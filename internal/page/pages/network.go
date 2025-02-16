// Network usage
//
// NETWORK USAGE
// -------------
// NS: default  In: 150MB  Out: 1.2GB
// NS: backend  In: 220MB  Out: 800MB
// NS: logging  In:  90MB  Out: 2.3GB

package pages

import (
  "fmt"
  "time"

	"github.com/DB-Vincent/go-eink-driver/pkg/display"
	"github.com/DB-Vincent/go-eink-driver/pkg/graphics"
	k8s "github.com/DB-Vincent/inkube/internal/kubernetes"
)

func NetworkPage(display *display.Display, kubeConn *k8s.KubernetesConnection) error {
	display.ClearCanvas()

	graphics.Text(display.Canvas, 10, 0, "This is the network page")

  graphics.Text(display.Canvas, 10, 105, fmt.Sprintf("Last refresh: %s", time.Now().Format("15:04:05")))
	display.DrawCanvas()

	return nil
}
