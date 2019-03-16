package main

import "time"

//Event passed to channel to control devices
type Event struct {
	Name    string
	Address string
	State   string
}

//Device defines characteristics of a Wemo device in config
type Device struct {
	Name      string        `json:"Name"`
	Address   string        `json:"Address"`
	DutyCycle time.Duration `json:"DutyCycle"`
	Frequency time.Duration `json:"Frequency"`
}

//Configuration defines the expected layout of the config file
type Configuration struct {
	NetworkInterface string   `json:"NetworkInterface"`
	Devices          []Device `json:"Devices"`
}
