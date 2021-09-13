package models

import "time"

type ShoppingList struct {
	ID          int          `gorm:"primaryKey;serial;unique;not null;autoIncrement"`
	Ingredients []Ingredient `gorm:"many2many:ShoppingList_Ingredients"`
	CreatedDate time.Time    `gorm:"not null;default:CURRENT_TIMESTAMP"`
	Completed   bool         `gorm:"default:false"`
}
