package web

import "time"

type DocumentLogCreateRequest struct {
	PengajuanId int    `validate:"required"`
	FileName    string `validate:"required"`
	Type        string `validate:"required"`
	FileUrl     string `validate:"required"`
	CreatedBy   int    `json:"created_by"`
}

type DocumentLogUpdateRequest struct {
	Id   int    `validate:"required"`
	Type string `json:"type" validate:"required"`
}

type DocumentLogResponse struct {
	Id          int       `json:"id"`
	PengajuanId int       `json:"pengajuan_id"`
	FileName    string    `json:"file_name"`
	Type        string    `json:"type"`
	FileUrl     string    `json:"file_url"`
	CreatedBy   int       `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type DocumentLogDetailResponse struct {
	Id        int               `json:"id"`
	FileName  string            `json:"file_name"`
	Type      string            `json:"type"`
	FileUrl   string            `json:"file_url"`
	CreatedBy int               `json:"created_by"`
	Pengajuan PengajuanResponse `json:"pengajuan"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type StatusLogCreateRequest struct {
	Feedback     string `json:"feedback" validate:"required"`
	CreatedBy    int    `validate:"required"`
	WorkFlowType string `json:"workflow_type" validate:"required"`
	Name         string `json:"name" validate:"required"`
	PengajuanId  int    `json:"pengajuan_id" validate:"required"`
}
