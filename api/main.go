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
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	}

	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_HOST")))
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	}
	defer client.Disconnect(ctx)

	database := client.Database("gauchadas")
	postsCollection := database.Collection("posts")

	postsService := services.NewPostsService(postsCollection)
	postsController := controllers.NewPostsController(postsService)

	router := mux.NewRouter()

	//Posts routes
	router.HandleFunc("/api/posts", middlewares.JSON(postsController.GetAll)).Methods("GET")
	router.HandleFunc("/api/posts", middlewares.JSON(postsController.Create)).Methods("POST")

	router.HandleFunc("/api/posts/{id}", middlewares.JSON(postsController.GetByID)).Methods("GET")
	router.HandleFunc("/api/posts/{id}", middlewares.JSON(postsController.Edit)).Methods("PUT")
	router.HandleFunc("/api/posts/{id}", middlewares.JSON(postsController.Delete)).Methods("DELETE")

	router.HandleFunc("/api/posts/{id}/comments", middlewares.JSON(postsController.GetComments)).Methods("GET")
	router.HandleFunc("/api/posts/{id}/comments", middlewares.JSON(postsController.CreateMainComment)).Methods("POST")

	router.HandleFunc("/api/posts/{id}/comments/{commentID}", middlewares.JSON(postsController.CreateSecondaryComment)).Methods("POST")
	router.HandleFunc("/api/posts/{id}/comments/{commentID}", middlewares.JSON(postsController.EditComment)).Methods("PUT")
	router.HandleFunc("/api/posts/{id}/comments/{commentID}", middlewares.JSON(postsController.DeleteComment)).Methods("DELETE")

	log.Printf("Listening server at %s...\n", os.Getenv("SERVER_URL"))
	http.ListenAndServe(os.Getenv("SERVER_URL"), router)
}
