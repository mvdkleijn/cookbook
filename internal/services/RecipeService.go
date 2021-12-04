package services

import (
	i "github.com/ihulsbus/cookbook/internal/interfaces"
	m "github.com/ihulsbus/cookbook/internal/models"
)

type RecipeService interface {
	Find(recipeID int) (m.RecipeDTO, error)
	FindAll() ([]m.RecipeDTO, error)
	Create(recipe m.RecipeDTO) (m.RecipeDTO, error)
	Update(recipe m.RecipeDTO) (m.RecipeDTO, error)
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

// Find contains the business logic to get all recipes
func (s recipeService) FindAll() ([]m.RecipeDTO, error) {
	var recipes []m.Recipe

	recipes, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return m.Recipe{}.ConvertAllToDTO(recipes), nil
}

// Find contains the business logic to get a specific recipe
func (s recipeService) Find(recipeID int) (m.RecipeDTO, error) {
	var recipe m.Recipe

	recipe, err := s.repository.Find(recipeID)
	if err != nil {
		return recipe.ConvertToDTO(), err
	}

	return recipe.ConvertToDTO(), nil
}

// Create handles the business logic for the creation of a recipe and passes the recipe object to the recipe repo for processing
func (s recipeService) Create(recipe m.RecipeDTO) (m.RecipeDTO, error) {

	response, err := s.repository.Create(recipe.ConvertFromDTO())
	if err != nil {
		return response.ConvertToDTO(), err
	}

	return response.ConvertToDTO(), nil
}

func (s recipeService) Update(recipe m.RecipeDTO) (m.RecipeDTO, error) {
	var updateRecipe m.Recipe

	updateRecipe, err := s.repository.Update(recipe.ConvertFromDTO())
	if err != nil {
		return updateRecipe.ConvertToDTO(), err
	}

	return updateRecipe.ConvertToDTO(), nil
}
