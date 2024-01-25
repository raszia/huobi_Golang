package main

import (
	"github.com/raszia/huobi_Golang/cmd/accountclientexample"
	"github.com/raszia/huobi_Golang/cmd/accountwebsocketclientexample"
	"github.com/raszia/huobi_Golang/cmd/algoorderclientexample"
	"github.com/raszia/huobi_Golang/cmd/commonclientexample"
	"github.com/raszia/huobi_Golang/cmd/crossmarginclientexample"
	"github.com/raszia/huobi_Golang/cmd/etfclientexample"
	"github.com/raszia/huobi_Golang/cmd/isolatedmarginclientexample"
	"github.com/raszia/huobi_Golang/cmd/marketclientexample"
	"github.com/raszia/huobi_Golang/cmd/marketwebsocketclientexample"
	"github.com/raszia/huobi_Golang/cmd/orderclientexample"
	"github.com/raszia/huobi_Golang/cmd/orderwebsocketclientexample"
	"github.com/raszia/huobi_Golang/cmd/stablecoinclientexample"
	"github.com/raszia/huobi_Golang/cmd/subuserclientexample"
	"github.com/raszia/huobi_Golang/cmd/walletclientexample"
	"github.com/raszia/huobi_Golang/logging/perflogger"
)

func main() {
	marketwebsocketclientexample.RunTicker()
	// runAll()
}

// Run all examples
func runAll() {
	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	algoorderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	subuserclientexample.RunAllExamples()
	stablecoinclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
	marketwebsocketclientexample.RunAllExamples()
	accountwebsocketclientexample.RunAllExamples()
	orderwebsocketclientexample.RunAllExamples()
}

// Run performance test
func runPerfTest() {
	perflogger.Enable(true)

	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
}
