package repositories

import (
	"medods_task/configs"
	"medods_task/internal/app/models"
)

type MainRepository struct {
}

func NewMainRepository() *MainRepository {
	return &MainRepository{}
}

func (m *MainRepository) Create(model models.Model) (interface{}, error) {
	result := configs.DbObject.Create(model)
	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}

func (m *MainRepository) SinglePreloading(model models.Model, target string, targetValue string) (interface{}, error) {
	result := configs.DbObject.Preload(target).Find(model, targetValue)
	if result.Error != nil {
		return nil, result.Error
	}

	return result, nil
}

func (m *MainRepository) GetAllBy(model models.Model, fieldName string, value interface{}) (interface{}, error) {
	var resultSlice []models.Model
	err := configs.DbObject.Model(model).Where(fieldName+" = ?", value).Find(&resultSlice).Error

	if err != nil {
		return nil, err
	}

	return resultSlice, nil
}
