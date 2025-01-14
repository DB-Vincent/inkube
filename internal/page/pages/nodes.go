package pages

import (
	"context"
	"fmt"
	"time"

	"github.com/DB-Vincent/go-eink-driver/pkg/display"
	"github.com/DB-Vincent/go-eink-driver/pkg/graphics"
	k8s "github.com/DB-Vincent/inkube/internal/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NodePage(display *display.Display, kubeConn *k8s.KubernetesConnection) error {
	nodes, err := kubeConn.Clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	display.ClearCanvas()

	for i, node := range nodes.Items {
		y := 10 + (i * 15)
		ready := "Not Ready"
		for _, condition := range node.Status.Conditions {
			if condition.Type == "Ready" {
				if condition.Status == "True" {
					ready = "Ready"
				}
				break
			}
		}
		graphics.Text(display.Canvas, 10, y, fmt.Sprintf("%s: %s", node.Name, ready))
	}

	graphics.Text(display.Canvas, 10, 105, fmt.Sprintf("Last refresh: %s", time.Now().Format("15:04:05")))
	display.DrawCanvas()

	return nil
}
