package repositories

import (
	"errors"
	"gorm.io/gorm"
	"pendaftaran-sidang-new/internal/model/entity"
)

type NotificationRepositoryImpl struct {
	DB *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &NotificationRepositoryImpl{
		DB: db,
	}
}

func (r NotificationRepositoryImpl) GetAll(filter map[string]interface{}) ([]entity.Notification, error) {
	var notifications []entity.Notification

	tx := r.DB

	for key, value := range filter {
		tx = tx.Where(key, value)
	}

	err := tx.Find(&notifications).Error
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

func (r NotificationRepositoryImpl) GetNotificationById(notificationId string) (*entity.Notification, error) {
	var notification *entity.Notification

	tx := r.DB

	err := tx.Model(&entity.Notification{}).Where("id = ?", notificationId).First(&notification).Error
	if err != nil {
		return nil, errors.New("data not found")
	}

	return notification, nil

}

func (r NotificationRepositoryImpl) UpdateNotification(notification *entity.Notification) (*entity.Notification, error) {
	tx := r.DB

	err := tx.Save(notification).Error
	if err != nil {
		return nil, err
	}

	return notification, nil
}

func (r NotificationRepositoryImpl) Save(notification []entity.Notification) ([]entity.Notification, error) {
	tx := r.DB

	err := tx.Create(&notification).Error
	if err != nil {
		return nil, err
	}

	return notification, nil
}
