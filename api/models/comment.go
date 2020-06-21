package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Content   string             `bson:"content,omitempty"`
	UserID    primitive.ObjectID `bson:"userID,omitempty"`
	PostID    primitive.ObjectID `bson:"postID,omitempty"`
	CommentID primitive.ObjectID `bson:"commentID,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty"`
	DeletedAt time.Time          `bson:"deletedAt,omitempty"`
}
