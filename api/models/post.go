package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title     string             `bson:"title,omitempty" json:"title,omitempty"`
	Body      string             `bson:"body,omitempty" json:"body,omitempty"`
	Picture   primitive.Binary   `bson:"picture,omitempty" json:"picture,omitempty"`
	UserID    primitive.ObjectID `bson:"userID,omitempty" json:"userID,string,omitempty"`
	CreatedAt primitive.DateTime `bson:"createdAt,omitempty" json:"createdAt,string,omitempty"`
	UpdatedAt primitive.DateTime `bson:"updatedAt,omitempty" json:"updatedAt,string,omitempty"`
	DeletedAt primitive.DateTime `bson:"deletedAt,omitempty" json:"deletedAt,string,omitempty"`
}
