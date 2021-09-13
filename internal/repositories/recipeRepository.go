package repositories

import (
	"fmt"
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

func (r gormRecipeRepository) FindAll() ([]m.Recipe, error) {
	var recipes []m.Recipe

	if err := r.db.Preload(clause.Associations).Find(&recipes).Error; err != nil {
		return nil, err
	}

	return recipes, nil
}

func (r gormRecipeRepository) Create(recipe m.RecipeInput) error {
	var recipeDB m.Recipe
	var ingredients []m.Recipe_Ingredient

	recipeDB.Title = recipe.Title
	recipeDB.Description = recipe.Description
	recipeDB.Method = recipe.Method
	recipeDB.PrepTime = recipe.PrepTime
	recipeDB.CookTime = recipe.CookTime
	recipeDB.TotalTime = recipe.TotalTime
	recipeDB.Amount_Persons = recipe.Amount_Persons

	if err := r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&recipeDB).Error; err != nil {
			return err
		}

		log.Printf("insertID: %d", recipeDB.ID)
		for _, ingredient := range recipe.Ingredients {
			var ingredientDB m.Ingredient
			var ingredientJoin m.Recipe_Ingredient

			ingredientDB.ID = ingredient.ID
			ingredientJoin.IngredientID = ingredient.ID
			ingredientDB.Name = ingredient.Name
			ingredientJoin.Amount = ingredient.Amount
			ingredientJoin.Unit = ingredient.Unit
			ingredientJoin.RecipeID = recipeDB.ID
			fmt.Printf("setting recipe ID %d to ingredient %d. result: %d", recipeDB.ID, ingredientJoin.IngredientID, ingredientJoin.RecipeID)

			recipeDB.Ingredients = append(recipeDB.Ingredients, ingredientDB)
			ingredients = append(ingredients, ingredientJoin)
		}
		fmt.Println(ingredients)

		fmt.Printf("ingredients: %v", ingredients)

		if err := tx.Create(&ingredients).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
