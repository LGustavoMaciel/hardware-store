package repository

import "hardware-store/api/models"

type ProductsRepository interface {
	Save(*models.Product) (*models.Product, error)
}