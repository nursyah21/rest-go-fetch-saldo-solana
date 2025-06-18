package helper

import "go.mongodb.org/mongo-driver/mongo"

func GetCollection(collectionName string) *mongo.Collection {
	return MongoClient.Database("fetch_saldo_solana").Collection(collectionName)
}
