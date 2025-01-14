package pages

import (
	"context"
	"fmt"
	"time"

	"github.com/DB-Vincent/go-eink-driver/pkg/display"
	"github.com/DB-Vincent/go-eink-driver/pkg/graphics"
	k8s "github.com/DB-Vincent/inkube/internal/kubernetes"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	metricsclientset "k8s.io/metrics/pkg/client/clientset/versioned"
)

type CPU struct {
	UsedCores    float64
	TotalCores   float64
	UsagePercent float64
}

type Memory struct {
	UsedBytes    int64
	TotalBytes   int64
	UsagePercent float64
}

type NodeMetrics struct {
	CPU
	Memory
}

func GetNodeResourceUsage(clientset *kubernetes.Clientset, metricsClient *metricsclientset.Clientset) (NodeMetrics, error) {
	nodeMetrics, err := metricsClient.MetricsV1beta1().NodeMetricses().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return NodeMetrics{}, fmt.Errorf("failed to get node metrics: %v", err)
	}

	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return NodeMetrics{}, fmt.Errorf("failed to get nodes: %v", err)
	}

	allocatable := make(map[string]v1.ResourceList)
	for _, node := range nodes.Items {
		allocatable[node.Name] = node.Status.Allocatable
	}

	nodeMetric := NodeMetrics{
		CPU: CPU{
			UsedCores:    0,
			TotalCores:   0,
			UsagePercent: 0,
		},
		Memory: Memory{
			UsedBytes:    0,
			TotalBytes:   0,
			UsagePercent: 0,
		},
	}

	for _, metric := range nodeMetrics.Items {
		if resources, ok := allocatable[metric.Name]; ok {
			cpuUsage := metric.Usage.Cpu().AsApproximateFloat64()
			cpuCapacity := resources.Cpu().AsApproximateFloat64()
			nodeMetric.CPU.UsedCores += cpuUsage
			nodeMetric.CPU.TotalCores += cpuCapacity

			memoryUsage := metric.Usage.Memory().Value()
			memoryCapacity := resources.Memory().Value()
			nodeMetric.Memory.UsedBytes += memoryUsage
			nodeMetric.Memory.TotalBytes += memoryCapacity

		}
	}
	nodeMetric.CPU.UsagePercent = (nodeMetric.CPU.UsedCores / nodeMetric.CPU.TotalCores) * 100
	nodeMetric.Memory.UsagePercent = (float64(nodeMetric.Memory.UsedBytes) / float64(nodeMetric.Memory.TotalBytes)) * 100

	return nodeMetric, nil
}

func ClusterPage(display *display.Display, kubeConn *k8s.KubernetesConnection) error {
	display.ClearCanvas()

	nodeMetrics, err := GetNodeResourceUsage(kubeConn.Clientset, kubeConn.MetricsClient)
	if err != nil {
		return fmt.Errorf("failed to get node metrics: %v", err)
	}

	graphics.Text(display.Canvas, 10, 15, fmt.Sprintf("CPU: %.2f/%.2fCores (%.2f%%)", nodeMetrics.CPU.UsedCores, nodeMetrics.CPU.TotalCores, nodeMetrics.CPU.UsagePercent))
	graphics.Text(display.Canvas, 10, 30, fmt.Sprintf("Memory: %d/%dMB (%.2f%%)", (nodeMetrics.Memory.UsedBytes/1024/1024), (nodeMetrics.Memory.TotalBytes/1024/1024), nodeMetrics.Memory.UsagePercent))

	graphics.Text(display.Canvas, 10, 105, fmt.Sprintf("Last refresh: %s", time.Now().Format("15:04:05")))
	display.DrawCanvas()

	return nil
}
