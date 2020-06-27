package controllers

import (
	"net/http"
	"time"

	"github.com/gauchadas/api/models"
	"github.com/gauchadas/api/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostsController struct {
	postsService *services.PostsService
}

func NewPostsController(postsService *services.PostsService) *PostsController {
	return &PostsController{
		postsService: postsService}
}

func (pc *PostsController) GetAll(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return nil, nil
}

func (pc *PostsController) Create(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var p models.Post
	// Memory read limit: 1MB
	err := r.ParseMultipartForm(1048576)
	if err != nil {
		return nil, err
	}

	p.Title = r.FormValue("title")
	p.Body = r.FormValue("body")
	// TODO get picture
	// TODO get userID
	p.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

	return pc.postsService.Create(&p)
}

func (pc *PostsController) GetByID(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return nil, nil
}

func (pc *PostsController) Edit(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return nil, nil
}

func (pc *PostsController) Delete(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return nil, nil
}

func (pc *PostsController) GetComments(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return nil, nil
}

func (pc *PostsController) CreateMainComment(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return nil, nil
}

func (pc *PostsController) CreateSecondaryComment(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return nil, nil
}

func (pc *PostsController) EditComment(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return nil, nil
}

func (pc *PostsController) DeleteComment(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return nil, nil
}
