package common

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"gopkg.in/mgo.v2"
)

type (
	configuration struct {
		ServerPort, MongoDBHost, MongoDBUser, MongoDBPwd, Database, ProductService string
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

// Session holds the mongodb session for database access
var session *mgo.Session

// Get database session
func GetSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{AppConfig.MongoDBHost},
			Username: AppConfig.MongoDBUser,
			Password: AppConfig.MongoDBPwd,
			Timeout:  60 * time.Second,
		})
		if err != nil {
			log.Fatalf("[GetSession]: %s\n", err)
		}
	}
	return session
}

// Create database session
func createDbSession() {
	var err error
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{AppConfig.MongoDBHost},
		Username: AppConfig.MongoDBUser,
		Password: AppConfig.MongoDBPwd,
		Timeout:  60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[createDbSession]: %s\n", err)
	}
}

func StartUp(configFile string) {
	if configFile != "" {
		jsonFile = configFile
	}

	// Initialize AppConfig variable
	initConfig()
	// Start a MongoDB session
	createDbSession()
}
