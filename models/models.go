package models

import "gorm.io/gorm"

type Expenses struct {
	gorm.Model
	Name       string     `json:"name"`
	Money      int        `json:"money"`
	Categories []Category `gorm:"foreignKey:CategoryID" json:"Categories"`
}

type Incomes struct {
	gorm.Model
	Money      int        `json:"money"`
	Name       string     `json:"name"`
	Categories []Category `gorm:"foreignKey: CategoryID" json:"Categories"`
}
type Category struct {
	gorm.Model
	Name       string `json:"name"`
	CategoryID int    `json:"CategoryID"`
}
