package helper

import (
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() {
	clientOpts := options.Client().ApplyURI(MONGO_URI)
	err := mgm.SetDefaultConfig(nil, "fetch_saldo_solana", clientOpts)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
}
