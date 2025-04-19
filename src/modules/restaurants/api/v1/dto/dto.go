package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateRestaurantReq struct {
	Name       string `json:"name" validate:"required"`
	LocationID string `json:"location_id"  validate:"required"`
	Email      string `json:"email" validate:"required,email"`
}

type CreateRestaurantRes struct {
	ID       primitive.ObjectID `json:"_id"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	APIKey   string             `json:"apikey"`
}
