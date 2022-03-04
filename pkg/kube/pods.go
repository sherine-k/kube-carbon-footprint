package kube

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const regionLabel = "topology.kubernetes.io/region"
const instanceTypeLabel = "node.kubernetes.io/instance-type"

type DatacenterInfo struct {
	Region       string
	InstanceType string
}

func (c *Client) GetPodDatacenterInfo(podName, namespace string) (*DatacenterInfo, error) {
	pod, err := c.clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	if pod.Spec.NodeName == "" {
		return nil, fmt.Errorf("no node name set for pod %s [namespace %s]", podName, namespace)
	}
	node, err := c.clientset.CoreV1().Nodes().Get(context.TODO(), pod.Spec.NodeName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return &DatacenterInfo{
		Region:       node.Labels[regionLabel],
		InstanceType: node.Labels[instanceTypeLabel],
	}, nil
}
