package models

type RecipeInput struct {
	ID             int               `json:"id"`
	Title          string            `json:"title"`
	Description    string            `json:"description"`
	Ingredients    []IngredientInput `json:"ingredients"`
	Method         string            `json:"method"`
	PrepTime       int               `json:"preptime"`
	CookTime       int               `json:"cooktime"`
	TotalTime      int               `json:"totaltime"`
	Amount_Persons int               `json:"persons"`
}

type Recipe struct {
	ID             int          `gorm:"primaryKey;serial;unique;not null;autoIncrement" json:"id"`
	Title          string       `gorm:"not null" json:"title"`
	Description    string       `gorm:"not null" json:"description"`
	Ingredients    []Ingredient `gorm:"many2many:Recipe_Ingredient" json:"ingredients"`
	Method         string       `gorm:"not null" json:"method"`
	PrepTime       int          `gorm:"default:0" json:"preptime"`
	CookTime       int          `gorm:"default:0" json:"cooktime"`
	TotalTime      int          `gorm:"default:0" json:"totaltime"`
	Amount_Persons int          `gorm:"default:0" json:"persons"`
}

type IngredientInput struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Unit   string `json:"unit"`
}

type Ingredient struct {
	ID   int    `gorm:"primaryKey;serial;unique;not null;autoIncrement" json:"id"`
	Name string `gorm:"not null" json:"name"`
}

type Recipe_Ingredient struct {
	RecipeID     int    `gorm:"primaryKey;not null" json:"-"`
	IngredientID int    `gorm:"primaryKey;not null" json:"-"`
	Amount       int    `gorm:"default:0" json:"amount"`
	Unit         string `gorm:"not null" json:"unit"`
}
