package pump

import (
	"log"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/host/v3"
	"periph.io/x/host/v3/rpi"
)

func init() {
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	t := time.NewTicker(500 * time.Millisecond)
	l := gpio.Low
	for {
		l = !l

		if err := rpi.P1_33.Out(l); err != nil {
			log.Fatal(err)
		}
		<-t.C
	}
	for l := gpio.Low; ; l = !l {
		// Lookup a pin by its location on the board:
		if err := rpi.P1_33.Out(l); err != nil {
			log.Fatal(err)
		}
		<-t.C
	}
}
