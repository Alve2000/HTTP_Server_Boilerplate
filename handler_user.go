
package main

import (
	"fmt"
	"time"
	"net/http"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/Alve2000/rssagg/internal/database"
)

// this function handles requests to create a user
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v",err))
		return
	}
	
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: 	   uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: 	   params.Name,
	})
	
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v",err))
		return
	}
	
	respondWithJSON(w, 201, databaseUserToUser(user))
}

// this function handles requests to get a user
func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, databaseUserToUser(user))
}

// this function handles requests to get all users
func (apiCfg *apiConfig) handlerGetAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := apiCfg.DB.GetAllUsers(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get all users: %v",err))
		return
	}

	respondWithJSON(w, 200, databaseUsersToUsers(users))
}

// this function handles requests for a user to get new posts from feeds following
func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID:	user.ID,
		Limit:	10,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get posts: %v",err))
		return
	}

	respondWithJSON(w, 200, databasePostsToPosts(posts))
}
