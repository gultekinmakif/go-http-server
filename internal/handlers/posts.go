package handlers

import (
	"net/http"
)

// createPostRequest is the JSON body for POST /posts.
type createPostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// CreatePost handles POST /posts.
// Body: {"title":"...","body":"..."}.
// 201 with the created post on success, 400 on bad body, 500 on db error.
func CreatePost(w http.ResponseWriter, r *http.Request) {
}

// GetPost handles GET /posts/{slug}.
// 200 + post on hit, 404 on miss, 500 on db error.
func GetPost(w http.ResponseWriter, r *http.Request) {
}

// GetPost handles GET /posts.
// 200 + post on hit, 404 on miss, 500 on db error.
func GetAllPost(w http.ResponseWriter, r *http.Request) {
}
