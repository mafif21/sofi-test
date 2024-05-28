package services

import "pendaftaran-sidang-new/internal/model/web"

type DocumentLogService interface {
	GetAllDocumentLog(filter map[string]interface{}, order string, limit int, page int) ([]web.DocumentLogResponse, error)
	GetDocumentLogById(documentId int) (*web.DocumentLogDetailResponse, error)
	GetLatestDocument(filter map[string]interface{}, pengajuanId int) (*web.DocumentLogDetailResponse, error)
	Create(request *web.DocumentLogCreateRequest) (*web.DocumentLogResponse, error)
	Update(request *web.DocumentLogUpdateRequest) (*web.DocumentLogResponse, error)
	Delete(documentId int) error
}
