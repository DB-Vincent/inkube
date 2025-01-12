package main

import (
	"context"
	"fmt"
	"time"

	"github.com/DB-Vincent/go-eink-driver/pkg/display"
	"github.com/DB-Vincent/go-eink-driver/pkg/graphics"
	"github.com/DB-Vincent/go-eink-driver/pkg/spi"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	k8s "github.com/DB-Vincent/inkube/internal/kubernetes"
)

func getPodCount(client *kubernetes.Clientset) (map[string]int, error) {
	pods, err := client.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list pods: %v", err)
	}

	running := 0
	pending := 0
	succeeded := 0
	failed := 0
	unknown := 0

	for _, pod := range pods.Items {
		switch pod.Status.Phase {
		case "Running":
			running++
		case "Pending":
			pending++
		case "Succeeded":
			succeeded++
		case "Failed":
			failed++
		default:
			unknown++
		}
	}

	return map[string]int{
		"running":   running,
		"pending":   pending,
		"succeeded": succeeded,
		"failed":    failed,
		"unknown":   unknown,
	}, nil
}

func showPodPage(display *display.Display, podCount map[string]int) {
	currentTime := time.Now().Format("15:04:05")
	display.ClearCanvas()

	graphics.Text(display.Canvas, 10, 0, fmt.Sprintf("%d running", podCount["running"]))
	graphics.Text(display.Canvas, 10, 15, fmt.Sprintf("%d pending", podCount["pending"]))
	graphics.Text(display.Canvas, 10, 30, fmt.Sprintf("%d succeeded", podCount["succeeded"]))
	graphics.Text(display.Canvas, 10, 45, fmt.Sprintf("%d failed", podCount["failed"]))
	graphics.Text(display.Canvas, 10, 60, fmt.Sprintf("%d unknown", podCount["unknown"]))

	graphics.Text(display.Canvas, 10, 100, fmt.Sprintf("Last refresh: %s", currentTime))

	display.DrawCanvas()
}

func main() {
	spi, err := spi.New()
	if err != nil {
		fmt.Println(err)
		return
	}

	display := display.New(spi, true)
	defer spi.Close()

	client, err := k8s.ConnectToCluster()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Initializing display")
	display.Init()

	for {
		podCount, err := getPodCount(client)
		if err != nil {
			fmt.Println(err)
			return
		}
		showPodPage(display, podCount)

		time.Sleep(30 * time.Second)
	}
}
