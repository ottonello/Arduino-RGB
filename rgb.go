package main

import (
	"log"
	"net/http"
	"time"
)

type Colores struct {
	run bool
	ar  DuinoAr
}

func main() {
	log.Println("Starting")

	main := Colores{true, Connect()}
	// wait for the duino to reset
	time.Sleep(time.Second * 2)

	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Start received")
		go main.start()
	})

	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Stop received")
		main.stop()
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (c *Colores) start() {
	c.run = true
	for c.run {
		log.Printf("Running %#v", c)
		go c.ar.receive()

		c.ar.sendRandomColor()
		time.Sleep(time.Second * 1)
	}
}

func (c *Colores) stop() {
	c.run = false
	log.Printf("Stopped %#v", c)
}
