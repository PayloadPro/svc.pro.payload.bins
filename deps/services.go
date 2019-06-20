package deps

import (
	"github.com/PayloadPro/svc.pro.payload.bins/services"
)

// Services wrapped in a container
type Services struct {
	Bin *services.BinService
}
