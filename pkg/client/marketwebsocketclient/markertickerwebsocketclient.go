package marketwebsocketclient

import (
	"encoding/json"
	"fmt"

	"github.com/raszia/huobi_Golang/logging/applogger"
	"github.com/raszia/huobi_Golang/pkg/client/websocketclientbase"
	"github.com/raszia/huobi_Golang/pkg/model/market"
)

// Responsible to handle last 24h marketTicker data from WebSocket
type MarketTickerWebSocketClient struct {
	websocketclientbase.WebSocketClientBase
}

// Initializer
func (p *MarketTickerWebSocketClient) Init(host string) *MarketTickerWebSocketClient {
	p.WebSocketClientBase.Init(host)
	return p
}

// Set callback handler
func (p *MarketTickerWebSocketClient) SetHandler(
	connectedHandler websocketclientbase.ConnectedHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, p.handleMessage, responseHandler)
}

// Request full marketTicker data
func (p *MarketTickerWebSocketClient) Request(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.ticker", symbol)
	req := fmt.Sprintf("{\"req\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(req)

	applogger.Info("WebSocket requested, topic=%s, clientId=%s", topic, clientId)
}

// Subscribe latest 24h market stats
func (p *MarketTickerWebSocketClient) Subscribe(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.ticker", symbol)
	sub := fmt.Sprintf("{\"sub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(sub)

	applogger.Info("WebSocket subscribed, topic=%s, clientId=%s", topic, clientId)
}

// Unsubscribe latest 24 market stats
func (p *MarketTickerWebSocketClient) UnSubscribe(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.ticker", symbol)
	unsub := fmt.Sprintf("{\"unsub\": \"%s\",\"id\": \"%s\" }", symbol, clientId)

	p.Send(unsub)

	applogger.Info("WebSocket unsubscribed, topic=%s, clientId=%s", topic, clientId)
}

func (p *MarketTickerWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubscribeMarketTickerResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
