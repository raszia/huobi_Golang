package market

import (
	"github.com/raszia/huobi_Golang/pkg/model/base"
	"github.com/shopspring/decimal"
)

type SubscribeMarketTickerResponse struct {
	base.WebSocketResponseBase
	Tick *MarkerTick
}

type MarkerTick struct {
	Amount    decimal.Decimal `json:"amount"`
	Count     int             `json:"count"`
	Open      decimal.Decimal `json:"open"`
	Close     decimal.Decimal `json:"close"`
	Low       decimal.Decimal `json:"low"`
	High      decimal.Decimal `json:"high"`
	Vol       decimal.Decimal `json:"vol"`
	Bid       decimal.Decimal `json:"bid"`
	BidSize   decimal.Decimal `json:"bidSize"`
	Ask       decimal.Decimal `json:"ask"`
	AskSize   decimal.Decimal `json:"askSize"`
	LastPrice decimal.Decimal `json:"lastPrice"`
	LastSize  decimal.Decimal `json:"lastSize"`
}
