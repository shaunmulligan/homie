package main

import (
	"fmt"
	"log"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

// LogValue Process sensor input and store in influxdb
func (db DbConfig) LogValue(s string, l string, v float64) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		//TODO: probably shouldn't hardcode database name
		Database:        "homie",
		Precision:       "s",
		RetentionPolicy: "four_months",
	})

	error_fail(err)

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
	err = db.client.Write(bp)
	log.Printf("sensor: %v, val: %v\n", s, v)
	error_fail(err)

}

// Connect to influxdb instance
func (db *DbConfig) Connect() {
	var err error
	db.client, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     db.Address,
		Username: db.Username,
		Password: db.Password,
	})
	error_fail(err)
	log.Println("Conected to DB.")
	q := client.NewQuery("CREATE DATABASE homie", "", "")
	if response, err := db.client.Query(q); err == nil && response.Error() == nil {
		fmt.Println(response.Results)
	}
	p := client.NewQuery("CREATE RETENTION POLICY four_months ON homie DURATION 16w REPLICATION 1 DEFAULT", "", "")
	if response, err := db.client.Query(p); err == nil && response.Error() == nil {
		fmt.Println(response.Results)
	}
}
