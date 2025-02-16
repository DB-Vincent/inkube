// Nodes page
//
// NODE STATUS
// -----------
// node1  [OK]  CPU: 30%  MEM: 2.1G/4G
// node2  [!!]  CPU: 85%  MEM: 3.5G/4G
// node3  [X]   NotReady (DiskPressure)

package pages

import (
  "fmt"
  "time"

	"github.com/DB-Vincent/go-eink-driver/pkg/display"
	"github.com/DB-Vincent/go-eink-driver/pkg/graphics"
	k8s "github.com/DB-Vincent/inkube/internal/kubernetes"
)

func NodePage(display *display.Display, kubeConn *k8s.KubernetesConnection) error {
	display.ClearCanvas()

	graphics.Text(display.Canvas, 10, 0, "This is the node page")

  graphics.Text(display.Canvas, 10, 105, fmt.Sprintf("Last refresh: %s", time.Now().Format("15:04:05")))
	display.DrawCanvas()

	return nil
}
