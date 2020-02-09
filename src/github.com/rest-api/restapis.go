package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rest-api/model"
)

var posts []model.Post

func main() {
	router := mux.NewRouter()

	post1 := model.Post{ID: "1", PostName: "first", PostDescription: "first Description"}
	posts = append(posts, post1)

	post2 := model.Post{ID: "2", PostName: "second", PostDescription: "second Description"}
	posts = append(posts, post2)

	router.HandleFunc("/posts", getPosts).Methods(http.MethodGet)
	router.HandleFunc("/posts/{id}", getPost).Methods(http.MethodGet)
	router.HandleFunc("/posts", createPost).Methods(http.MethodPost)

	serverError := http.ListenAndServe(":8000", router)

	if serverError != nil {
		fmt.Println("server error " + serverError.Error())
	}
}

func getPost(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for _, p := range posts {
		if p.ID == params["id"] {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	json.NewEncoder(w).Encode(model.Post{})
}

func getPosts(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func createPost(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := model.Post{}
	json.NewDecoder(request.Body).Decode(&post)
	posts = append(posts, post)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}
