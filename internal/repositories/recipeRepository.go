package repositories

import (
	"log"

	m "github.com/ihulsbus/cookbook/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type gormRecipeRepository struct {
	db *gorm.DB
}

func NewGormRecipeRepository(db *gorm.DB) gormRecipeRepository {
	return gormRecipeRepository{
		db: db,
	}
}

// FindAll retrieves all recipes from the database and returns them in a slice
func (r gormRecipeRepository) FindAll() ([]m.Recipe, error) {
	var recipes []m.Recipe

	if err := r.db.Preload(clause.Associations).Find(&recipes).Error; err != nil {
		return nil, err
	}

	return recipes, nil
}

// Find searches for a specific recipe in the database and returns it when found.
func (r gormRecipeRepository) Find(recipeID int) (m.Recipe, error) {
	var recipe m.Recipe

	if err := r.db.Preload(clause.Associations).Where("ID = ?", recipeID).Find(&recipe).Error; err != nil {
		return recipe, err
	}

	return recipe, nil
}

// FindRecipeIngredients finds all ingredients associated to a recipe and returns them in a slice
func (r gormRecipeRepository) FindRecipeIngredients(recipeID int) ([]m.Recipe_Ingredient, error) {
	var recipeIngredients []m.Recipe_Ingredient

	return recipeIngredients, nil
}

// Create handles the creation of a recipe and stores the relevant information in the database
func (r gormRecipeRepository) Create(recipe m.Recipe) (m.Recipe, error) {

	if err := r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&recipe).Error; err != nil {
			return err
		}

		log.Printf("insertID: %d", recipe.ID)

		return nil
	}); err != nil {
		return recipe, err
	}

	return recipe, nil
}

func (r gormRecipeRepository) Update(recipe m.Recipe) (m.Recipe, error) {

	if err := r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Updates(&recipe).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return recipe, err
	}

	return recipe, nil
}
