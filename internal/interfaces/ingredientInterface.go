package interfaces

import (
	m "github.com/ihulsbus/cookbook/internal/models"
)

type IngredientRepository interface {
	FindAll() ([]m.Ingredient, error)
	Create(ingredient m.Ingredient) (m.Ingredient, error)
	Update(ingredient m.Ingredient) (m.Ingredient, error)
}
