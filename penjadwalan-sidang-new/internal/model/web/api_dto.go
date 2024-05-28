package web

import "time"

type Pengajuan struct {
	Id             int    `json:"id"`
	UserId         int    `json:"user_id"`
	Nim            string `json:"nim"`
	Pembimbing1Id  int    `json:"pembimbing1_id"`
	Pembimbing2Id  int    `json:"pembimbing2_id"`
	Judul          string `json:"judul"`
	Eprt           int    `json:"eprt"`
	DocTa          string `json:"doc_ta"`
	Makalah        string `json:"makalah"`
	Tak            int    `json:"tak"`
	Status         string `json:"status"`
	SksLulus       int    `json:"sks_lulus"`
	SksBelumLulus  int    `json:"sks_belum_lulus"`
	IsEnglish      bool   `json:"is_english"`
	PeriodId       int    `json:"period_id"`
	SkPenguji      string `json:"sk_penguji"`
	FormBimbingan1 int    `json:"form_bimbingan1"`
	FormBimbingan2 int    `json:"form_bimbingan2"`
	Kk             string `json:"kk"`
	Slide          struct {
		Id          int       `json:"id"`
		PengajuanId int       `json:"pengajuan_id"`
		FileName    string    `json:"file_name"`
		Type        string    `json:"type"`
		FileUrl     string    `json:"file_url"`
		CreatedBy   int       `json:"created_by"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	} `json:"slide"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PengajuanResponseApi struct {
	Code   int       `json:"code"`
	Status string    `json:"status"`
	Data   Pengajuan `json:"data"`
}

type PengajuanDatasResponseApi struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    []Pengajuan `json:"data"`
}

type GetMemberUserLoggedInApi struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		Members []struct {
			UserId    int       `json:"user_id"`
			TeamId    int       `json:"team_id"`
			Nim       int       `json:"nim"`
			Username  string    `json:"username"`
			Name      string    `json:"name"`
			Pengajuan Pengajuan `json:"pengajuan"`
		} `json:"members"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"data"`
}
