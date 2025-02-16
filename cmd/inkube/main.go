package main

import (
	"fmt"
	"time"
  "strconv"

	"github.com/DB-Vincent/go-eink-driver/pkg/display"
	"github.com/DB-Vincent/go-eink-driver/pkg/spi"
	"github.com/DB-Vincent/inkube/internal/config"
	k8s "github.com/DB-Vincent/inkube/internal/kubernetes"
	"github.com/DB-Vincent/inkube/internal/page"
	"github.com/DB-Vincent/inkube/internal/page/pages"
)

func main() {
  configPath := "config.toml"
	config, err := config.Load(configPath)
	if err != nil {
		fmt.Printf("failed to load config: %v", err)
    return
	}

	spi, err := spi.New()
	if err != nil {
		fmt.Println(err)
		return
	}
	einkDisplay := display.New(spi, true)
	defer spi.Close()

	kubeConn := k8s.KubernetesConnection{}
	_, err = kubeConn.ConnectToCluster(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Initializing display")
	einkDisplay.Init()

	pm := page.NewManager(einkDisplay, &kubeConn)
	pm.AddPage(pages.PodPage)
	pm.AddPage(pages.ClusterPage)

  refreshRate, _ := strconv.Atoi(config.Display.Refresh) 

	for {
		if err := pm.CurrentPage(); err != nil {
			fmt.Println(err)
		}

		time.Sleep(time.Duration(refreshRate) * time.Second)
		pm.NextPage()
	}
}
