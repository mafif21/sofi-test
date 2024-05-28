package web

import "time"

type TeamCreateRequest struct {
	UserId int    `json:"user_id" validate:"required"`
	Name   string `json:"name" validate:"required"`
}

type TeamUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `json:"name" validate:"required"`
}

type TeamResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MemberData struct {
	UserId      int                `json:"user_id"`
	TeamId      int                `json:"team_id"`
	Nim         int                `json:"nim"`
	Username    string             `json:"username"`
	Name        string             `json:"name"`
	PeminatanId int                `json:"peminatan_id"`
	Pengajuan   *PengajuanResponse `json:"pengajuan"`
}

type AvailableMember struct {
	UserId int    `json:"user_id"`
	TeamId int    `json:"team_id"`
	Nim    int    `json:"nim"`
	Name   string `json:"name"`
}

type TeamResponseDetail struct {
	Id        int          `json:"id"`
	Name      string       `json:"name"`
	Members   []MemberData `json:"members"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type MemberRequest struct {
	UserId int `json:"user_id" validate:"required"`
	TeamId int `json:"team_id" validate:"required"`
}
