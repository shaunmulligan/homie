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
	for {
		time.Sleep(2 * time.Second)
		t, h, err := g.ReadDHT(grovepi.D4)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(t)
		fmt.Println(h)
	}
}
