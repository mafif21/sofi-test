package repositories

import (
	"errors"
	"gorm.io/gorm"
	"penjadwalan-sidang-new/internal/model/entity"
	"reflect"
	"strings"
	"time"
)

type ScheduleRepositoryImpl struct {
	DB *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &ScheduleRepositoryImpl{DB: db}
}

func (r ScheduleRepositoryImpl) FindAll(filter map[string]interface{}, page int) ([]entity.Schedule, error) {
	var allSchedules []entity.Schedule
	dataLimit := 10

	tx := r.DB

	for key, value := range filter {
		switch key {
		case "penguji":
			pengujiID, ok := value.(int)
			if !ok {
				return nil, errors.New("Nilai penguji tidak valid")
			}
			tx = tx.Where("penguji1_id = ? OR penguji2_id = ?", pengujiID, pengujiID)

		default:
			if reflect.TypeOf(value).Kind() == reflect.Slice {
				tx = tx.Where(key+" IN (?)", value)
			} else {
				tx = tx.Where(key, value)
			}
		}
	}

	err := tx.Limit(dataLimit).Offset((page - 1) * dataLimit).Order("created_at DESC").Find(&allSchedules).Error
	if err != nil {
		return nil, err
	}

	return allSchedules, nil
}

func (r ScheduleRepositoryImpl) FindById(scheduleId int) (*entity.Schedule, error) {
	var foundData *entity.Schedule

	err := r.DB.Where("id = ?", scheduleId).First(&foundData).Error
	if err != nil {
		return nil, errors.New("schedule is not found")
	}

	return foundData, nil
}

func (r ScheduleRepositoryImpl) FindPengajuanSchedules(pengajuanId int, filter map[string]interface{}, order string) ([]entity.Schedule, error) {
	var foundData []entity.Schedule

	tx := r.DB.Begin()

	err := tx.Model(&entity.Schedule{}).Where("pengajuan_id = ?", pengajuanId).Error
	if err != nil {
		return nil, errors.New("pengajuan dont have schedule")
	}

	for key, value := range filter {
		tx = tx.Where(key, value)
	}

	tx.Order("created_at " + order).Find(&foundData)

	return foundData, nil
}

func (r ScheduleRepositoryImpl) FindByPengajuanId(pengajuanId int) (*entity.Schedule, error) {
	var foundData *entity.Schedule

	err := r.DB.Where("pengajuan_id = ?", pengajuanId).Where("status = ?", "belum dilaksanakan").First(&foundData).Error
	if err != nil {
		return nil, errors.New("make sure user have submit new schedule")
	}

	return foundData, nil
}

func (r ScheduleRepositoryImpl) FindMemberSchedule(schedule *entity.Schedule) ([]entity.Schedule, error) {
	var memberSchedules []entity.Schedule

	err := r.DB.Where("date_time = ? AND ruang = ? AND penguji1_id = ? AND penguji2_id = ? AND status = ?", schedule.DateTime, schedule.Room, schedule.Penguji1Id, schedule.Penguji2Id, schedule.Status).
		Find(&memberSchedules).Error

	if err != nil {
		return nil, err
	}

	return memberSchedules, nil
}

func (r ScheduleRepositoryImpl) CheckAvailRoom(dateTime time.Time, room string, pengajuanId []int, condition string) ([]entity.Schedule, error) {
	timeEnd := dateTime.Add(2 * time.Hour)
	var foundData []entity.Schedule

	tx := r.DB

	if strings.EqualFold(condition, "create") {
		tx.Where("ruang = ? AND status = ?", room, "belum dilaksanakan").
			Where("((date_time >= ? AND date_time < ?) OR (date_time < ? AND date_time > ?))", dateTime, timeEnd, timeEnd, dateTime.Add(-2*time.Hour)).Find(&foundData)
	} else {
		tx.Where("ruang = ? AND status = ?", room, "belum dilaksanakan").
			Where("((date_time >= ? AND date_time < ?) OR (date_time < ? AND date_time > ?))", dateTime, timeEnd, timeEnd, dateTime.Add(-2*time.Hour)).
			Where("pengajuan_id NOT IN ?", pengajuanId).Find(&foundData)
	}

	err := tx.Error

	if err != nil {
		return nil, err
	}

	return foundData, nil
}

func (r ScheduleRepositoryImpl) CheckAvailUser(dateTime time.Time, userId int, column string, pengajuanId []int, condition string) ([]entity.Schedule, error) {
	timeEnd := dateTime.Add(2 * time.Hour)
	var foundData []entity.Schedule

	tx := r.DB

	if strings.EqualFold(condition, "create") {
		tx.Where(column+" = ? AND status = ?", userId, "belum dilaksanakan").
			Where("((date_time >= ? AND date_time < ?) OR (date_time < ? AND date_time > ?))", dateTime, timeEnd, timeEnd, dateTime.Add(-2*time.Hour)).
			Find(&foundData)
	} else {
		tx.Where(column+" = ? AND status = ?", userId, "belum dilaksanakan").
			Where("((date_time >= ? AND date_time < ?) OR (date_time < ? AND date_time > ?))", dateTime, timeEnd, timeEnd, dateTime.Add(-2*time.Hour)).
			Where("pengajuan_id NOT IN ?", pengajuanId).
			Find(&foundData)
	}

	err := tx.Error

	if err != nil {
		return nil, err
	}

	return foundData, nil
}

func (r ScheduleRepositoryImpl) Save(newSchedules []entity.Schedule) ([]entity.Schedule, error) {
	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	//existsStatus := []string{"belum dilaksanakan", "sedang dilaksanakan"}

	//var validData []entity.Schedule
	//for _, schedule := range newSchedules {
	//	var count int64
	//	if err := tx.Model(&entity.Schedule{}).
	//		Where("pengajuan_id = ? AND status IN (?)", schedule.PengajuanId, existsStatus).
	//		Count(&count).Error; err != nil {
	//		tx.Rollback()
	//		return nil, err
	//	}
	//	if count == 0 {
	//		validData = append(validData, schedule)
	//	}
	//}

	if err := tx.Create(&newSchedules).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return newSchedules, nil
}

func (r ScheduleRepositoryImpl) UpdateMany(newSchedules []entity.Schedule) ([]entity.Schedule, error) {
	err := r.DB.Save(&newSchedules).Error
	if err != nil {
		return nil, err
	}

	return newSchedules, nil
}

func (r ScheduleRepositoryImpl) Update(newSchedules *entity.Schedule) (*entity.Schedule, error) {
	err := r.DB.Updates(&newSchedules).Error
	if err != nil {
		return nil, err
	}

	return newSchedules, nil
}

func (r ScheduleRepositoryImpl) Delete(schedules []entity.Schedule) error {
	err := r.DB.Delete(&schedules).Error
	if err != nil {
		return err
	}

	return nil
}
