package services

import (
	i "github.com/ihulsbus/cookbook/internal/interfaces"
	m "github.com/ihulsbus/cookbook/internal/models"
)

type IngredientService interface {
	FindAll() ([]m.Ingredient, error)
	Create(ingredient m.Ingredient) (m.Ingredient, error)
}

// NewRecipeService creates a new RecipeService instance
func NewIngredientService(repo i.IngredientRepository) ingredientService {
	return ingredientService{
		repository: repo,
	}
}

type ingredientService struct {
	repository i.IngredientRepository
}

func (s ingredientService) FindAll() ([]m.Ingredient, error) {
	var ingredients []m.Ingredient

	ingredients, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return ingredients, nil
}

func (s ingredientService) Create(ingredient m.Ingredient) (m.Ingredient, error) {
	var response m.Ingredient

	if err := s.repository.Create(ingredient); err != nil {
		return response, err
	}

	return response, nil
}
