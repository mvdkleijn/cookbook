package services

import (
	i "github.com/ihulsbus/cookbook/internal/interfaces"
	m "github.com/ihulsbus/cookbook/internal/models"
)

type IngredientService interface {
	IngredientFindAll() ([]m.Ingredient, error)
	IngredientFindSingle(ingredientID int) ([]m.Ingredient, error)
	SectionsFindAll() ([]m.Section, error)
	SectionsFindSingle(sectionID int) ([]m.Section, error)
	FindRecipeIngredients(recipeID int) ([]m.IngredientDTO, error)
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

func (s ingredientService) IngredientFindAll() ([]m.Ingredient, error) {
	var ingredients []m.Ingredient

	ingredients, err := s.repository.IngredientFindAll()
	if err != nil {
		return nil, err
	}

	return ingredients, nil
}

func (s ingredientService) IngredientFindSingle(ingredientID int) ([]m.Ingredient, error) {
	var ingredients []m.Ingredient

	ingredients, err := s.repository.IngredientFindSingle(ingredientID)
	if err != nil {
		return nil, err
	}

	return ingredients, nil
}

func (s ingredientService) SectionsFindAll() ([]m.Section, error) {
	var sections []m.Section

	sections, err := s.repository.SectionFindAll()
	if err != nil {
		return nil, err
	}

	return sections, nil
}

func (s ingredientService) SectionsFindSingle(sectionID int) ([]m.Section, error) {
	var sections []m.Section

	sections, err := s.repository.SectionFindSingle(sectionID)
	if err != nil {
		return nil, err
	}

	return sections, nil
}

func (s ingredientService) FindRecipeIngredients(recipeID int) ([]m.IngredientDTO, error) {
	var data []m.Recipe_Ingredient

	data, err := s.repository.FindRecipeIngredients(recipeID)
	if err != nil {
		return nil, err
	}

	return m.Recipe_Ingredient{}.ConvertAllToDTO(data), nil
}

func (s ingredientService) Create(ingredient m.Ingredient) (m.Ingredient, error) {
	var response m.Ingredient

	response, err := s.repository.CreateIngredient(ingredient)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (s ingredientService) Update(ingredient m.Ingredient) (m.Ingredient, error) {
	var response m.Ingredient

	response, err := s.repository.UpdateIngredient(ingredient)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (s ingredientService) Delete(ingredient m.Ingredient) error {

	// TODO: check if there are recipies using the ingredient. If so, an error should be returned and the ingredient should not be deleted.
	err := s.repository.DeleteIngredient(ingredient)
	if err != nil {
		return err
	}

	return nil
}
