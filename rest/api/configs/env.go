// package configs

// import (
// 	"log"
// 	"os"
// 	"strconv"

// 	"github.com/joho/godotenv"
// )

// func GetSettings() map[string]string {
// 	// Global Settings

// 	ignoreEnviroment, err := strconv.ParseBool(os.Getenv("IGNORE_ENVIRONMENT"))
// 	if !ignoreEnviroment {
// 		err = godotenv.Load()
// 	}

// 	if err != nil {
// 		log.Fatal("Error loading .env file: ", err.Error())
// 	}

// 	settings := make(map[string]string)

// 	settings["API_V1"] = "/api/v1"
// 	settings["PORT"] = os.Getenv("PORT")

// 	// // Mongo Settings
// 	// settings["MONGO_URI"] = os.Getenv("MONGO_URI")
// 	// settings["MONGO_DATABASE"] = os.Getenv("MONGO_DATABASE")

// 	// // Auth0 Settings
// 	// settings["JWT_ISSUER"] = os.Getenv("JWT_ISSUER")
// 	// settings["JWT_ALGORITHM"] = os.Getenv("JWT_ALGORITHM")
// 	// settings["JWT_AUDIENCE"] = os.Getenv("JWT_AUDIENCE")
// 	// settings["JWT_URL_RSA_KEY"] = os.Getenv("JWT_URL_RSA_KEY")

// 	// // GCP
// 	// settings["BUCKET_NAME"] = os.Getenv("BUCKET_NAME")
// 	// settings["BUCKET_SOP_SERIES_SIMULATION"] = os.Getenv("BUCKET_SOP_SERIES_SIMULATION")
// 	// settings["GCP_SERVICE_ACCOUNT"] = os.Getenv("GCP_SERVICE_ACCOUNT")

// 	// // Model API
// 	// settings["MODEL_API_URL"] = os.Getenv("MODEL_API_URL")

// 	// // FeatureStore API
// 	// settings["FS_API_URL"] = os.Getenv("FS_API_URL")

// 	// // PUB/SUB
// 	// settings["GCP_JNJ_SOP_TOPIC_ID"] = os.Getenv("GCP_JNJ_SOP_TOPIC_ID")
// 	// settings["GCP_PROJECT_ID"] = os.Getenv("GCP_PROJECT_ID")

// 	return settings
// }
