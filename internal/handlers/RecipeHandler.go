package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	c "github.com/ihulsbus/cookbook/internal/config"
	m "github.com/ihulsbus/cookbook/internal/models"
)

func RecipeGetAll(w http.ResponseWriter, r *http.Request) {
	var data []m.RecipeDTO
	var responseCode int

	data, err := c.RecipeService.FindAll()
	if err != nil {
		response500WithDetails(w, err.Error())
	}

	responseCode = 200
	respondWithJSON(w, responseCode, data)
}

func RecipeGet(w http.ResponseWriter, r *http.Request) {
	var data m.RecipeDTO
	var responseCode int

	vars := mux.Vars(r)
	rID, err := strconv.Atoi(vars["recipeID"])
	if err != nil {
		response500WithDetails(w, err.Error())
	}

	data, err = c.RecipeService.Find(rID)
	if err != nil {
		response500WithDetails(w, err.Error())
	}

	responseCode = 200
	respondWithJSON(w, responseCode, data)
}

func RecipeCreate(w http.ResponseWriter, r *http.Request) {
	var recipe m.RecipeDTO
	var data m.RecipeDTO

	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(r.Body)
	if err != nil {
		response500WithDetails(w, err.Error())
		return
	}

	body := buffer.String()

	if err = json.Unmarshal([]byte(body), &recipe); err != nil {
		response500WithDetails(w, err.Error())
		return
	}

	data, err = c.RecipeService.Create(recipe)
	if err != nil {
		response500WithDetails(w, err.Error())
	}

	respondWithJSON(w, 201, data)
}
