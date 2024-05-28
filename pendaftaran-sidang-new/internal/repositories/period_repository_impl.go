package repositories

import (
	"errors"
	"gorm.io/gorm"
	"pendaftaran-sidang-new/internal/model/entity"
	"time"
)

type PeriodRepositoryImpl struct {
	DB *gorm.DB
}

func NewPeriodRepository(db *gorm.DB) PeriodRepository {
	return &PeriodRepositoryImpl{DB: db}
}

func (r PeriodRepositoryImpl) FindAll(filter map[string]interface{}) ([]entity.Period, error) {
	var periods []entity.Period

	tx := r.DB

	for key, value := range filter {
		tx = tx.Where(key, value)
	}

	err := tx.Find(&periods).Error
	if err != nil {
		return nil, err
	}

	return periods, nil
}

func (r PeriodRepositoryImpl) FindPeriodById(periodId int) (*entity.Period, error) {
	var period *entity.Period

	tx := r.DB

	err := tx.Model(&entity.Period{}).Preload("Pengajuans").Take(&period, "id = ?", periodId).Error
	if err != nil {
		return nil, errors.New("data not found")
	}

	return period, nil
}

func (r PeriodRepositoryImpl) Save(period *entity.Period) (*entity.Period, error) {
	tx := r.DB

	var existingPeriod entity.Period
	err := tx.Model(&entity.Period{}).Where("name = ?", period.Name).First(&existingPeriod).Error
	if err == nil {
		return nil, errors.New("found same duplicate name")
	}

	err = tx.Create(&period).Error
	if err != nil {
		return nil, err
	}

	return period, nil
}

func (r PeriodRepositoryImpl) Update(period *entity.Period) (*entity.Period, error) {
	tx := r.DB

	var existingPeriod entity.Period
	err := tx.Model(&entity.Period{}).Where("name = ?", period.Name).Where("id != ?", period.ID).First(&existingPeriod).Error
	if err == nil {
		return nil, errors.New("found same duplicate name")
	}

	err = tx.Save(&period).Error
	if err != nil {
		return nil, err
	}

	return period, nil
}

func (r PeriodRepositoryImpl) Delete(periodId int) error {
	err := r.DB.Delete(&entity.Period{}, "id = ?", periodId).Error
	if err != nil {
		return err
	}

	return nil
}

func (r PeriodRepositoryImpl) FindPeriodByTime(time time.Time) (*entity.Period, error) {
	var foundPeriod entity.Period

	tx := r.DB

	err := tx.Where("? BETWEEN start_date AND end_date", time).Take(&foundPeriod).Error
	if err != nil {
		return nil, errors.New("time is not valid with period data")
	}

	return &foundPeriod, nil
}
