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

func (ps *PostsService) Edit(data []byte, postID string) (interface{}, error) {
	post, err := ps.postExists(postID)
	if err != nil {
		return nil, err
	}

	var p models.Post
	err = json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}

	update := bson.D{
		{"$set", p},
		{"$currentDate", bson.D{
			{"updatedAt", true},
		}},
	}
	filter := bson.D{
		{"_id", post.ID},
	}
	updateResult, err := ps.postsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return updateResult, nil
}

func (ps *PostsService) GetComments(postID string) (interface{}, error) {
	post, err := ps.postExists(postID)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	filter := bson.D{
		{"postID", post.ID},
		{"deletedAt", bson.D{
			{"$exists", false},
		}},
	}

	cursor, err := ps.commentsCollection.Find(ctx, filter)
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

func (ps *PostsService) CreateSecondaryComment(data []byte, postID string, mainCommentID string) (interface{}, error) {
	// Check if post exists
	post, err := ps.postExists(postID)
	if err != nil {
		return nil, err
	}

	// Check if main comment exists
	mainComment, err := ps.commentExists(mainCommentID)
	if err != nil {
		return nil, err
	}
	if !mainComment.CommentID.IsZero() {
		return nil, errors.New("Not a main comment")
	}

	// Fill model and store secondary comment
	var c models.Comment
	err = json.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}
	// ¡OJO! PODRIAN MANDAR JASON CON deletedAt U OTRO CAMPO, Y HARIA CAGADAS.
	c.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	c.PostID = post.ID
	c.CommentID = mainComment.ID

	insertResult, err := ps.commentsCollection.InsertOne(context.TODO(), c)
	if err != nil {
		return nil, err
	}
	return insertResult, nil
}

func (ps *PostsService) EditComment(data []byte, postID string, commentID string) (interface{}, error) {
	_, err := ps.postExists(postID)
	if err != nil {
		return nil, err
	}

	comment, err := ps.commentExists(commentID)
	if err != nil {
		return nil, err
	}

	var c models.Comment
	err = json.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}

	update := bson.D{
		{"$set", c},
		{"$currentDate", bson.D{
			{"updatedAt", true},
		}},
	}
	filter := bson.D{
		{"_id", comment.ID},
	}
	updateResult, err := ps.commentsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return updateResult, nil
}

func (ps *PostsService) DeleteComment(postID string, commentID string) (interface{}, error) {
	_, err := ps.postExists(postID)
	if err != nil {
		return nil, err
	}

	cID, err := primitive.ObjectIDFromHex(commentID)
	filter := bson.D{
		{"_id", cID},
	}
	deleteByKeyResult, err := ps.commentsCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	if deleteByKeyResult.DeletedCount == 0 {
		return nil, errors.New("Comment not found.")
	}

	filter = bson.D{
		{"commentID", cID},
	}
	deleteByForeingKeyResult, err := ps.commentsCollection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	deleteResults := bson.D{
		{"deletedByItsKey", deleteByKeyResult},
		{"deletedByForeingKey", deleteByForeingKeyResult},
	}
	return deleteResults, nil
}

// If post exists, returns postId as an objectID.
func (ps *PostsService) postExists(postID string) (*models.Post, error) {
	pID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return nil, err
	}
	var post models.Post
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
	return &post, nil
}

// If comment exists, returns comment as a Comment model.
func (ps *PostsService) commentExists(commentID string) (*models.Comment, error) {
	cID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return nil, err
	}
	var comment models.Comment
	filter := bson.D{
		{"_id", cID},
		{"deletedAt", bson.D{
			{"$exists", false},
		}},
	}
	err = ps.commentsCollection.FindOne(context.TODO(), filter).Decode(&comment)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}
