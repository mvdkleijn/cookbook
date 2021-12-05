package services

import (
	i "github.com/ihulsbus/cookbook/internal/interfaces"
	m "github.com/ihulsbus/cookbook/internal/models"
)

type IngredientService interface {
	FindAll() ([]m.Ingredient, error)
	Create(ingredient m.Ingredient) (m.Ingredient, error)
	Update(ingredient m.Ingredient) (m.Ingredient, error)
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

	response, err := s.repository.Create(ingredient)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (s ingredientService) Update(ingredient m.Ingredient) (m.Ingredient, error) {
	var response m.Ingredient

	response, err := s.repository.Update(ingredient)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (s ingredientService) Delete(ingredient m.Ingredient) error {

	// TODO: check if there are recipies using the ingredient. If so, an error should be returned and the ingredient should not be deleted.
	err := s.repository.Delete(ingredient)
	if err != nil {
		return err
	}

	return nil
}
