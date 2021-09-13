package models

import "time"

type MealPlan struct {
	ID        int       `gorm:"primaryKey;serial;unique;not null;autoIncrement"`
	Name      string    `gorm:"not null"`
	Notes     string    `json:"notes"`
	Recipes   []Recipe  `gorm:"many2many:MealPlan_Recipe" json:"recipes"`
	StartDate time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	EndDate   time.Time `gorm:"not null"`
}

type MealPlan_Recipe struct {
	MealPlanID int       `gorm:"not null"`
	RecipeID   int       `gorm:"not null"`
	Date       time.Time `gorm:"not null"`
}
