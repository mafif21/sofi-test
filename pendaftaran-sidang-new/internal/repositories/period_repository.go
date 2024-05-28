package repositories

import (
	"pendaftaran-sidang-new/internal/model/entity"
	"time"
)

type PeriodRepository interface {
	FindAll(filter map[string]interface{}) ([]entity.Period, error)
	FindPeriodById(periodId int) (*entity.Period, error)
	Save(period *entity.Period) (*entity.Period, error)
	Update(period *entity.Period) (*entity.Period, error)
	Delete(periodId int) error
	FindPeriodByTime(time time.Time) (*entity.Period, error)
}
