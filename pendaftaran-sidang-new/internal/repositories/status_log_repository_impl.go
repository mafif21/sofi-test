package repositories

import (
	"errors"
	"gorm.io/gorm"
	"pendaftaran-sidang-new/internal/model/entity"
)

type StatusLogRepositoryImpl struct {
	DB *gorm.DB
}

func NewStatusLogRepository(db *gorm.DB) StatusLogRepository {
	return &StatusLogRepositoryImpl{
		DB: db,
	}
}

func (r StatusLogRepositoryImpl) FindAll(filter map[string]interface{}) ([]entity.StatusLog, error) {
	var statusLogs []entity.StatusLog

	tx := r.DB

	for key, value := range filter {
		tx = tx.Where(key, value)
	}

	err := tx.Order("created_at DESC").Find(&statusLogs).Error
	if err != nil {
		return nil, err
	}

	return statusLogs, nil
}

func (r StatusLogRepositoryImpl) FindStatusLogById(statusId int) (*entity.StatusLog, error) {
	var statusLog *entity.StatusLog

	tx := r.DB

	err := tx.Model(&entity.StatusLog{}).Where("id = ?", statusId).First(&statusLog).Error

	if err != nil {
		return nil, errors.New("data not found")
	}

	return statusLog, nil
}

func (r StatusLogRepositoryImpl) Save(statusLog *entity.StatusLog) (*entity.StatusLog, error) {
	tx := r.DB

	err := tx.Create(&statusLog).Error
	if err != nil {
		return nil, err
	}

	return statusLog, nil
}

func (r StatusLogRepositoryImpl) Delete(statusId int) error {
	tx := r.DB

	err := tx.Delete(&entity.StatusLog{}, "id = ?", statusId).Error
	if err != nil {
		return err
	}

	return nil
}
