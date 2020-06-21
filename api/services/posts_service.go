package services

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type PostsService struct {
	postsCollection *mongo.Collection
}

func NewPostsService(postsCollection *mongo.Collection) *PostsService {
	return &PostsService{
		postsCollection: postsCollection}
}
