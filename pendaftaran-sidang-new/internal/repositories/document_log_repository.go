package repositories

import (
	"pendaftaran-sidang-new/internal/model/entity"
)

type DocumentLogRepository interface {
	FindAll(filter map[string]interface{}, order string, limit int, page int) ([]entity.DocumentLog, error)
	FindDocumentLogById(documentId int) (*entity.DocumentLog, error)
	FindLatestDocument(filter map[string]interface{}, pengajuanId int) (*entity.DocumentLog, error)
	Save(document *entity.DocumentLog) (*entity.DocumentLog, error)
	Update(document *entity.DocumentLog) (*entity.DocumentLog, error)
	Delete(documentId int) error
}
