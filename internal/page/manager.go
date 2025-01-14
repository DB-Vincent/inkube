package page

import (
	"github.com/DB-Vincent/go-eink-driver/pkg/display"
	k8s "github.com/DB-Vincent/inkube/internal/kubernetes"
)

type PageManager struct {
	display     *display.Display
	kubeConn    *k8s.KubernetesConnection
	pages       []func(*display.Display, *k8s.KubernetesConnection) error
	currentPage int
}

func NewManager(eink *display.Display, kubeConn *k8s.KubernetesConnection) *PageManager {
	return &PageManager{
		display:     eink,
		kubeConn:    kubeConn,
		pages:       make([]func(*display.Display, *k8s.KubernetesConnection) error, 0),
		currentPage: 0,
	}
}

func (pm *PageManager) AddPage(page func(*display.Display, *k8s.KubernetesConnection) error) {
	pm.pages = append(pm.pages, page)
}

func (pm *PageManager) NextPage() {
	if (pm.currentPage + 1) >= len(pm.pages) {
		pm.currentPage = 0
	} else {
		pm.currentPage++
	}
}

func (pm *PageManager) CurrentPage() error {
	return pm.pages[pm.currentPage](pm.display, pm.kubeConn)
}
