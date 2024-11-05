package models

import "gorm.io/gorm"

// Product model definition
type Product struct {
	gorm.Model
	Name  string
	Price float64
}
