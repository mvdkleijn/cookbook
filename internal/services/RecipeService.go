package services

import (
	i "github.com/ihulsbus/cookbook/internal/interfaces"
	m "github.com/ihulsbus/cookbook/internal/models"
)

type RecipeService interface {
	FindAll() ([]m.Recipe, error)
	Create(recipe m.RecipeInput) (m.Recipe, error)
}

// NewRecipeService creates a new RecipeService instance
func NewRecipeService(repo i.RecipeRepository) recipeService {
	return recipeService{
		repository: repo,
	}
}

type recipeService struct {
	repository i.RecipeRepository
}

func (s recipeService) FindAll() ([]m.Recipe, error) {
	var recipes []m.Recipe

	recipes, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return recipes, nil
}

func (s recipeService) Create(recipe m.RecipeInput) (m.Recipe, error) {
	var response m.Recipe

	if err := s.repository.Create(recipe); err != nil {
		return response, err
	}

	return response, nil
}

// converting
func convertRecipeToDataObject(recipe m.Recipe) m.Recipe {
	return m.Recipe{
		ID:             recipe.ID,
		Title:          recipe.Title,
		Description:    recipe.Description,
		Method:         recipe.Method,
		PrepTime:       recipe.PrepTime,
		CookTime:       recipe.CookTime,
		TotalTime:      recipe.TotalTime,
		Amount_Persons: recipe.Amount_Persons,
	}
}

func convertRecipesToDataObject(recipes []m.Recipe) []m.Recipe {
	var data []m.Recipe

	for _, recipe := range recipes {
		data = append(data, convertRecipeToDataObject(recipe))
	}

	return data
}
