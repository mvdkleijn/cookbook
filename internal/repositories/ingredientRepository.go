package repositories

import (
	m "github.com/ihulsbus/cookbook/internal/models"
	"gorm.io/gorm"
)

type gormIngredientRepository struct {
	db *gorm.DB
}

func NewGormIngredientRepository(db *gorm.DB) gormIngredientRepository {
	return gormIngredientRepository{
		db: db,
	}
}

func (r gormIngredientRepository) FindAll() ([]m.Ingredient, error) {
	var ingredients []m.Ingredient

	if err := r.db.Find(&ingredients).Error; err != nil {
		return nil, err
	}

	return ingredients, nil
}

func (r gormIngredientRepository) FindRecipeIngredients(recipeID int) ([]m.Recipe_Ingredient, error) {
	var RecipeIngredients []m.Recipe_Ingredient

	if err := r.db.Find(&RecipeIngredients).Where("recipe_id = ?", recipeID).Error; err != nil {
		return nil, err
	}

	return RecipeIngredients, nil
}

func (r gormIngredientRepository) Create(ingredient m.Ingredient) (m.Ingredient, error) {

	if err := r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&ingredient).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return ingredient, err
	}

	return ingredient, nil
}

func (r gormIngredientRepository) Update(ingredient m.Ingredient) (m.Ingredient, error) {

	if err := r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&ingredient).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return ingredient, err
	}

	return ingredient, nil
}

func (r gormIngredientRepository) Delete(ingredient m.Ingredient) error {

	if err := r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Delete(&ingredient).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
