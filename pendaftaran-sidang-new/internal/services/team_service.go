package services

import "pendaftaran-sidang-new/internal/model/web"

type TeamService interface {
	GetAllTeam(filter map[string]interface{}, page int, limit int) ([]web.TeamResponse, error)
	GetTeamById(teamId int) (*web.TeamResponse, error)
	GetTeamByUserId(userId int) (*web.TeamResponseDetail, error)
	GetAvailableMember() ([]web.AvailableMember, error)
	Create(request *web.TeamCreateRequest) (*web.TeamResponse, error)
	Update(request *web.TeamUpdateRequest) (*web.TeamResponse, error)
	Delete(teamId int, userId int) error
	AddMember(request *web.MemberRequest) (*string, error)
	LeaveTeam(request *web.MemberRequest) (*string, error)
}
