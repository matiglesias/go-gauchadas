package services

import (
	"context"

	"github.com/gauchadas/api/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostsService struct {
	postsCollection *mongo.Collection
}

func NewPostsService(postsCollection *mongo.Collection) *PostsService {
	return &PostsService{
		postsCollection: postsCollection}
}

func (ps *PostsService) Create(p *models.Post) (interface{}, error) {
	insertResult, err := ps.postsCollection.InsertOne(context.TODO(), p)
	if err != nil {
		return nil, err
	}

	/* 	ctx := context.TODO()
	   	cursor, err := ps.postsCollection.Find(ctx, bson.D{{}})
	   	if err != nil {
	   		return nil, err
	   	}
	   	defer cursor.Close(ctx)

	   	var results []models.Post
	   	err = cursor.All(ctx, &results)
	   	if err != nil {
	   		return nil, err
	   	}
	   		for _, result := range results {
	   		fmt.Println("EL BODY ES " + result.Body)
	   		fmt.Println("EL TITULO ES " + result.Title)
	   		fmt.Println("LA HORA ES " + result.CreatedAt.Time().String())
	   		fmt.Println("")
	   	} */
	return insertResult.InsertedID, nil
}
