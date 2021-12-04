package models

type IngredientDTO struct {
	RecipeID     int    `json:"recipeid"`
	IngredientID int    `json:"ingredientid"`
	SectionID    int    `json:"sectionid"`
	Amount       int    `json:"amount"`
	Unit         string `json:"unit"`
}

type Ingredient struct {
	ID   int    `gorm:"primaryKey;serial;unique;not null;autoIncrement" json:"id"`
	Name string `gorm:"not null" json:"name"`
}

type Recipe_Ingredient struct {
	RecipeID     int               `gorm:"primaryKey;not null" json:"-"`
	IngredientID int               `gorm:"primaryKey;not null" json:"-"`
	SectionID    IngredientSection `gorm:"many2many:Section" json:"-"`
	Amount       int               `gorm:"default:0" json:"amount"`
	Unit         string            `gorm:"not null" json:"unit"`
}

type IngredientSection struct {
	ID   int    `gorm:"primaryKey;serial;unique;not null;autoIncrement" json:"id"`
	Name string `gorm:"not null" json:"name"`
}

// ConvertToDTO converts a single m.Recipe_Ingredient to a single m.IngredientDTO
func (i Recipe_Ingredient) ConvertToDTO() IngredientDTO {
	return IngredientDTO{
		RecipeID:     i.RecipeID,
		IngredientID: i.IngredientID,
		SectionID:    i.SectionID.ID,
		Amount:       i.Amount,
		Unit:         i.Unit,
	}
}

// ConvertAllToDTO converts a slice of m.Recipe_Ingredient to a slice of m.IngredientDTO objcets
func (i Recipe_Ingredient) ConvertAllToDTO(ingredients []Recipe_Ingredient) []IngredientDTO {
	var data []IngredientDTO

	for _, ingredient := range ingredients {
		data = append(data, ingredient.ConvertToDTO())
	}

	return data
}
