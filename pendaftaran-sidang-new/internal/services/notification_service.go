package services

import "pendaftaran-sidang-new/internal/model/web"

type NotificationService interface {
	GetAllNotification(filter map[string]interface{}) ([]web.NotificationResponse, error)
	GetAllNotificationById(id string) (*web.NotificationResponse, error)
	Update(request *web.NotificationUpdateRequest) (*web.NotificationResponse, error)
	Create(request *web.NotificationCreateRequest) ([]web.NotificationResponse, error)
}
