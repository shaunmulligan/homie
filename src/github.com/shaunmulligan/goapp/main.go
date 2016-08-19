package main

import (
	"fmt"
	"time"

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

func getMeasurements(g *grovepi.GrovePi) Measurement {
	air, err := g.AnalogRead(grovepi.A0)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("Air Quality: %v\n", air)

	light, err := g.AnalogRead(grovepi.A1)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("Light Level: %v\n", light)

	t, err := g.Temp(grovepi.A2)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("Current Temperature is %f\n", t)

	con, err := g.ReadDustSensor()
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("dust concentration = %f pcf/0.01cf\n", con)

	m := Measurement{airQuality: air, lightLevel: light, temperature: t, dustLevel: con}
	return m
}

func main() {
	var g grovepi.GrovePi
	g = *grovepi.InitGrovePi(0x04)
	defer g.CloseDevice()
	time.Sleep(2 * time.Second)

	v, err := g.Version()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Grovepi Firmware Version: %v\n", v)

	err = g.EnableDustSensor()
	if err != nil {
		fmt.Println(err)
	}
	m := getMeasurements(&g)
	fmt.Println(m)
}
