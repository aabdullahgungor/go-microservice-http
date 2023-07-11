package product

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Brand struct {
	Id   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string             `bson:"name,omitempty" json:"name"`
}

type Product struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string             `bson:"name,omitempty" json:"name"`
	Price      float64            `bson:"price,omitempty" json:"price"`
	Quantity   int64              `bson:"quantity,omitempty" json:"quantity"`
	Status     bool               `bson:"status,omitempty" json:"status"`
	Date       time.Time          `bson:"date,omitempty" json:"date"`
	CategoryId primitive.ObjectID `bson:"categoryId,omitempty" json:"category_id"`
	Brand      Brand              `bson:"brand,omitempty" json:"brand"`
	Colors     []string           `bson:"colors,omitempty" json:"colors"`
}
