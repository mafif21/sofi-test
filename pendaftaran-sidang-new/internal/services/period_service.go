package services

import (
	"pendaftaran-sidang-new/internal/model/web"
)

type PeriodService interface {
	GetAllPeriod(filter map[string]interface{}) ([]web.PeriodResponse, error)
	GetPeriodById(periodId int) (*web.PeriodDetailResponse, error)
	Create(request *web.PeriodCreateRequest) (*web.PeriodResponse, error)
	Update(request *web.PeriodUpdateRequest) (*web.PeriodResponse, error)
	Delete(periodId int) error
	GetPeriodByTime() (*web.PeriodResponse, error)
}
