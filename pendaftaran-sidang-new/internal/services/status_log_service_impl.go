package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"pendaftaran-sidang-new/internal/exception"
	"pendaftaran-sidang-new/internal/helper"
	"pendaftaran-sidang-new/internal/model/entity"
	"pendaftaran-sidang-new/internal/model/web"
	"pendaftaran-sidang-new/internal/repositories"
)

type StatusLogServiceImpl struct {
	StatusLogRepository repositories.StatusLogRepository
	PengajuanRepository repositories.PengajuanRepository
	Validator           *validator.Validate
}

func NewStatusLogService(statusLogRepository repositories.StatusLogRepository, pengajuanRepository repositories.PengajuanRepository, validator *validator.Validate) StatusLogService {
	return StatusLogServiceImpl{StatusLogRepository: statusLogRepository, PengajuanRepository: pengajuanRepository, Validator: validator}
}

func (s StatusLogServiceImpl) GetAllStatusLogs(filter map[string]interface{}) ([]web.StatusLogResponse, error) {
	statusLogs, err := s.StatusLogRepository.FindAll(filter)

	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusInternalServerError, "internal server error")
	}

	var statusLogsResponse []web.StatusLogResponse
	for _, status := range statusLogs {
		statusLogsResponse = append(statusLogsResponse, helper.ToStatusLogResponse(&status))
	}

	return statusLogsResponse, nil
}

func (s StatusLogServiceImpl) GetStatusLogsById(statusId int) (*web.StatusLogResponse, error) {
	status, err := s.StatusLogRepository.FindStatusLogById(statusId)

	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	response := helper.ToStatusLogResponse(status)
	return &response, nil
}

func (s StatusLogServiceImpl) Create(request *web.StatusLogCreateRequest) (*web.StatusLogResponse, error) {
	if err := s.Validator.Struct(request); err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	foundPengajuan, err := s.PengajuanRepository.FindPengajuanById(request.PengajuanId)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	newStatusLogs := &entity.StatusLog{
		Feedback:     request.Feedback,
		CreatedBy:    request.CreatedBy,
		WorkFlowType: request.WorkFlowType,
		Name:         request.Name,
		PengajuanID:  foundPengajuan.ID,
	}

	savedData, err := s.StatusLogRepository.Save(newStatusLogs)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "failed to save status logs")
	}

	response := helper.ToStatusLogResponse(savedData)
	return &response, nil
}

func (s StatusLogServiceImpl) Delete(statusId int) error {
	found, err := s.StatusLogRepository.FindStatusLogById(statusId)

	if err != nil {
		return exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	err = s.StatusLogRepository.Delete(found.ID)
	if err != nil {
		return exception.NewErrorResponse(fiber.StatusNotFound, "failed to delete status logs")
	}

	return nil
}
