package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Content   string             `bson:"content,omitempty" json:"content,omitempty"`
	UserID    primitive.ObjectID `bson:"userID,omitempty" json:"userID,omitempty"`
	PostID    primitive.ObjectID `bson:"postID,omitempty" json:"postID,omitempty"`
	CommentID primitive.ObjectID `bson:"commentID,omitempty" json:"commentID,omitempty"`
	CreatedAt primitive.DateTime `bson:"createdAt,omitempty" json:"createdAt,string,omitempty"`
	UpdatedAt primitive.DateTime `bson:"updatedAt,omitempty" json:"updatedAt,string,omitempty"`
	DeletedAt primitive.DateTime `bson:"deletedAt,omitempty" json:"deletedAt,string,omitempty"`
}
