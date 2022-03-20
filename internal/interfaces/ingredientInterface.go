package interfaces

import (
	m "github.com/ihulsbus/cookbook/internal/models"
)

type IngredientRepository interface {
	IngredientFindAll() ([]m.Ingredient, error)
	IngredientFindSingle(ingredientID int) ([]m.Ingredient, error)
	SectionFindAll() ([]m.Section, error)
	SectionFindSingle(sectionID int) ([]m.Section, error)
	FindRecipeIngredients(recipeID int) ([]m.Recipe_Ingredient, error)
	CreateIngredient(ingredient m.Ingredient) (m.Ingredient, error)
	UpdateIngredient(ingredient m.Ingredient) (m.Ingredient, error)
	DeleteIngredient(ingredient m.Ingredient) error
}
