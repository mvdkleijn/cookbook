package services

import (
	m "github.com/ihulsbus/cookbook/internal/models"
	r "github.com/ihulsbus/cookbook/internal/repositories"
)

type FindAllRecipes func(s *RecipeService) ([]m.RecipeDTO, error)
type FindSingleRecipe func(s *RecipeService, recipeID int) (m.Recipe, error)
type CreateRecipe func(s *RecipeService, recipe m.RecipeDTO) (m.RecipeDTO, error)
type UpdateRecipe func(s *RecipeService, recipe m.RecipeDTO) (m.RecipeDTO, error)
type DeleteRecipe func(s *RecipeService, recipe m.RecipeDTO) error

type RecipeService struct {
	repo *r.RecipeRepository

	FindAllRecipes   FindAllRecipes
	FindSingleRecipe FindSingleRecipe
	CreateRecipe     CreateRecipe
	UpdateRecipe     UpdateRecipe
	DeleteRecipe     DeleteRecipe
}

// NewRecipeService creates a new RecipeService instance
func NewRecipeService(recipeRepo *r.RecipeRepository) *RecipeService {
	return &RecipeService{
		repo: recipeRepo,

		FindAllRecipes:   findAllRecipes,
		FindSingleRecipe: findSingleRecipe,
		CreateRecipe:     createRecipe,
		UpdateRecipe:     updateRecipe,
		DeleteRecipe:     deleteRecipe,
	}
}

// Find contains the business logic to get all recipes
func findAllRecipes(s *RecipeService) ([]m.RecipeDTO, error) {
	var recipes []m.Recipe

	recipes, err := s.repo.FindAll(s.repo)
	if err != nil {
		return nil, err
	}

	return m.Recipe{}.ConvertAllToDTO(recipes), nil
}

// Find contains the business logic to get a specific recipe
func findSingleRecipe(s *RecipeService, recipeID int) (m.Recipe, error) {
	var recipe m.Recipe

	recipe, err := s.repo.Find(s.repo, recipeID)
	if err != nil {
		return recipe, err
	}

	return recipe, nil
}

// Create handles the business logic for the creation of a recipe and passes the recipe object to the recipe repo for processing
func createRecipe(s *RecipeService, recipe m.RecipeDTO) (m.RecipeDTO, error) {

	response, err := s.repo.Create(s.repo, recipe.ConvertFromDTO())
	if err != nil {
		return response.ConvertToDTO(), err
	}

	return response.ConvertToDTO(), nil
}

func updateRecipe(s *RecipeService, recipe m.RecipeDTO) (m.RecipeDTO, error) {
	var updatedRecipe m.Recipe
	var originalRecipe m.Recipe

	originalRecipe, err := s.repo.Find(s.repo, recipe.ID)
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

	if recipe.Amount_Persons == 0 {
		recipe.Amount_Persons = originalRecipe.Amount_Persons
	}

	updatedRecipe, err = s.repo.Update(s.repo, recipe.ConvertFromDTO())
	if err != nil {
		return updatedRecipe.ConvertToDTO(), err
	}

	return updatedRecipe.ConvertToDTO(), nil
}

func deleteRecipe(s *RecipeService, recipe m.RecipeDTO) error {

	if err := s.repo.Delete(s.repo, recipe.ConvertFromDTO()); err != nil {
		return err
	}

	return nil
}
