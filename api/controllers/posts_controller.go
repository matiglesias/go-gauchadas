package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/gauchadas/api/services"
	"github.com/gorilla/mux"
)

type PostsController struct {
	postsService *services.PostsService
}

func NewPostsController(postsService *services.PostsService) *PostsController {
	return &PostsController{
		postsService: postsService}
}

func (pc *PostsController) GetAll(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return pc.postsService.GetAll()
}

func (pc *PostsController) Create(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return pc.postsService.Create(rBody)
}

func (pc *PostsController) GetByID(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	postID := mux.Vars(r)["id"]
	return pc.postsService.GetByID(postID)
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
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	postID := mux.Vars(r)["id"]
	return pc.postsService.CreateMainComment(rBody, postID)
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
