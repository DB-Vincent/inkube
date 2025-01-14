package main

import (
	"fmt"
	"time"

	"github.com/DB-Vincent/go-eink-driver/pkg/display"
	"github.com/DB-Vincent/go-eink-driver/pkg/spi"
	k8s "github.com/DB-Vincent/inkube/internal/kubernetes"
	"github.com/DB-Vincent/inkube/internal/page"
	"github.com/DB-Vincent/inkube/internal/page/pages"
)

func main() {
	spi, err := spi.New()
	if err != nil {
		fmt.Println(err)
		return
	}
	einkDisplay := display.New(spi, true)
	defer spi.Close()

	kubeConn := k8s.KubernetesConnection{}
	_, err = kubeConn.ConnectToCluster()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Initializing display")
	einkDisplay.Init()

	pm := page.NewManager(einkDisplay, &kubeConn)
	pm.AddPage(pages.PodPage)
	pm.AddPage(pages.NodePage)
	pm.AddPage(pages.ClusterPage)

	for {
		if err := pm.CurrentPage(); err != nil {
			fmt.Println(err)
		}

		time.Sleep(30 * time.Second)
		pm.NextPage()
	}
}
