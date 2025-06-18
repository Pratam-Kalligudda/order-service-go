package repository

import "gorm.io/gorm"

type CartRepository interface{}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}
