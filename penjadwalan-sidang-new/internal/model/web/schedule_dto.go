package web

import (
	"time"
)

type ValidationCreateSchedule struct {
	Members []struct {
		UserId        int    `json:"user_id"`
		PengajuanId   int    `json:"pengajuan_id"`
		Pembimbing1Id int    `json:"pembimbing1_id"`
		Pembimbing2Id int    `json:"pembimbing2_id"`
		Kk            string `json:"kk" validate:"required"`
	} `json:"members"`
}

type ScheduleCreateRequest struct {
	DateTime time.Time `json:"date_time" validate:"required"`
	Room     string    `json:"room" validate:"required"`
	Penguji1 struct {
		Id  int    `json:"id" validate:"required"`
		Jfa string `json:"jfa" validate:"required"`
		Kk  string `json:"kk" validate:"required"`
	} `json:"penguji1" validate:"required"`
	Penguji2 struct {
		Id  int    `json:"id" validate:"required"`
		Jfa string `json:"jfa" validate:"required"`
		Kk  string `json:"kk" validate:"required"`
	} `json:"penguji2" validate:"required"`
	ValidationCreateSchedule
}

type ScheduleUpdateRequest struct {
	Id       int       `json:"id"`
	DateTime time.Time `json:"date_time" validate:"required"`
	Room     string    `json:"room" validate:"required"`
	Penguji1 struct {
		Id  int    `json:"id" validate:"required"`
		Jfa string `json:"jfa" validate:"required"`
		Kk  string `json:"kk" validate:"required"`
	} `json:"penguji1" validate:"required"`
	Penguji2 struct {
		Id  int    `json:"id" validate:"required"`
		Jfa string `json:"jfa" validate:"required"`
		Kk  string `json:"kk" validate:"required"`
	} `json:"penguji2" validate:"required"`
	ValidationCreateSchedule
}

type ScheduleUpdateStatusRequest struct {
	Id     int    `json:"id"`
	Status string `json:"status" validate:"required"`
}

type ScheduleResponse struct {
	ID               int       `json:"id"`
	DateTime         time.Time `json:"date_time"`
	Room             string    `json:"room"`
	Penguji1Id       int       `json:"penguji1_id"`
	Penguji2Id       int       `json:"penguji2_id"`
	Status           string    `json:"status"`
	Decision         string    `json:"decision"`
	RevisionDuration int       `json:"revision_duration"`
	FlagAddRevision  bool      `json:"flag_add_revision"`
	FlagChangeScores bool      `json:"flag_change_scores"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type ScheduleDetailResponse struct {
	ID               int                   `json:"id"`
	DateTime         time.Time             `json:"date_time"`
	Room             string                `json:"room"`
	Penguji1Id       int                   `json:"penguji1_id"`
	Penguji2Id       int                   `json:"penguji2_id"`
	Status           string                `json:"status"`
	Decision         string                `json:"decision"`
	RevisionDuration int                   `json:"revision_duration"`
	FlagAddRevision  bool                  `json:"flag_add_revision"`
	FlagChangeScores bool                  `json:"flag_change_scores"`
	Pengajuan        PengajuanRelationship `json:"pengajuan"`
	CreatedAt        time.Time             `json:"created_at"`
	UpdatedAt        time.Time             `json:"updated_at"`
}

type PengajuanRelationship struct {
	Id             int               `json:"id"`
	UserId         int               `json:"user_id"`
	Nim            string            `json:"nim"`
	Pembimbing1Id  int               `json:"pembimbing1_id"`
	Pembimbing2Id  int               `json:"pembimbing2_id"`
	Judul          string            `json:"judul"`
	Eprt           int               `json:"eprt"`
	DocTa          string            `json:"doc_ta"`
	Makalah        string            `json:"makalah"`
	Tak            int               `json:"tak"`
	Status         string            `json:"status"`
	SksLulus       int               `json:"sks_lulus"`
	SksBelumLulus  int               `json:"sks_belum_lulus"`
	IsEnglish      bool              `json:"is_english"`
	PeriodID       int               `json:"period_id"`
	SkPenguji      string            `json:"sk_penguji"`
	FormBimbingan1 int               `json:"form_bimbingan1"`
	FormBimbingan2 int               `json:"form_bimbingan2"`
	Kk             string            `json:"kk"`
	Slide          SlideRelationship `json:"slide"`
	CreatedAt      time.Time         `json:"created_at"`
	Updated_at     time.Time         `json:"updated_at"`
}

type SlideRelationship struct {
	Id          int       `json:"id"`
	PengajuanId int       `json:"pengajuan_id"`
	FileName    string    `json:"file_name"`
	Type        string    `json:"type"`
	FileUrl     string    `json:"file_url"`
	CreatedBy   int       `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
