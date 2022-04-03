package repositories

import (
	m "github.com/ihulsbus/cookbook/internal/models"
	log "github.com/sirupsen/logrus"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Find func(r *RecipeRepository, recipeID int) (m.Recipe, error)
type FindAll func(r *RecipeRepository) ([]m.Recipe, error)
type Create func(r *RecipeRepository, recipe m.Recipe) (m.Recipe, error)
type Update func(r *RecipeRepository, recipe m.Recipe) (m.Recipe, error)
type Delete func(r *RecipeRepository, recipe m.Recipe) error

type RecipeRepository struct {
	db *gorm.DB

	Find    Find
	FindAll FindAll
	Create  Create
	Update  Update
	Delete  Delete
}

func NewRecipeRepository(db *gorm.DB) *RecipeRepository {
	return &RecipeRepository{
		db: db,

		Find:    find,
		FindAll: findAll,
		Create:  create,
		Update:  update,
		Delete:  delete,
	}
}

// FindAll retrieves all recipes from the database and returns them in a slice
func findAll(r *RecipeRepository) ([]m.Recipe, error) {
	var recipes []m.Recipe

	if err := r.db.Preload(clause.Associations).Find(&recipes).Error; err != nil {
		return nil, err
	}

	return recipes, nil
}

// Find searches for a specific recipe in the database and returns it when found.
func find(r *RecipeRepository, recipeID int) (m.Recipe, error) {
	var recipe m.Recipe

	if err := r.db.Preload(clause.Associations).Where(&m.Recipe{ID: recipeID}).Find(&recipe).Error; err != nil {
		return recipe, err
	}

	log.Info(recipe)
	return recipe, nil
}

// FindRecipeIngredients finds all ingredients associated to a recipe and returns them in a slice
// func findRecipeIngredients(r *RecipeRepository, recipeID int) ([]m.Recipe_Ingredient, error) {
// 	var recipeIngredients []m.Recipe_Ingredient

// 	return recipeIngredients, nil
// }

// Create handles the creation of a recipe and stores the relevant information in the database
func create(r *RecipeRepository, recipe m.Recipe) (m.Recipe, error) {

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

func update(r *RecipeRepository, recipe m.Recipe) (m.Recipe, error) {

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

func delete(r *RecipeRepository, recipe m.Recipe) error {

	if err := r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Delete(&recipe).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
