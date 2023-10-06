package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetSettings() map[string]string {
	// Global Settings

	ignoreEnviroment, err := strconv.ParseBool(os.Getenv("IGNORE_ENVIRONMENT"))
	if !ignoreEnviroment {
		err = godotenv.Load()
	}

	if err != nil {
		log.Fatal("Error loading .env file: ", err.Error())
	}

	settings := make(map[string]string)

	settings["POSTGRES"] = os.Getenv("POSTGRES")

	return settings
}
