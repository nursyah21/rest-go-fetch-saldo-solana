package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	RPC_URI   string
	MONGO_URI string
)

func NewConfig() {
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

	clientOpts := options.Client().ApplyURI(MONGO_URI)
	err = mgm.SetDefaultConfig(nil, "fetch_saldo_solana", clientOpts)
	if err != nil {
		log.Fatal(err.Error())
	}

}
