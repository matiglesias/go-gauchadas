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

// Cambiar todos los bson.D a bson.M, ya que no me importa el orden de los campos.
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
	if p.DeletedAt != 0 || p.UpdatedAt != 0 {
		return nil, errors.New("request have an unespected date field")
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
		return nil, errors.New("identificador de post invalido")
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
	if p.DeletedAt != 0 || p.UpdatedAt != 0 || p.CreatedAt != 0 {
		return nil, errors.New("request have an unespected date field")
	}

	update := bson.D{
		{"$set", p},
		{"$currentDate", bson.D{
			{"updatedAt", true},
		}},
	}
	filter := bson.D{{"_id", post.ID}}
	updateResult, err := ps.postsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return updateResult, nil
}

func (ps *PostsService) Delete(postID string) (interface{}, error) {
	post, err := ps.postExists(postID)
	if err != nil {
		return nil, err
	}

	update := bson.D{
		{"$currentDate", bson.D{
			{"deletedAt", true},
		}},
	}

	filter := bson.D{{"postID", post.ID}}
	_, err = ps.commentsCollection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	filter = bson.D{{"_id", post.ID}}
	_, err = ps.postsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return "true", nil
}

func (ps *PostsService) Restore(postID string) (interface{}, error) {
	post, err := ps.postDeleted(postID)
	if err != nil {
		return nil, err
	}

	update := bson.D{
		{"$unset", bson.D{
			{"deletedAt", ""},
		}},
	}

	/* TODO: Post owner shouldnt be able to restore soft-deleted comments by comments owners. */
	filter := bson.D{{"postID", post.ID}}
	_, err = ps.commentsCollection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	filter = bson.D{{"_id", post.ID}}
	_, err = ps.postsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return "true", nil
}

func (ps *PostsService) GetComments(postID string) (interface{}, error) {
	post, err := ps.postExists(postID)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	filter := bson.D{
		{"postID", post.ID},
		{"deletedAt", bson.D{{"$exists", false}}},
		{"commentID", bson.D{{"$exists", false}}},
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
	post, err := ps.postExists(postID)
	if err != nil {
		return nil, err
	}
	var c models.Comment
	err = json.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}
	if c.DeletedAt != 0 || c.UpdatedAt != 0 {
		return nil, errors.New("request have an unespected date field")
	}

	c.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	c.PostID = post.ID
	insertResult, err := ps.commentsCollection.InsertOne(context.TODO(), c)
	if err != nil {
		return nil, err
	}

	createdComment := bson.M{
		"_id":       insertResult.InsertedID.(primitive.ObjectID),
		"content":   c.Content,
		"postID":    c.PostID,
		"createdAt": c.CreatedAt,
	}
	return createdComment, nil
}

func (ps *PostsService) GetCommentResponses(postID string, commentID string) (interface{}, error) {
	_, err := ps.postExists(postID)
	if err != nil {
		return nil, err
	}

	comment, err := ps.commentExists(commentID)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	filter := bson.D{
		{"commentID", comment.ID},
		{"deletedAt", bson.D{{"$exists", false}}},
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

func (ps *PostsService) CreateSecondaryComment(data []byte, postID string, mainCommentID string) (interface{}, error) {
	post, err := ps.postExists(postID)
	if err != nil {
		return nil, err
	}
	mainComment, err := ps.commentExists(mainCommentID)
	if err != nil {
		return nil, err
	}
	if !mainComment.CommentID.IsZero() {
		return nil, errors.New("not a main comment")
	}

	var c models.Comment
	err = json.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}
	if c.DeletedAt != 0 || c.UpdatedAt != 0 {
		return nil, errors.New("request have an unespected date field")
	}

	c.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	c.PostID = post.ID
	c.CommentID = mainComment.ID
	insertResult, err := ps.commentsCollection.InsertOne(context.TODO(), c)
	if err != nil {
		return nil, err
	}

	createdResponse := bson.M{
		"_id":       insertResult.InsertedID.(primitive.ObjectID),
		"content":   c.Content,
		"commentID": c.CommentID,
		"postID":    c.PostID,
		"createdAt": c.CreatedAt,
	}
	return createdResponse, nil
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
	if c.DeletedAt != 0 || c.UpdatedAt != 0 || c.CreatedAt != 0 || !c.CommentID.IsZero() || !c.PostID.IsZero() {
		return nil, errors.New("request have an unespected field")
	}

	update := bson.D{
		{"$set", c},
		{"$currentDate", bson.D{
			{"updatedAt", true},
		}},
	}
	filter := bson.D{{"_id", comment.ID}}
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
	if err != nil {
		return nil, err
	}

	update := bson.D{
		{"$currentDate", bson.D{
			{"deletedAt", true}},
		},
	}
	filter := bson.D{{"_id", cID}}
	deleteByKeyResult, err := ps.commentsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	if deleteByKeyResult.MatchedCount == 0 {
		return nil, errors.New("comment not found")
	}

	filter = bson.D{{"commentID", cID}}
	deleteByForeingKeyResult, err := ps.commentsCollection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	deleteResults := bson.M{
		"deleteByKeyResult":        deleteByKeyResult,
		"deleteByForeingKeyResult": deleteByForeingKeyResult,
	}
	return deleteResults, nil
}

// If post exists, returns post as a Post struct.
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

// If post have been soft-deleted, returns post as a Post struct.
func (ps *PostsService) postDeleted(postID string) (*models.Post, error) {
	pID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return nil, err
	}

	var post models.Post
	filter := bson.D{
		{"_id", pID},
		{"deletedAt", bson.D{
			{"$exists", true},
		}},
	}
	err = ps.postsCollection.FindOne(context.TODO(), filter).Decode(&post)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// If comment exists, returns comment as a Comment struct.
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
