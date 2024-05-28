package repositories

import (
	"pendaftaran-sidang-new/internal/model/entity"
)

type PengajuanRepository interface {
	FindAll(filter map[string]interface{}) ([]entity.Pengajuan, error)
	FindPengajuanById(pengajuanId int) (*entity.Pengajuan, error)
	FindPengajuanByUserId(userId int) (*entity.Pengajuan, error)
	Save(pengajuan *entity.Pengajuan) (*entity.Pengajuan, error)
	Update(pengajuan *entity.Pengajuan) (*entity.Pengajuan, error)
	UpdateStatus(pengajuan *entity.Pengajuan, statusLog *entity.StatusLog, notification *entity.Notification) (*entity.Pengajuan, error)
	Delete(pengajuan *entity.Pengajuan) error
}
