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

func (pc *PostsController) GetAll(w http.ResponseWriter, r *http.Request) {
}

func (pc *PostsController) Create(w http.ResponseWriter, r *http.Request) {
}

func (pc *PostsController) GetByID(w http.ResponseWriter, r *http.Request) {
}

func (pc *PostsController) Edit(w http.ResponseWriter, r *http.Request) {
}

func (pc *PostsController) Delete(w http.ResponseWriter, r *http.Request) {
}

func (pc *PostsController) GetPostsComments(w http.ResponseWriter, r *http.Request) {
}

func (pc *PostsController) CreateMainPostsComment(w http.ResponseWriter, r *http.Request) {
}

func (pc *PostsController) CreateSecondaryPostsComment(w http.ResponseWriter, r *http.Request) {
}

func (pc *PostsController) EditPostComment(w http.ResponseWriter, r *http.Request) {
}

func (pc *PostsController) DeletePostComment(w http.ResponseWriter, r *http.Request) {
}
