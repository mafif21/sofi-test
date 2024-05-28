package services

import (
	"penjadwalan-sidang-new/internal/model/web"
)

type ScheduleService interface {
	GetAll(filter map[string]interface{}, page int, token string) ([]web.ScheduleDetailResponse, error)
	GetPengajuanSchedules(pengajuanId int, filter map[string]interface{}, order string) ([]web.ScheduleResponse, error)
	GetById(scheduleId int, token string) (*web.ScheduleDetailResponse, error)
	GetByPengajuan(pengajuan *web.Pengajuan) (*web.ScheduleDetailResponse, error)
	Create(request *web.ScheduleCreateRequest, token string) (*[]web.ScheduleDetailResponse, error)
	Update(request *web.ScheduleUpdateRequest, token string) (*[]web.ScheduleDetailResponse, error)
	Delete(scheduleId int, token string) error
	AddFlag(code string, scheduleId int) (*web.ScheduleResponse, error)
	ChangeStatus(request *web.ScheduleUpdateStatusRequest) (*web.ScheduleResponse, error)
}
