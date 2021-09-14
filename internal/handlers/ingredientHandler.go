package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

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
