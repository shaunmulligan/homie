package main

import (
	"fmt"
	"log"

	"github.com/influxdata/influxdb/client/v2"
)

func error_check(err error) {
	if err != nil {
		fmt.Printf("error %s", err)
		log.Println("There was an error:", err)
	}
}

func error_fail(err error) {
	if err != nil {
		log.Fatalln("There was a fatal error:", err)
	}
}

type DbConfig struct {
	Address   string
	Username  string
	Password  string
	Database  string
	Precision string
	client    client.Client
}
