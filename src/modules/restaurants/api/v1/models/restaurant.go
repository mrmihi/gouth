package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goose/src/database"
	"time"
)

type Restaurant struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name,omitempty"`
	Email      string             `json:"email" bson:"email,omitempty"`
	Password   string             `json:"password" bson:"password,omitempty"`
	APIKey     string             `json:"apikey" bson:"apikey,omitempty"`
	LocationID string             `json:"location_id" bson:"location_id,omitempty"`
	CreatedAt  string             `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt  string             `json:"updated_at" bson:"updated_at,omitempty"`
}

func (r Restaurant) WithDefaults() Restaurant {
	r.CreatedAt = time.Now().Format(time.RFC3339)
	r.UpdatedAt = time.Now().Format(time.RFC3339)
	return r
}

func (r Restaurant) Secure() Restaurant {
	r.Password = ""
	return r
}

func SyncIndexes() {
	database.UseDefault().Collection("restaurants").Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: -1}},
		Options: options.Index().SetUnique(true),
	})
}
