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

func IngredientGetAll(w http.ResponseWriter, r *http.Request) {
	var data []m.Ingredient
	var responseCode int

	data, err := c.IngredientService.FindAll()
	if err != nil {
		response500WithDetails(w, err.Error())
	}

	responseCode = 200
	respondWithJSON(w, responseCode, data)
}

func RecipeIngredientGet(w http.ResponseWriter, r *http.Request) {
	var data []m.IngredientDTO
	var responseCode int

	vars := mux.Vars(r)
	rID, err := strconv.Atoi(vars["recipeID"])
	if err != nil {
		response500WithDetails(w, err.Error())
		return
	}

	data, err = c.IngredientService.FindRecipeIngredients(rID)

	responseCode = 200
	respondWithJSON(w, responseCode, data)
}

func IngredientCreate(w http.ResponseWriter, r *http.Request) {
	var ingredient m.Ingredient
	var data m.Ingredient

	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(r.Body)
	if err != nil {
		response500WithDetails(w, err.Error())
		return
	}

	body := buffer.String()

	if err = json.Unmarshal([]byte(body), &ingredient); err != nil {
		response500WithDetails(w, err.Error())
		return
	}

	data, err = c.IngredientService.Create(ingredient)
	if err != nil {
		response500WithDetails(w, err.Error())
	}

	respondWithJSON(w, 201, data)
}
