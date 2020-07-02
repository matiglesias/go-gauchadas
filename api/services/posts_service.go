package services

import (
	"context"
	"encoding/json"
	"time"

	"errors"

	"github.com/gauchadas/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostsService struct {
	postsCollection    *mongo.Collection
	commentsCollection *mongo.Collection
}

func NewPostsService(postsCollection *mongo.Collection, commentsCollection *mongo.Collection) *PostsService {
	return &PostsService{
		postsCollection:    postsCollection,
		commentsCollection: commentsCollection}
}

func (ps *PostsService) CreateMainComment(data []byte, postID string) (interface{}, error) {
	pID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return nil, err
	}
	var post bson.M
	filter := bson.D{
		{"_id", pID},
		{"deletedAt", bson.D{
			{"$exists", false},
		}},
	}
	err = ps.postsCollection.FindOne(context.TODO(), filter).Decode(&post)
	if err != nil {
		return nil, err
	}

	var c models.Comment
	err = json.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}
	/* 	// ¡OJO! PODRIAN MANDAR JASON CON deletedAt U OTRO CAMPO, Y HARIA CAGADAS. WENO
	   	// AUNQ SERIA MAS IMPORTANTE EN EDITAR COMENTARIO JEP
	   	zeroTime := time.Unix(0, 0)
	   	if  !c.DeletedAt.Time().Equal(zeroTime) {
	   		return
	   	} */
	c.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	c.PostID = pID

	insertResult, err := ps.commentsCollection.InsertOne(context.TODO(), c)
	if err != nil {
		return nil, err
	}
	return insertResult, nil
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

func (ps *PostsService) GetByID(postID string) (interface{}, error) {
	id, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return nil, errors.New("Identificador de post invalido.")
	}

	filter := bson.D{
		{"_id", id},
		{"deletedAt", bson.D{
			{"$exists", false},
		}},
	}
	var post bson.M
	err = ps.postsCollection.FindOne(context.TODO(), filter).Decode(&post)
	if err != nil {
		return nil, err
	}
	// Must return Picture too
	return post, nil
}
