package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title,omitempty"`
	Body      string             `bson:"body,omitempty"`
	Picture   primitive.Binary   `bson:"picture,omitempty"`
	UserID    primitive.ObjectID `bson:"userID,omitempty"`
	CreatedAt primitive.DateTime `bson:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime `bson:"updatedAt,omitempty"`
	DeletedAt primitive.DateTime `bson:"deletedAt,omitempty"`
}
