package repositories

import (
	"pendaftaran-sidang-new/internal/model/entity"
)

type NotificationRepository interface {
	GetAll(filter map[string]interface{}) ([]entity.Notification, error)
	GetNotificationById(notificationId string) (*entity.Notification, error)
	UpdateNotification(notification *entity.Notification) (*entity.Notification, error)
	Save(notification []entity.Notification) ([]entity.Notification, error)
}
