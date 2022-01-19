package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string    `json:"name"`
	Price       float32   `json:"price"`
	Quantity    float32   `json:"quantity"`
	Unit        string    `json:"unit"`
	WarehouseId uint      `json:"-"`
	Warehouse   Warehouse `json:"warehouse" gorm:"foreignKey:warehouse_id;references:ID"`
}
