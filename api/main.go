package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gauchadas/api/controllers"
	"github.com/gauchadas/api/middlewares"
	"github.com/gauchadas/api/services"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_HOST")))
	if err != nil {
		log.Fatalf("Error connecting to database, not coming through %v", err)
	}
	defer client.Disconnect(ctx)

	database := client.Database("gauchadas")
	postsCollection := database.Collection("posts")
	commentsCollection := database.Collection("comments")

	postsService := services.NewPostsService(postsCollection, commentsCollection)
	postsController := controllers.NewPostsController(postsService)

	router := mux.NewRouter()

	//Posts routes
	router.HandleFunc("/api/posts", middlewares.JSON(postsController.GetAll)).Methods("GET")
	router.HandleFunc("/api/posts", middlewares.JSON(postsController.Create)).Methods("POST", "OPTIONS")

	router.HandleFunc("/api/posts/{id}", middlewares.JSON(postsController.GetByID)).Methods("GET")
	router.HandleFunc("/api/posts/{id}", middlewares.JSON(postsController.Edit)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/posts/{id}", middlewares.JSON(postsController.Delete)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/posts/{id}/comments", middlewares.JSON(postsController.GetComments)).Methods("GET")
	router.HandleFunc("/api/posts/{id}/comments", middlewares.JSON(postsController.CreateMainComment)).Methods("POST", "OPTIONS")

	router.HandleFunc("/api/posts/{id}/comments/{commentID}", middlewares.JSON(postsController.GetCommentResponses)).Methods("GET")
	router.HandleFunc("/api/posts/{id}/comments/{commentID}", middlewares.JSON(postsController.CreateSecondaryComment)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/posts/{id}/comments/{commentID}", middlewares.JSON(postsController.EditComment)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/posts/{id}/comments/{commentID}", middlewares.JSON(postsController.DeleteComment)).Methods("DELETE", "OPTIONS")

	router.Use(mux.CORSMethodMiddleware(router))

	log.Printf("Listening server at *:8080...\n")
	http.ListenAndServe(":8080", router)
}
