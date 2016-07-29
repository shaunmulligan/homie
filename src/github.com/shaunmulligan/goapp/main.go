package main

import (
	"fmt"
	"time"

	"github.com/shaunmulligan/goapp/grovepi"
)

func main() {
	var g grovepi.GrovePi
	g = *grovepi.InitGrovePi(0x04)
	defer g.CloseDevice()
	time.Sleep(2 * time.Second)

	v, err := g.Version()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)

	air, err := g.AnalogRead(grovepi.A0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(air)

	light, err := g.AnalogRead(grovepi.A1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(light)

	temp, err := g.AnalogRead(grovepi.A2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(temp)

	t, err := g.Temp(grovepi.A2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)

	err = g.EnableDustSensor()
	if err != nil {
		fmt.Println(err)
	}

	con, err := g.ReadDustSensor()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("dust concentration = %f pcf/0.01cf\n", con)

}
