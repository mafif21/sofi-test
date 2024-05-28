package repositories

import "pendaftaran-sidang-new/internal/model/entity"

type StatusLogRepository interface {
	FindAll(filter map[string]interface{}) ([]entity.StatusLog, error)
	FindStatusLogById(statusId int) (*entity.StatusLog, error)
	Save(statusLog *entity.StatusLog) (*entity.StatusLog, error)
	Delete(statusId int) error
}
