package models

import "gorm.io/gorm"

type Expenses struct {
	gorm.Model
	Name       string   `json:"name"`
	Money      int      `json:"money"`
	CategoryID uint     `gorm:"not null" json:"category_id"`
	Categories Category `gorm:"foreignKey:CategoryID" json:"category"`
	// Categories []Category `gorm:"foreignKey:CategoryID" json:"Categories"`
}

type Incomes struct {
	gorm.Model
	Money      int      `json:"money"`
	Name       string   `json:"name"`
	CategoryID uint     `json:"category_id"`                           // Поле для ID категории
	Categories Category `gorm:"foreignKey:CategoryID" json:"category"` // Связь с категорией
	// Categories []Category `gorm:"foreignKey: CategoryID" json:"category"`
}
type Category struct {
	gorm.Model
	Name string `json:"name"`
}
