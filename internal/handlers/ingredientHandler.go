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

// Get all ingredients
func IngredientGetAll(w http.ResponseWriter, r *http.Request) {
	var data []m.Ingredient
	var responseCode int

	data, err := c.IngredientService.IngredientFindAll()
	if err != nil {
		response500WithDetails(w, err.Error())
	}

	responseCode = 200
	respondWithJSON(w, responseCode, data)
}

// Get a single ingredient
func IngredientGetSingle(w http.ResponseWriter, r *http.Request) {
	var data []m.Ingredient
	var responseCode int

	vars := mux.Vars(r)
	iID, err := strconv.Atoi(vars["ingredientID"])
	if err != nil {
		response500WithDetails(w, err.Error())
		return
	}

	data, err = c.IngredientService.IngredientFindSingle(iID)
	if err != nil {
		response500WithDetails(w, err.Error())
	}

	responseCode = 200
	respondWithJSON(w, responseCode, data)
}

// Get all ingredient sections
func SectionsGetAll(w http.ResponseWriter, r *http.Request) {
	var data []m.Section
	var responseCode int

	data, err := c.IngredientService.SectionsFindAll()
	if err != nil {
		response500WithDetails(w, err.Error())
	}

	responseCode = 200
	respondWithJSON(w, responseCode, data)
}

// Get a single ingredient section
func SectionsGetSingle(w http.ResponseWriter, r *http.Request) {
	var data []m.Section
	var responseCode int

	vars := mux.Vars(r)
	sID, err := strconv.Atoi(vars["sectionID"])
	if err != nil {
		response500WithDetails(w, err.Error())
		return
	}

	data, err = c.IngredientService.SectionsFindSingle(sID)
	if err != nil {
		response500WithDetails(w, err.Error())
	}

	responseCode = 200
	respondWithJSON(w, responseCode, data)
}

// Get ingredients belonging to a recipe
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
	if err != nil {
		response500WithDetails(w, err.Error())
	}

	responseCode = 200
	respondWithJSON(w, responseCode, data)
}

// Create an ingredient
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
