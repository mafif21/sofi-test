package repositories

import (
	"penjadwalan-sidang-new/internal/model/entity"
	"time"
)

type ScheduleRepository interface {
	FindAll(filter map[string]interface{}, page int) ([]entity.Schedule, error)
	FindPengajuanSchedules(pengajuanId int, filter map[string]interface{}, order string) ([]entity.Schedule, error)
	FindById(scheduleId int) (*entity.Schedule, error)
	FindByPengajuanId(pengajuanId int) (*entity.Schedule, error)
	FindMemberSchedule(schedule *entity.Schedule) ([]entity.Schedule, error)
	CheckAvailRoom(dateTime time.Time, room string, pengajuanId []int, condition string) ([]entity.Schedule, error)
	CheckAvailUser(dateTime time.Time, userId int, column string, pengajuanId []int, condition string) ([]entity.Schedule, error)
	Save(newSchedules []entity.Schedule) ([]entity.Schedule, error)
	UpdateMany(newSchedules []entity.Schedule) ([]entity.Schedule, error)
	Update(newSchedules *entity.Schedule) (*entity.Schedule, error)
	Delete(schedules []entity.Schedule) error
}
