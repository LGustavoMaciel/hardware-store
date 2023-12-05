package repository

import (
	"github.com/jinzhu/gorm"

	"hardware-store/api/models"
)


type CategoriesRepository interface{
	Save(*models.Category) (*models.Category, error)
} 

type categoriesRepositoryImpl struct {
	db *gorm.DB
}


func NewCategoriesRepositoryImpl(db *gorm.DB) *categoriesRepositoryImpl {
	return &categoriesRepositoryImpl{db}
}


func (r *categoriesRepositoryImpl) Save(category *models.Category) (*models.Category, error) {
	tx := r.db.Begin()
    err := tx.Debug().Model(&models.Category{}).Create(category).Error

    if err!= nil {
        tx.Rollback()
        return nil, err
    }
    return category, tx.Commit().Error
}