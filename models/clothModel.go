package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NEED  TO give firs letter as uppercase here
type Cloth struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       *string            `json:"name" validate:"required,min=2,max=100"`
	Price      *float64           `json:"price" validate:"required"`
	Cloth_image *string            `json:"cloth_image" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Cloth_id    string             `json:"cloth_id"`
	Inventory_id    *string            `json:"inventory_id" validate:"required"`
}