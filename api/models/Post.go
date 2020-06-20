package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title,omitempty"`
	Body      string             `bson:"body,omitempty"`
	Picture   string             `bson:"picture,omitempty"`
	UserID    primitive.ObjectID `bson:"userID,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty"`
	DeletedAt time.Time          `bson:"deletedAt,omitempty"`
}
