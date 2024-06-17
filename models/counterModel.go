package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Counter struct {
	ID               primitive.ObjectID `bson:"_id"`
	Number_of_guests *int               `json:"number_of_guests" validate:"required"`
	Counter_number     *int               `json:"counter_number" validate:"required"`
	Created_at       time.Time          `json:"created_at"`
	Updated_at       time.Time          `json:"updated_at"`
	Counter_id         string             `json:"counter_id"`
}