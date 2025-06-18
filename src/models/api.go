package models

import (
	"context"
	"fetch-saldo/src/helper"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type API struct {
	ID        string             `json:"id,omitempty" bson:"_id,omitempty"`
	Api       string             `json:"api" bson:"api"`
	CreatedAt primitive.DateTime `json:"created_at"`
}

func CreateAPI(apiKey string) error {
	api := API{
		Api:       apiKey,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	_, err := helper.GetCollection("apis").InsertOne(context.Background(), api)
	return err
}

func ApiExist(apiKey string) bool {
	var result API
	err := helper.GetCollection("apis").FindOne(context.Background(), bson.M{"api": apiKey}).Decode(&result)

	return err == nil
}
