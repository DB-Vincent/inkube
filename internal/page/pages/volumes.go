// Persistent volume overview
//
// STORAGE USAGE
// -------------
// PVC: logs-vol  85% (9.8G/12G)
// PVC: db-vol    60% (6G/10G)
// PVC: cache     95% (950M/1G) [!!]

package pages

import (
  "fmt"
  "time"

	"github.com/DB-Vincent/go-eink-driver/pkg/display"
	"github.com/DB-Vincent/go-eink-driver/pkg/graphics"
	k8s "github.com/DB-Vincent/inkube/internal/kubernetes"
)

func VolumePage(display *display.Display, kubeConn *k8s.KubernetesConnection) error {
	display.ClearCanvas()

	graphics.Text(display.Canvas, 10, 0, "This is the volumes page")

  graphics.Text(display.Canvas, 10, 105, fmt.Sprintf("Last refresh: %s", time.Now().Format("15:04:05")))
	display.DrawCanvas()

	return nil
}
