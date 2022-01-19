package models

import "gorm.io/gorm"

type Warehouse struct {
	gorm.Model
	Name     string `json:"name" gorm:"size:255"`
	Location string `json:"location" gorm:"size:255"`
}
