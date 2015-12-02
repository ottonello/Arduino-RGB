package main

import (
	"bufio"
	"log"
	"math/rand"
	"time"

	"github.com/tarm/serial"
)

type DuinoAr struct {
	s *serial.Port
}

func Connect() (ar DuinoAr) {
	c := &serial.Config{Name: "COM3", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	return DuinoAr{s}
}

func (d DuinoAr) sendRandomColor() {
	color := randomColor()
	_, err := d.s.Write(color)

	if err != nil {
		log.Fatal(err)
	}
}

func (d DuinoAr) receive() {
	reader := bufio.NewReader(d.s)
	reply, err := reader.ReadBytes('\x0a')
	if err != nil {
		panic(err)
	}
	log.Printf("%q", reply)
}

func randomColor() (rgb []byte) {
	rand.Seed(time.Now().UnixNano()) // Try changing this number!
	return []byte{'s', byte(rand.Intn(255)), byte(rand.Intn(255)), byte(rand.Intn(255))}
}
