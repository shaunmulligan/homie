package main

import (
	"flag"
	"fmt"

	"github.com/shaunmulligan/datalogger"
)

// ./createPoint -reading=12.3
func main() {
	var dbssl = flag.Bool("dbssl", false, "Whether to use HTTPS to connect to InfluxDB")
	var dbhost = flag.String("dbhost", "localhost", "Hostname of the InfluxDB server.")
	var dbport = flag.Int("dbport", 8086, "Port of the InfluxDB server")
	var dbuser = flag.String("dbuser", "", "Username for the InfluxDB")
	var dbpass = flag.String("dbpass", "", "Password for the InfluxDB")
	var dbname = flag.String("dbname", "homie", "Name of the InfluxDB database")
	var sensor = flag.String("sensor", "temperature", "type of sensor reading")
	var location = flag.String("location", "home", "location of sensor")
	var reading = flag.Float64("reading", 0.0, "value of the sensor reading")
	flag.Parse()

	dbprotocol := "http"
	if *dbssl {
		dbprotocol = "https"
	}
	db := datalogger.DbConfig{
		Address:   fmt.Sprintf("%s://%s:%d", dbprotocol, *dbhost, *dbport),
		Database:  *dbname,
		Precision: "s",
	}
	if *dbuser != "" && *dbpass != "" {
		db.Username = *dbuser
		db.Password = *dbpass
	}
	db.Connect()

	db.LogValue(*sensor, *location, *reading)
}
