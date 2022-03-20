package repositories

import (
	"fmt"

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

func (r gormIngredientRepository) IngredientFindAll() ([]m.Ingredient, error) {
	var ingredients []m.Ingredient

	if err := r.db.Find(&ingredients).Error; err != nil {
		return nil, err
	}

	return ingredients, nil
}

func (r gormIngredientRepository) IngredientFindSingle(ingredientID int) ([]m.Ingredient, error) {
	var ingredient []m.Ingredient

	if err := r.db.Where("id = ?", ingredientID).Find(&ingredient).Error; err != nil {
		return nil, err
	}

	return ingredient, nil
}

func (r gormIngredientRepository) SectionFindAll() ([]m.Section, error) {
	var sections []m.Section

	if err := r.db.Find(&sections).Error; err != nil {
		return nil, err
	}

	return sections, nil
}

func (r gormIngredientRepository) SectionFindSingle(sectionID int) ([]m.Section, error) {
	var section []m.Section

	if err := r.db.Where("id = ?", sectionID).Find(&section).Error; err != nil {
		return nil, err
	}

	return section, nil
}

func (r gormIngredientRepository) FindRecipeIngredients(recipeID int) ([]m.Recipe_Ingredient, error) {
	var RecipeIngredients []m.Recipe_Ingredient

	if err := r.db.Preload("SectionID").Where("recipe_id = ?", recipeID).Find(&RecipeIngredients).Error; err != nil {
		return nil, err
	}
	fmt.Println(RecipeIngredients)
	return RecipeIngredients, nil
}

func (r gormIngredientRepository) CreateIngredient(ingredient m.Ingredient) (m.Ingredient, error) {

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

func (r gormIngredientRepository) UpdateIngredient(ingredient m.Ingredient) (m.Ingredient, error) {

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

func (r gormIngredientRepository) DeleteIngredient(ingredient m.Ingredient) error {

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
