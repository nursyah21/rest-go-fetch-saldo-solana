package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	RPC_URI    string
	MONGO_URI  string
	SECRET_KEY string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system env vars")
	}

	RPC_URI = os.Getenv("RPC_URI")
	if RPC_URI == "" {
		log.Fatal("RPC_URI is not set in environment or .env file")
	}

	MONGO_URI = os.Getenv("MONGO_URI")
	if MONGO_URI == "" {
		log.Fatal("MONGO_URI is not set in environment or .env file")
	}

	SECRET_KEY = os.Getenv("SECRET_KEY")
	if SECRET_KEY == "" {
		log.Fatal("SECRET_KEY is not set in environment or .env file")
	}
}
