package main

import (
	"time"

	wemo "github.com/danward79/go.wemo"
	log "github.com/sirupsen/logrus"
)

func scheduler(device Device, eventqueue chan Event) {
	log.Info("Scheduler started for ", device.Name)
	for {
		eventqueue <- Event{Name: device.Name, Address: device.Address, State: "On"}
		time.Sleep(device.DutyCycle)

		eventqueue <- Event{Name: device.Name, Address: device.Address, State: "Off"}
		time.Sleep(device.Frequency - device.DutyCycle)
	}
}

func control(e Event) {
	device := &wemo.Device{Host: e.Address}

	switch e.State {
	case "On":
		log.Info(e.Name, " is ", device.GetBinaryState())
		device.On()
	case "Off":
		log.Info(e.Name, " is ", device.GetBinaryState())
		device.Off()
	}
}
