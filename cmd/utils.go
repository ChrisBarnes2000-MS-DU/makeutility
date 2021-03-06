package cmd

import "time"

// GTZone stand for General Time Zone
var GTZone = `2006-01-02T15:04:05.000-07:00`

func getTimeStamp(date string) int {
	// https://play.golang.org/p/ouiDtIVjQI
	t, e := time.Parse(GTZone, date)
	if e != nil {
		panic(e)
	}

	return int(t.UTC().UnixNano() / 1000000)
}
