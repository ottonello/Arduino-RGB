package main

/*
 To run you need the Firmata firnmware on your Arduino,
 read: http://gobot.io/documentation/platforms/arduino/
*/
import (
	"math/rand"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
)

func main() {
	gbot := gobot.NewGobot()

	firmataAdaptor := firmata.NewFirmataAdaptor("arduino", "COM3")

	red := gpio.NewDirectPinDriver(firmataAdaptor, "pin", "6")
	green := gpio.NewDirectPinDriver(firmataAdaptor, "pin", "10")
	blue := gpio.NewDirectPinDriver(firmataAdaptor, "pin", "11")

	work := func() {
		gobot.Every(1*time.Second, func() {
			r := byte(rand.Intn(255))
			g := byte(rand.Intn(255))
			b := byte(rand.Intn(255))

			red.PwmWrite(r)
			green.PwmWrite(g)
			blue.PwmWrite(b)
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{red, green, blue},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
