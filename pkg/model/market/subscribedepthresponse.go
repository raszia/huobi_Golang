package market

import (
	"github.com/raszia/huobi_Golang/pkg/model/base"
)

type SubscribeDepthResponse struct {
	base.WebSocketResponseBase
	Data *Depth
	Tick *Depth
}
