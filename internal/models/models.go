package models

type Section struct {
	ID   int    `gorm:"primaryKey;serial;unique;not null;autoIncrement" json:"id"`
	Name string `gorm:"not null" json:"name"`
}

type Ingredient struct {
	ID   int    `gorm:"primaryKey;serial;unique;not null;autoIncrement" json:"id"`
	Name string `gorm:"not null" json:"name"`
}

type Recipe struct {
	ID             int          `gorm:"primaryKey;serial;unique;not null;autoIncrement" json:"id"`
	Title          string       `gorm:"not null" json:"title"`
	Description    string       `gorm:"not null" json:"description"`
	Ingredients    []Ingredient `gorm:"many2many:Recipe_Ingredient" json:"ingredients"`
	Sections       []Section    `gorm:"many2many:Recipe_Ingredient" json:"sections"`
	Method         string       `gorm:"" json:"method"`
	PrepTime       int          `gorm:"default:0" json:"preptime"`
	CookTime       int          `gorm:"default:0" json:"cooktime"`
	TotalTime      int          `gorm:"default:0" json:"totaltime"`
	Amount_Persons int          `gorm:"default:0" json:"persons"`
}

type Recipe_Ingredient struct {
	RecipeID     int    `gorm:"primaryKey;not null" json:"recipeid"`
	IngredientID int    `gorm:"primaryKey;not null" json:"ingredientid"`
	SectionID    int    `gorm:"primaryKey;not null" json:"sectionid"`
	Amount       int    `gorm:"default:0" json:"amount"`
	Unit         string `gorm:"not null" json:"unit"`
}
