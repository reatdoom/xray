package main

import (
	"context"
	statsService "github.com/xtls/xray-core/app/stats/command"
)

func queryTraffic(c statsService.StatsServiceClient, ptn string, reset bool) (traffic int64, err error) {
	traffic = -1
	resp, err := c.QueryStats(context.Background(), &statsService.QueryStatsRequest{
		Pattern: ptn,
		Reset_: reset,
	})
	if err != nil {
		return
	}
	// Get traffic data
	stat := resp.GetStat()
	if len(stat) != 0 {
		traffic = stat[0].Value
	}

	return
}
