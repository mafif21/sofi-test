package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"pendaftaran-sidang-new/internal/exception"
	"pendaftaran-sidang-new/internal/helper"
	"pendaftaran-sidang-new/internal/model/entity"
	"pendaftaran-sidang-new/internal/model/web"
	"pendaftaran-sidang-new/internal/repositories"
	"time"
)

type PeriodServiceImpl struct {
	PeriodRepository repositories.PeriodRepository
	Validator        *validator.Validate
}

func NewPeriodService(periodRepository repositories.PeriodRepository, validator *validator.Validate) PeriodService {
	return &PeriodServiceImpl{
		PeriodRepository: periodRepository, Validator: validator,
	}
}

func (s PeriodServiceImpl) GetAllPeriod(filter map[string]interface{}) ([]web.PeriodResponse, error) {
	var allPeriodResponse []web.PeriodResponse

	periodDatas, err := s.PeriodRepository.FindAll(filter)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusInternalServerError, "internal server error")
	}

	for _, period := range periodDatas {
		allPeriodResponse = append(allPeriodResponse, helper.ToPeriodResponse(&period))
	}

	return allPeriodResponse, nil
}

func (s PeriodServiceImpl) GetPeriodById(periodId int) (*web.PeriodDetailResponse, error) {
	foundPeriod, err := s.PeriodRepository.FindPeriodById(periodId)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	response := helper.ToPeriodDetailResponse(foundPeriod)
	return &response, nil
}

func (s PeriodServiceImpl) Create(request *web.PeriodCreateRequest) (*web.PeriodResponse, error) {
	if err := s.Validator.Struct(request); err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	if request.StartDate.After(request.EndDate) || request.StartDate.Equal(request.EndDate) {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "period date is not valid")
	}

	newPeriod := &entity.Period{
		Name:        request.Name,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
		Description: request.Description,
	}

	newPeriod, err := s.PeriodRepository.Save(newPeriod)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	response := helper.ToPeriodResponse(newPeriod)
	return &response, nil
}

func (s PeriodServiceImpl) Update(request *web.PeriodUpdateRequest) (*web.PeriodResponse, error) {
	if err := s.Validator.Struct(request); err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	if request.StartDate.After(request.EndDate) || request.StartDate.Equal(request.EndDate) {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "period date is not valid")
	}

	foundPeriod, err := s.PeriodRepository.FindPeriodById(request.Id)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	foundPeriod.Name = request.Name
	foundPeriod.StartDate = request.StartDate
	foundPeriod.EndDate = request.EndDate
	foundPeriod.Description = request.Description

	update, err := s.PeriodRepository.Update(foundPeriod)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	response := helper.ToPeriodResponse(update)
	return &response, nil
}

func (s PeriodServiceImpl) Delete(periodId int) error {
	foundPeriod, err := s.PeriodRepository.FindPeriodById(periodId)
	if err != nil {
		return exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	err = s.PeriodRepository.Delete(foundPeriod.ID)
	if err != nil {
		return exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	return nil
}

func (s PeriodServiceImpl) GetPeriodByTime() (*web.PeriodResponse, error) {
	now := time.Now()

	foundPeriod, err := s.PeriodRepository.FindPeriodByTime(now)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	response := helper.ToPeriodResponse(foundPeriod)
	return &response, nil
}
