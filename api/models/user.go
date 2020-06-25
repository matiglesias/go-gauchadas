package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Surname   string             `bson:"surname,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Password  string             `bson:"password,omitempty"`
	CreatedAt primitive.DateTime `bson:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime `bson:"updatedAt,omitempty"`
	DeletedAt primitive.DateTime `bson:"deletedAt,omitempty"`
}
