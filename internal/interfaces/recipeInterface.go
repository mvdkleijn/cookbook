package interfaces

import (
	m "github.com/ihulsbus/cookbook/internal/models"
)

type RecipeRepository interface {
	Find(recipeID int) (m.Recipe, error)
	FindAll() ([]m.Recipe, error)
	Create(recipe m.Recipe) (m.Recipe, error)
	Update(recipe m.Recipe) (m.Recipe, error)
	Delete(recipe m.Recipe) error
}
