package models

type UserData struct {
	RecipeID int    `gorm:"not null"`
	Notes    string `gorm:"not null"`
}
