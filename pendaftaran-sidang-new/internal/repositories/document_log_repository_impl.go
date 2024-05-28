package repositories

import (
	"errors"
	"gorm.io/gorm"
	"pendaftaran-sidang-new/internal/model/entity"
)

type DocumentLogRepositoryImpl struct {
	DB *gorm.DB
}

func NewDocumentLogRepository(db *gorm.DB) DocumentLogRepository {
	return &DocumentLogRepositoryImpl{
		DB: db,
	}
}

func (r DocumentLogRepositoryImpl) FindAll(filter map[string]interface{}, order string, limit int, page int) ([]entity.DocumentLog, error) {
	var documents []entity.DocumentLog

	tx := r.DB

	for key, value := range filter {
		tx = tx.Where(key, value)
	}

	err := tx.Order("created_at " + order).Offset((page - 1) * limit).Limit(limit).Find(&documents).Error
	if err != nil {
		return nil, err
	}

	return documents, nil
}

func (r DocumentLogRepositoryImpl) FindDocumentLogById(documentId int) (*entity.DocumentLog, error) {
	var document *entity.DocumentLog

	tx := r.DB

	err := tx.Where("id = ?", documentId).Preload("Pengajuan").First(&document).Error
	if err != nil {
		return nil, errors.New("document not found")
	}

	return document, nil
}

func (r DocumentLogRepositoryImpl) FindLatestDocument(filter map[string]interface{}, pengajuanId int) (*entity.DocumentLog, error) {
	var latestDocument entity.DocumentLog

	tx := r.DB

	for key, value := range filter {
		tx = tx.Where(key, value)
	}

	err := tx.Where("pengajuan_id = ?", pengajuanId).Preload("Pengajuan").Order("created_at DESC").First(&latestDocument).Error
	if err != nil {
		return nil, errors.New("document not found")
	}

	return &latestDocument, nil
}

func (r DocumentLogRepositoryImpl) Save(document *entity.DocumentLog) (*entity.DocumentLog, error) {
	tx := r.DB

	err := tx.Create(document).Error
	if err != nil {
		return nil, err
	}

	return document, nil
}

func (r DocumentLogRepositoryImpl) Update(document *entity.DocumentLog) (*entity.DocumentLog, error) {
	tx := r.DB

	err := tx.Save(&document).Error
	if err != nil {
		return nil, err
	}

	return document, nil
}

func (r DocumentLogRepositoryImpl) Delete(documentId int) error {
	tx := r.DB

	err := tx.Delete(&entity.DocumentLog{}, "id = ?", documentId).Error
	if err != nil {
		return err
	}

	return nil
}
