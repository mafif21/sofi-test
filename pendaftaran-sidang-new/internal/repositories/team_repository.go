package repositories

import "pendaftaran-sidang-new/internal/model/entity"

type TeamRepository interface {
	FindAll(filter map[string]interface{}, page int, limit int) ([]entity.Team, error)
	FindTeamById(teamId int) (*entity.Team, error)
	Save(team *entity.Team, pengajuan *entity.Pengajuan, statusLog *entity.StatusLog) (*entity.Team, error)
	Update(team *entity.Team) (*entity.Team, error)
	Delete(teamId int) error
}
