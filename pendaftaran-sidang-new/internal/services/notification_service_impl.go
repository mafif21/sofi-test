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

type NotificationServiceImpl struct {
	Repositories repositories.NotificationRepository
	Validator    *validator.Validate
}

func NewNotificationService(repositories repositories.NotificationRepository, validator *validator.Validate) NotificationService {
	return &NotificationServiceImpl{Repositories: repositories, Validator: validator}
}

func (s NotificationServiceImpl) GetAllNotification(filter map[string]interface{}) ([]web.NotificationResponse, error) {
	allNotification, err := s.Repositories.GetAll(filter)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusInternalServerError, "internal server error")
	}

	var notificationResponse []web.NotificationResponse
	for _, notification := range allNotification {
		notificationResponse = append(notificationResponse, helper.ToNotificationResponse(&notification))
	}

	return notificationResponse, nil
}

func (s NotificationServiceImpl) GetAllNotificationById(id string) (*web.NotificationResponse, error) {
	found, err := s.Repositories.GetNotificationById(id)

	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	response := helper.ToNotificationResponse(found)
	return &response, nil
}

func (s NotificationServiceImpl) Update(request *web.NotificationUpdateRequest) (*web.NotificationResponse, error) {
	foundById, err := s.Repositories.GetNotificationById(request.Id)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	if foundById.UserId != request.UserId {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, "data not found")
	}

	foundById.ReadAt = request.ReadAt

	notificationUpdated, err := s.Repositories.UpdateNotification(foundById)
	response := helper.ToNotificationResponse(notificationUpdated)
	return &response, nil
}

func (s NotificationServiceImpl) Create(request *web.NotificationCreateRequest) ([]web.NotificationResponse, error) {
	if err := s.Validator.Struct(request); err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	var notificationCreate []entity.Notification
	for _, uid := range request.UserId {
		newNotification := entity.Notification{
			UserId:  uid,
			Title:   request.Title,
			Message: request.Message,
			Url:     request.Url,
		}

		notificationCreate = append(notificationCreate, newNotification)
	}

	var allNotifications []web.NotificationResponse

	newNotifications, err := s.Repositories.Save(notificationCreate)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	for _, notification := range newNotifications {
		allNotifications = append(allNotifications, helper.ToNotificationResponse(&notification))
	}

	return allNotifications, nil
}
