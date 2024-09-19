package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ZhanabekZ/rss_aggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handler_CreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json: "name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().Local().UTC(),
		UpdatedAt: time.Now().Local().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't create the user: %v", err))
		return
	}
	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handler_GetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, databaseUserToUser(user))
}
