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
	Delete(recipe m.RecipeDTO) error
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
	var updatedRecipe m.Recipe
	var originalRecipe m.Recipe

	originalRecipe, err := s.repository.Find(recipe.ID)
	if err != nil {
		return recipe, err
	}

	if recipe.Title == "" {
		recipe.Title = originalRecipe.Title
	}

	if recipe.Description == "" {
		recipe.Description = originalRecipe.Description
	}

	if recipe.PrepTime == 0 {
		recipe.PrepTime = originalRecipe.PrepTime
	}

	if recipe.CookTime == 0 {
		recipe.CookTime = originalRecipe.CookTime
	}

	if recipe.TotalTime == 0 {
		recipe.TotalTime = recipe.PrepTime + recipe.CookTime
	}

	if recipe.Amount_Persons == 0 {
		recipe.Amount_Persons = originalRecipe.Amount_Persons
	}

	updatedRecipe, err = s.repository.Update(recipe.ConvertFromDTO())
	if err != nil {
		return updatedRecipe.ConvertToDTO(), err
	}

	return updatedRecipe.ConvertToDTO(), nil
}

func (s recipeService) Delete(recipe m.RecipeDTO) error {

	if err := s.repository.Delete(recipe.ConvertFromDTO()); err != nil {
		return err
	}

	return nil
}
