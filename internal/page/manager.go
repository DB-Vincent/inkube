package page

import (
	"github.com/DB-Vincent/go-eink-driver/pkg/display"
	"k8s.io/client-go/kubernetes"
)

type PageManager struct {
	display     *display.Display
	client      *kubernetes.Clientset
	pages       []func(*display.Display, *kubernetes.Clientset) error
	currentPage int
}

func NewManager(eink *display.Display, client *kubernetes.Clientset) *PageManager {
	return &PageManager{
		display:     eink,
		client:      client,
		pages:       make([]func(*display.Display, *kubernetes.Clientset) error, 0),
		currentPage: 0,
	}
}

func (pm *PageManager) AddPage(page func(*display.Display, *kubernetes.Clientset) error) {
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
	return pm.pages[pm.currentPage](pm.display, pm.client)
}
