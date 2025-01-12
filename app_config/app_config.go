package app_config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ENV ENV
}

type ENV struct {
	PORT string

	// core database
	DATABASE_SOURCE         string
	DATABASE_NAME           string
	COLLECTION_USERS        string
	COLLECTION_PROFILES     string
	COLLECTION_MATCHES      string
	TESTING_DATABASE_SOURCE string
	TESTING_DATABASE_NAME   string
}

var (
	appConfig     *AppConfig
	appConfigTest *AppConfig
)

// Get initialized config.
func Get(isProd bool) *AppConfig {
	if !isProd {
		if appConfigTest == nil {
			appConfigTest = createAppConfigTest()
		}
		return appConfigTest
	}

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	if err := checkENV(); err != nil {
		err = godotenv.Load()
		if err != nil {
			log.Println(" (-) file .env not found, using global variable")
		}

		err = checkENV()
		if err != nil {
			log.Panic(err)
		}
	}

	return createAppConfig()
}

func checkENV() error {
	if _, ok := os.LookupEnv("PORT"); !ok {
		return fmt.Errorf(" (x) config: PORT has not been set")
	}
	if _, ok := os.LookupEnv("DATABASE_SOURCE"); !ok {
		return fmt.Errorf(" (x) config: DATABASE_SOURCE has not been set")
	}
	if _, ok := os.LookupEnv("DATABASE_NAME"); !ok {
		return fmt.Errorf(" (x) config: DATABASE_NAME has not been set")
	}
	if _, ok := os.LookupEnv("COLLECTION_USERS"); !ok {
		return fmt.Errorf(" (x) config: COLLECTION_USERS has not been set")
	}
	if _, ok := os.LookupEnv("COLLECTION_PROFILES"); !ok {
		return fmt.Errorf(" (x) config: COLLECTION_PROFILES has not been set")
	}
	if _, ok := os.LookupEnv("COLLECTION_MATCHES"); !ok {
		return fmt.Errorf(" (x) config: COLLECTION_MATCHES has not been set")
	}
	return nil
}

func createAppConfig() *AppConfig {
	return &AppConfig{
		ENV: ENV{
			PORT:                os.Getenv("PORT"),
			DATABASE_SOURCE:     os.Getenv("DATABASE_SOURCE"),
			DATABASE_NAME:       os.Getenv("DATABASE_NAME"),
			COLLECTION_USERS:    os.Getenv("COLLECTION_USERS"),
			COLLECTION_PROFILES: os.Getenv("COLLECTION_PROFILES"),
			COLLECTION_MATCHES:  os.Getenv("COLLECTION_MATCHES"),
		},
	}
}

func createAppConfigTest() *AppConfig {
	return &AppConfig{
		ENV: ENV{
			TESTING_DATABASE_SOURCE: os.Getenv("TESTING_DATABASE_SOURCE"),
			TESTING_DATABASE_NAME:   os.Getenv("TESTING_DATABASE_NAME"),
			COLLECTION_USERS:        "test_users",
			COLLECTION_PROFILES:     "test_profiles",
			COLLECTION_MATCHES:      "test_matches",
		},
	}
}
