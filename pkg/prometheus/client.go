package prometheus

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/config"
	"github.com/prometheus/common/model"
	"github.com/sirupsen/logrus"
)

var clog = logrus.WithField("module", "prometheus")

type Config struct {
	URL                string
	Timeout            time.Duration
	Token              string
	InsecureSkipVerify bool
}

type Client struct {
	cfg   Config
	v1api v1.API
}

func roundTripper(cfg Config) http.RoundTripper {
	t := http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   cfg.Timeout,
			KeepAlive: cfg.Timeout,
		}).DialContext,
		TLSHandshakeTimeout: cfg.Timeout,
	}
	if cfg.InsecureSkipVerify {
		t.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	if cfg.Token != "" {
		return config.NewAuthorizationCredentialsRoundTripper("Bearer", config.Secret(cfg.Token), &t)
	}
	return &t
}

func NewClient(cfg Config) (*Client, error) {
	transport := roundTripper(cfg)
	client, err := api.NewClient(api.Config{
		Address:      cfg.URL,
		RoundTripper: transport,
	})
	if err != nil {
		return nil, err
	}
	v1api := v1.NewAPI(client)
	return &Client{
		cfg:   cfg,
		v1api: v1api,
	}, nil
}

func (c *Client) executePrometheusQuery(query string) (model.Matrix, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.cfg.Timeout)
	defer cancel()
	r := v1.Range{
		Start: time.Now().Add(-time.Hour),
		End:   time.Now(),
		Step:  time.Minute,
	}
	result, warnings, err := c.v1api.QueryRange(ctx, query, r)
	if err != nil {
		return nil, err
	}
	if len(warnings) > 0 {
		clog.Warnf("Warnings: %v", warnings)
	}
	if result.Type() == model.ValMatrix {
		return result.(model.Matrix), nil
	}
	return nil, fmt.Errorf("invalid query, matrix expected: %s", query)
}

func (c *Client) GetCPUAvgByPod(pod string, namespace string) (model.Matrix, error) {
	query := fmt.Sprintf("(sum(node_namespace_pod_container:container_cpu_usage_seconds_total:sum_irate{pod='%s', namespace='%s'}) by (namespace,node,pod) / on(node) group_left() machine_cpu_cores) * 100", pod, namespace)
	return c.executePrometheusQuery(query)
}

func (c *Client) GetCPUMetrics() (model.Matrix, error) {
	query := "(node_namespace_pod_container:container_cpu_usage_seconds_total:sum_irate / on(node) group_left() machine_cpu_cores) * 100"
	return c.executePrometheusQuery(query)
}
