package interfaces

import (
	m "github.com/ihulsbus/cookbook/internal/models"
)

type RecipeRepository interface {
	FindAll() ([]m.Recipe, error)
	Create(recipe m.RecipeInput) (m.Recipe, error)
}
