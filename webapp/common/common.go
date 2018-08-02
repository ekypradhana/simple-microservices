package common

import (
	"encoding/json"
	"log"
	"os"
)

type (
	configuration struct {
		ListeningPort, TransactionService, ProductService string
	}
)

// AppConfig holds the configuration values from config.json file
var AppConfig configuration
var jsonFile string = "config.json"

// Initialize AppConfig
func initConfig() {
	file, err := os.Open(jsonFile)
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
}

func StartUp(configFile string) {
	if configFile != "" {
		jsonFile = configFile
	}

	// Initialize AppConfig variable
	initConfig()
}
