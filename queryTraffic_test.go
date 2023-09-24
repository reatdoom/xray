package main

import (
	"fmt"
	"testing"
)

func TestQueryTraffic(t *testing.T) {
	var (
		xrayCtl *XrayController
		cfg     = &BaseConfig{
			APIAddress: "127.0.0.1",
			APIPort:    10085,
		}
	)
	xrayCtl = new(XrayController)
	err := xrayCtl.Init(cfg)
	defer xrayCtl.CmdConn.Close()
	if err != nil {
		t.Errorf("Failed %s", err)
	}
	ptn := "user>>>love@xray.com>>>traffic>>>uplink"
	trafficData, err := queryTraffic(xrayCtl.SsClient, ptn, false)
	if err != nil {
		t.Errorf("Failed %s", err)
	}
	fmt.Println(trafficData)

}
