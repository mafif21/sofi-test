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

type DocumentLogServiceImpl struct {
	DocumentLogRepository repositories.DocumentLogRepository
	PengajuanRepository   repositories.PengajuanRepository
	Validator             *validator.Validate
}

func NewDocumentLogService(documentLogRepo repositories.DocumentLogRepository, pengajuanRepository repositories.PengajuanRepository, validator *validator.Validate) DocumentLogService {
	return &DocumentLogServiceImpl{
		DocumentLogRepository: documentLogRepo,
		PengajuanRepository:   pengajuanRepository,
		Validator:             validator,
	}
}

func (s DocumentLogServiceImpl) GetAllDocumentLog(filter map[string]interface{}, order string, limit int, page int) ([]web.DocumentLogResponse, error) {
	var allDocuments []web.DocumentLogResponse

	documents, err := s.DocumentLogRepository.FindAll(filter, order, limit, page)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusInternalServerError, "internal server error")
	}

	for _, document := range documents {
		allDocuments = append(allDocuments, helper.ToDocumentLogResponse(&document))
	}

	return allDocuments, nil
}

func (s DocumentLogServiceImpl) GetDocumentLogById(documentId int) (*web.DocumentLogDetailResponse, error) {
	found, err := s.DocumentLogRepository.FindDocumentLogById(documentId)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	response := helper.ToDocumentLogDetailResponse(found)
	return &response, nil
}

func (s DocumentLogServiceImpl) GetLatestDocument(filter map[string]interface{}, pengajuanId int) (*web.DocumentLogDetailResponse, error) {
	foundPengajuan, err := s.PengajuanRepository.FindPengajuanById(pengajuanId)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	document, err := s.DocumentLogRepository.FindLatestDocument(filter, foundPengajuan.ID)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	response := helper.ToDocumentLogDetailResponse(document)
	return &response, nil
}

func (s DocumentLogServiceImpl) Create(request *web.DocumentLogCreateRequest) (*web.DocumentLogResponse, error) {
	if err := s.Validator.Struct(request); err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	newDocument := &entity.DocumentLog{
		PengajuanID: request.PengajuanId,
		FileName:    request.FileName,
		Type:        request.Type,
		FileUrl:     request.FileUrl,
		CreatedBy:   request.CreatedBy,
	}

	savedDocument, err := s.DocumentLogRepository.Save(newDocument)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	response := helper.ToDocumentLogResponse(savedDocument)
	return &response, nil
}

func (s DocumentLogServiceImpl) Update(request *web.DocumentLogUpdateRequest) (*web.DocumentLogResponse, error) {
	if err := s.Validator.Struct(request); err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	found, err := s.DocumentLogRepository.FindDocumentLogById(request.Id)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	found.Type = request.Type

	updatedData, err := s.DocumentLogRepository.Update(found)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	response := helper.ToDocumentLogResponse(updatedData)
	return &response, nil
}

func (s DocumentLogServiceImpl) Delete(documentId int) error {
	found, err := s.DocumentLogRepository.FindDocumentLogById(documentId)
	if err != nil {
		return exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	err = s.DocumentLogRepository.Delete(found.ID)
	if err != nil {
		return exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	return nil
}
