package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	fast "github.com/ddo/go-fast"
	"github.com/jasonlvhit/gocron"
	"github.com/shaunmulligan/goapp/grovepi"
)

//A Measurement looks like this:
// Air Quality: 29
// Light Level: 480
// Current Temperature is 21.556412
// dust concentration = 546.579531 pcf/0.01cf
type Measurement struct {
	airQuality  int
	lightLevel  int
	temperature float64
	dustLevel   float64
}

type Sensors struct {
	*grovepi.GrovePi
}

// InitSensors enables GrovePi hat and the dust Sensor
func InitSensors() (Sensors, error) {
	g := grovepi.InitGrovePi(0x04)
	time.Sleep(2 * time.Second)
	//show the version of GrovePi Firmware
	v, err := g.Version()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Grovepi Firmware Version: %v\n", v)

	// Enable the dust sensor
	err = g.EnableDustSensor()
	if err != nil {
		fmt.Printf("damn errors: %v", err)
	}

	return Sensors{g}, nil
}

func CleanUpSensors(s Sensors) {
	s.CloseDevice()
}

func getMeasurements(s Sensors) *Measurement {

	air, err := s.AnalogRead(grovepi.A0)
	if err != nil {
		fmt.Printf("come on... %v", err)
	}

	light, err := s.AnalogRead(grovepi.A1)
	if err != nil {
		fmt.Printf("come on... %v", err)
	}

	temp, err := s.Temp(grovepi.A2)
	if err != nil {
		fmt.Printf("come on... %v", err)
	}

	dust, err := s.ReadDustSensor()
	if err != nil {
		fmt.Printf("come on... %v", err)
	}
	// Enable the dust sensor
	err = s.EnableDustSensor()
	if err != nil {
		fmt.Printf("damn errors: %v", err)
	}
	return &Measurement{airQuality: air, lightLevel: light, temperature: temp, dustLevel: dust}
}

func main() {
	g, _ := InitSensors()
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		sig := <-signalChannel
		switch sig {
		case syscall.SIGKILL:
			//handle SIGINT
			fmt.Println("got SIGKILL")
			CleanUpSensors(g)
		case syscall.SIGTERM:
			//handle SIGTERM
			fmt.Println("got SIGTERM")
			CleanUpSensors(g)
		}
	}()

	db := DbConfig{
		Address:   "http://127.0.0.1:8086",
		Database:  "homie",
		Precision: "s",
	}
	db.Connect()

	fmt.Println("Starting looper")
	gocron.Every(5).Minutes().Do(db.insertMeasurement, g)
	gocron.Every(20).Minutes().Do(db.measureInternet)
	<-gocron.Start()

}

// TODO: break this into separate goroutines for each sensor
func (db *DbConfig) insertMeasurement(s Sensors) {
	m := getMeasurements(s)
	fmt.Printf("===========================================\n")
	db.LogValue("air", "home", float64(m.airQuality))
	db.LogValue("light", "home", float64(m.lightLevel))
	db.LogValue("dust", "home", m.dustLevel)
	db.LogValue("temperature", "home", m.temperature)

}

func (db *DbConfig) measureInternet() {
	fastCom := fast.New()
	err := fastCom.Init()
	if err != nil {
		panic(err)
	}
	// get urls
	urls, err := fastCom.GetUrls()
	if err != nil {
		panic(err)
	}

	// measure
	KbpsChan := make(chan float64)
	go func() {
		sumK := 0.0
		i := 0.0
		for Kbps := range KbpsChan {
			// fmt.Printf("%.2f Kbps %.2f Mbps\n", Kbps, Kbps/1000)
			sumK = sumK + Kbps
			i++
		}
		avgSpeed := (sumK / i) / 1000
		fmt.Printf("Average Internet Speed: %v\n", avgSpeed)
		db.LogValue("internet", "home", avgSpeed)
	}()

	err = fastCom.Measure(urls, KbpsChan)

	if err != nil {
		fmt.Println("I am gonna panic!!")
		panic(err)
	}
}
