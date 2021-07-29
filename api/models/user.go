package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	Surname   string             `bson:"surname,omitempty" json:"surname,omitempty"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
	Password  string             `bson:"password,omitempty" json:"password,omitempty"`
	CreatedAt primitive.DateTime `bson:"createdAt,omitempty" json:"createdAt,string,omitempty"`
	UpdatedAt primitive.DateTime `bson:"updatedAt,omitempty" json:"updatedAt,string,omitempty"`
	DeletedAt primitive.DateTime `bson:"deletedAt,omitempty" json:"deletedAt,string,omitempty"`
}
