package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gauchadas/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostsService struct {
	postsCollection *mongo.Collection
}

func NewPostsService(postsCollection *mongo.Collection) *PostsService {
	return &PostsService{
		postsCollection: postsCollection}
}

func (ps *PostsService) Create(data []byte) (interface{}, error) {
	var p models.Post
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	p.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

	insertResult, err := ps.postsCollection.InsertOne(context.TODO(), p)
	if err != nil {
		return nil, err
	}
	return insertResult, nil
}
