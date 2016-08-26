package main

import (
	"log"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

const (
	database = "homie"
)

func main() {
	c := influxDBClient()
	sensor := "temperature"
	loc := "home"
	logValue(c, sensor, loc, 123.4)
	time.Sleep(10 * time.Second)
	logValue(c, sensor, loc, 13.4)
}

func influxDBClient() client.Client {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	return c
}

func logValue(c client.Client, s string, l string, v float64) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  database,
		Precision: "s",
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	// Create a point and add to batch
	tags := map[string]string{"sensor": s, "location": l}
	fields := map[string]interface{}{
		"value": v,
	}
	pt, err := client.NewPoint("sensors", tags, fields, time.Now())

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	bp.AddPoint(pt)

	err = c.Write(bp)
	if err != nil {
		log.Fatal(err)
	}
}
