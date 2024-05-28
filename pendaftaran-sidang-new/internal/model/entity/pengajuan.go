package entity

import (
	"time"
)

type Pengajuan struct {
	ID             int           `gorm:"primaryKey;column:id;autoIncrement"`
	UserId         int           `gorm:"column:user_id"`
	Nim            string        `gorm:"column:nim"`
	Pembimbing1Id  int           `gorm:"column:pembimbing1_id"`
	Pembimbing2Id  int           `gorm:"column:pembimbing2_id"`
	Judul          string        `gorm:"column:judul"`
	Eprt           int           `gorm:"column:eprt"`
	DocTa          string        `gorm:"column:doc_ta"`
	Makalah        string        `gorm:"column:makalah"`
	Tak            int           `gorm:"column:tak"`
	Status         string        `gorm:"column:status;default:belum disetujui admin"`
	SksLulus       int           `gorm:"column:sks_lulus"`
	SksBelumLulus  int           `gorm:"column:sks_belum_lulus"`
	IsEnglish      bool          `gorm:"column:is_english"`
	PeriodID       int           `gorm:"column:period_id;foreignKey:PeriodID;references:ID"`
	SkPenguji      string        `gorm:"column:sk_penguji"`
	FormBimbingan1 int           `gorm:"column:form_bimbingan1"`
	FormBimbingan2 int           `gorm:"column:form_bimbingan2"`
	DocumentLogs   []DocumentLog `gorm:"foreignKey:PengajuanID;references:ID"`
	StatusLogs     []StatusLog   `gorm:"foreignKey:PengajuanID;references:ID"`
	ScheduleID     int           `gorm:"column:schedule_id;foreignKey:ScheduleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;default:null"`
	Kk             string        `gorm:"column:kk"`
	CreatedAt      time.Time     `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time     `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (pengajuan *Pengajuan) TableName() string {
	return "pengajuans"
}
