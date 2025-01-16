package pages

import (
	"context"
	"fmt"
	"time"

	"github.com/DB-Vincent/go-eink-driver/pkg/display"
	"github.com/DB-Vincent/go-eink-driver/pkg/graphics"
	k8s "github.com/DB-Vincent/inkube/internal/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type PodCount struct {
	Running   int
	Pending   int
	Succeeded int
	Failed    int
	Unknown   int
}

func getPodCount(client *kubernetes.Clientset) (PodCount, error) {
	pods, err := client.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return PodCount{}, fmt.Errorf("failed to list pods: %v", err)
	}

	podCount := PodCount{}
	podCount.Running, podCount.Pending, podCount.Succeeded, podCount.Failed, podCount.Unknown = 0, 0, 0, 0, 0

	for _, pod := range pods.Items {
		switch pod.Status.Phase {
		case "Running":
			podCount.Running++
		case "Pending":
			podCount.Pending++
		case "Succeeded":
			podCount.Succeeded++
		case "Failed":
			podCount.Failed++
		default:
			podCount.Unknown++
		}
	}
	return podCount, nil
}

func PodPage(display *display.Display, kubeConn *k8s.KubernetesConnection) error {
	podCount, err := getPodCount(kubeConn.Clientset)
	if err != nil {
		return err
	}

	display.ClearCanvas()
	graphics.Text(display.Canvas, 10, 0, fmt.Sprintf("%d running", podCount.Running))
	graphics.Text(display.Canvas, 10, 15, fmt.Sprintf("%d pending", podCount.Pending))
	graphics.Text(display.Canvas, 10, 30, fmt.Sprintf("%d succeeded", podCount.Succeeded))
	graphics.Text(display.Canvas, 10, 45, fmt.Sprintf("%d failed", podCount.Failed))
	graphics.Text(display.Canvas, 10, 60, fmt.Sprintf("%d unknown", podCount.Unknown))
	graphics.Text(display.Canvas, 10, 105, fmt.Sprintf("Last refresh: %s", time.Now().Format("15:04:05")))
	display.DrawCanvas()

	return nil
}
