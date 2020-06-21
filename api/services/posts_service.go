package services

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostsService struct {
	postsCollection *mongo.Collection
}

func NewPostsService(postsCollection *mongo.Collection) *PostsService {
	return &PostsService{
		postsCollection: postsCollection}
}

func (ps *PostsService) GetAll() (interface{}, error) {
	ctx := context.TODO()
	filter := bson.D{
		{"deletedAt", bson.D{
			{"$exists", false},
		}},
	}
	cursor, err := ps.postsCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
