package services

import "pendaftaran-sidang-new/internal/model/web"

type StatusLogService interface {
	GetAllStatusLogs(filter map[string]interface{}) ([]web.StatusLogResponse, error)
	GetStatusLogsById(statusId int) (*web.StatusLogResponse, error)
	Create(request *web.StatusLogCreateRequest) (*web.StatusLogResponse, error)
	Delete(statusId int) error
}
