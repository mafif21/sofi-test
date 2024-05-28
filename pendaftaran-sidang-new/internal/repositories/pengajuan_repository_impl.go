package repositories

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"pendaftaran-sidang-new/internal/model/entity"
	"reflect"
)

type PengajuanRepositoryImpl struct {
	DB *gorm.DB
}

func NewPengajuanRepository(db *gorm.DB) PengajuanRepository {
	return &PengajuanRepositoryImpl{
		DB: db,
	}
}

func (r PengajuanRepositoryImpl) FindAll(filter map[string]interface{}) ([]entity.Pengajuan, error) {
	var allPengajuans []entity.Pengajuan

	tx := r.DB

	for key, value := range filter {
		switch key {
		case "pembimbing":
			pembimbingId, ok := value.(int)
			if !ok {
				return nil, errors.New("id is not valid")
			}
			tx = tx.Where("pembimbing1_id = ? OR pembimbing2_id = ?", pembimbingId, pembimbingId)

		default:
			if reflect.TypeOf(value).Kind() == reflect.Slice {
				tx = tx.Where(key+" IN (?)", value)
			} else {
				tx = tx.Where(key, value)
			}
		}
	}

	err := tx.Order("created_at DESC").Find(&allPengajuans).Error
	if err != nil {
		return nil, err
	}

	return allPengajuans, nil
}

func (r PengajuanRepositoryImpl) FindPengajuanById(pengajuanId int) (*entity.Pengajuan, error) {
	var pengajuan *entity.Pengajuan

	tx := r.DB

	err := tx.Model(&entity.Pengajuan{}).Preload("StatusLogs").
		Preload("DocumentLogs", func(db *gorm.DB) *gorm.DB {
			return db.Where("type = ?", "slide").Limit(1).Order("created_at DESC")
		}).
		Where("id = ?", pengajuanId).First(&pengajuan).Error
	if err != nil {
		return nil, errors.New("data pengajuan not found")
	}

	return pengajuan, nil
}

func (r PengajuanRepositoryImpl) FindPengajuanByUserId(userId int) (*entity.Pengajuan, error) {
	var pengajuan *entity.Pengajuan

	tx := r.DB

	err := tx.Model(&entity.Pengajuan{}).Preload("StatusLogs").
		Preload("DocumentLogs", func(db *gorm.DB) *gorm.DB {
			return db.Where("type = ?", "slide").Limit(1).Order("created_at DESC")
		}).
		Where("user_id = ?", userId).First(&pengajuan).Error

	if err != nil {
		return nil, errors.New("data pengajuan not found")
	}

	return pengajuan, nil
}

func (r PengajuanRepositoryImpl) Save(pengajuan *entity.Pengajuan) (*entity.Pengajuan, error) {
	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&pengajuan).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	newStatusLogs := &entity.StatusLog{
		Feedback:     "-",
		CreatedBy:    pengajuan.UserId,
		WorkFlowType: "pengajuan",
		Name:         "pengajuan",
		PengajuanID:  pengajuan.ID,
	}

	newDocLogs := []entity.DocumentLog{
		{
			PengajuanID: pengajuan.ID,
			FileName:    pengajuan.Makalah,
			Type:        "makalah",
			FileUrl:     fmt.Sprintf("/public/makalah/%s", pengajuan.Makalah),
			CreatedBy:   pengajuan.UserId,
		},
		{
			PengajuanID: pengajuan.ID,
			FileName:    pengajuan.DocTa,
			Type:        "draft",
			FileUrl:     fmt.Sprintf("/public/makalah/%s", pengajuan.DocTa),
			CreatedBy:   pengajuan.UserId,
		},
	}

	newNotification := []entity.Notification{
		{
			UserId:  pengajuan.Pembimbing1Id,
			Title:   "Mahasiswa daftar sidang",
			Message: "Mahasiswa dengan nim " + pengajuan.Nim + " telah mendaftarkan sidang",
			Url:     "/pengajuan/get",
		},
		{
			UserId:  pengajuan.Pembimbing2Id,
			Title:   "Mahasiswa daftar sidang",
			Message: "Mahasiswa dengan nim " + pengajuan.Nim + " telah mendaftarkan sidang",
			Url:     "/pengajuan/get",
		},
	}

	if err := tx.Create(newStatusLogs).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Create(newDocLogs).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Create(newNotification).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return pengajuan, nil
}

func (r PengajuanRepositoryImpl) Update(pengajuan *entity.Pengajuan) (*entity.Pengajuan, error) {
	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Save(&pengajuan).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	newStatusLogs := &entity.StatusLog{
		Feedback:     "-",
		CreatedBy:    pengajuan.UserId,
		WorkFlowType: "pengajuan",
		Name:         "perbaikan berkas",
		PengajuanID:  pengajuan.ID,
	}

	newDocLogs := []entity.DocumentLog{
		{
			PengajuanID: pengajuan.ID,
			FileName:    pengajuan.Makalah,
			Type:        "makalah",
			FileUrl:     fmt.Sprintf("/public/makalah/%s", pengajuan.Makalah),
			CreatedBy:   pengajuan.UserId,
		},
		{
			PengajuanID: pengajuan.ID,
			FileName:    pengajuan.DocTa,
			Type:        "draft",
			FileUrl:     fmt.Sprintf("/public/makalah/%s", pengajuan.DocTa),
			CreatedBy:   pengajuan.UserId,
		},
	}

	newNotification := []entity.Notification{
		{
			UserId:  pengajuan.Pembimbing1Id,
			Title:   "Mahasiswa Edit Data Pengajuan Sidang",
			Message: "Mahasiswa dengan nim " + pengajuan.Nim + " melakukan perubahan data daftar sidang",
			Url:     "/pengajuan/get",
		},
		{
			UserId:  pengajuan.Pembimbing2Id,
			Title:   "Mahasiswa Edit Data Pengajuan Sidang",
			Message: "Mahasiswa dengan nim " + pengajuan.Nim + " melakukan perubahan data daftar sidang",
			Url:     "/pengajuan/get",
		},
	}

	if err := tx.Create(newStatusLogs).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Create(newDocLogs).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Create(newNotification).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return pengajuan, nil
}

func (r PengajuanRepositoryImpl) UpdateStatus(pengajuan *entity.Pengajuan, statusLog *entity.StatusLog, notification *entity.Notification) (*entity.Pengajuan, error) {
	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Save(&pengajuan).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Create(statusLog).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if notification != nil {
		if err := tx.Create(notification).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return pengajuan, nil
}

func (r PengajuanRepositoryImpl) Delete(pengajuan *entity.Pengajuan) error {
	if err := r.DB.Where("pengajuan_id = ?", pengajuan.ID).Delete(&entity.StatusLog{}).Error; err != nil {
		return err
	}

	if err := r.DB.Where("pengajuan_id = ?", pengajuan.ID).Delete(&entity.DocumentLog{}).Error; err != nil {
		return err
	}

	if err := r.DB.Delete(&entity.Pengajuan{}, pengajuan.ID).Error; err != nil {
		return err
	}

	return nil
}
