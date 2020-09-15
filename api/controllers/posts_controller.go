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

func (pc *PostsController) GetAll(r *http.Request) (interface{}, error) {
	return pc.postsService.GetAll()
}

func (pc *PostsController) Create(r *http.Request) (interface{}, error) {
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return pc.postsService.Create(rBody)
}

func (pc *PostsController) GetByID(r *http.Request) (interface{}, error) {
	postID := mux.Vars(r)["id"]
	return pc.postsService.GetByID(postID)
}

func (pc *PostsController) Edit(r *http.Request) (interface{}, error) {
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	postID := mux.Vars(r)["id"]
	return pc.postsService.Edit(rBody, postID)
}

func (pc *PostsController) Delete(r *http.Request) (interface{}, error) {
	postID := mux.Vars(r)["id"]
	return pc.postsService.Delete(postID)
}

func (pc *PostsController) GetComments(r *http.Request) (interface{}, error) {
	postID := mux.Vars(r)["id"]
	return pc.postsService.GetComments(postID)
}

func (pc *PostsController) CreateMainComment(r *http.Request) (interface{}, error) {
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	postID := mux.Vars(r)["id"]
	return pc.postsService.CreateMainComment(rBody, postID)
}

func (pc *PostsController) GetCommentResponses(r *http.Request) (interface{}, error) {
	postID := mux.Vars(r)["id"]
	commentID := mux.Vars(r)["commentID"]
	return pc.postsService.GetCommentResponses(postID, commentID)
}

func (pc *PostsController) CreateSecondaryComment(r *http.Request) (interface{}, error) {
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	postID := mux.Vars(r)["id"]
	mainCommentID := mux.Vars(r)["commentID"]
	return pc.postsService.CreateSecondaryComment(rBody, postID, mainCommentID)
}

func (pc *PostsController) EditComment(r *http.Request) (interface{}, error) {
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	postID := mux.Vars(r)["id"]
	commentID := mux.Vars(r)["commentID"]
	return pc.postsService.EditComment(rBody, postID, commentID)
}

func (pc *PostsController) DeleteComment(r *http.Request) (interface{}, error) {
	postID := mux.Vars(r)["id"]
	commentID := mux.Vars(r)["commentID"]
	return pc.postsService.DeleteComment(postID, commentID)
}
