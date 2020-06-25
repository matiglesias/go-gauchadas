package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Content   string             `bson:"content,omitempty"`
	UserID    primitive.ObjectID `bson:"userID,omitempty"`
	PostID    primitive.ObjectID `bson:"postID,omitempty"`
	CommentID primitive.ObjectID `bson:"commentID,omitempty"`
	CreatedAt primitive.DateTime `bson:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime `bson:"updatedAt,omitempty"`
	DeletedAt primitive.DateTime `bson:"deletedAt,omitempty"`
}
