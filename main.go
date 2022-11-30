package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/kardianos/osext"
	log "github.com/sirupsen/logrus"
)

//Global
var config Configuration

func init() {
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal("Unable to get local directory ", err)
	}

	configFile, err := os.Open(folderPath + "\\config.json")
	if err != nil {
		log.Fatal("Config file not found ", err)
	}
	defer configFile.Close()

	byteValue, err := ioutil.ReadAll(configFile)
	json.Unmarshal(byteValue, &config)
}

func main() {
	eventqueue := make(chan Event)

	for _, d := range config.Devices {
		go scheduler(d, eventqueue)
	}

	for {
		control(<-eventqueue)
	}
}
