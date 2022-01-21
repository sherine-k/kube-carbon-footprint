package kube

import (
	"os"

	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var klog = logrus.WithField("module", "kube")

type Client struct {
	clientset *kubernetes.Clientset
}

func LoadKubeClient(path string) (*Client, error) {
	cls, err := kubernetes.NewForConfig(loadKubeConfig(path))
	if err != nil {
		return nil, err
	}
	return &Client{
		clientset: cls,
	}, nil
}

func loadKubeConfig(path string) *rest.Config {
	var config *rest.Config
	var err error
	if path != "" {
		flog := klog.WithField("kubeConfig", path)
		flog.Info("Using command line supplied kube config")
		config, err = clientcmd.BuildConfigFromFlags("", path)
		if err != nil {
			flog.WithError(err).Fatal("Can't load kube config file")
		}
	} else if kfgPath := os.Getenv("KUBECONFIG"); kfgPath != "" {
		flog := klog.WithField("kubeConfig", kfgPath)
		flog.Info("Using environment KUBECONFIG")
		config, err = clientcmd.BuildConfigFromFlags("", kfgPath)
		if err != nil {
			flog.WithError(err).WithField("kubeConfig", kfgPath).
				Fatal("can't find provided KUBECONFIG env path")
		}
	} else {
		klog.Info("Using in-cluster kube config")
		config, err = rest.InClusterConfig()
		if err != nil {
			klog.WithError(err).Fatal("can't load in-cluster REST config")
		}
	}
	return config
}
