package controllers

import (
	"net/http"

	"github.com/gauchadas/api/services"
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
	return nil, nil
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
