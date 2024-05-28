package repositories

import (
	"errors"
	"gorm.io/gorm"
	"pendaftaran-sidang-new/internal/model/entity"
)

type TeamRepositoryImpl struct {
	DB *gorm.DB
}

func NewTeamRepository(db *gorm.DB) TeamRepository {
	return &TeamRepositoryImpl{
		DB: db,
	}
}

func (r TeamRepositoryImpl) FindAll(filter map[string]interface{}, page int, limit int) ([]entity.Team, error) {
	var teams []entity.Team

	tx := r.DB

	for key, value := range filter {
		tx = tx.Where(key, value)
	}

	err := tx.Order("created_at DESC").Offset((page - 1) * limit).Limit(limit).Find(&teams).Error
	if err != nil {
		return nil, err
	}

	return teams, nil

}

func (r TeamRepositoryImpl) FindTeamById(teamId int) (*entity.Team, error) {
	var team entity.Team

	tx := r.DB

	err := tx.Where("id = ?", teamId).First(&team).Error
	if err != nil {
		return nil, errors.New("data is not found")
	}

	return &team, nil
}

func (r TeamRepositoryImpl) Save(team *entity.Team, pengajuan *entity.Pengajuan, statusLog *entity.StatusLog) (*entity.Team, error) {
	tx := r.DB.Begin()

	if tx.Error != nil {
		return nil, tx.Error
	}

	if err := tx.Create(&team).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Updates(&pengajuan).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Create(&statusLog).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return team, nil
}

func (r TeamRepositoryImpl) Update(team *entity.Team) (*entity.Team, error) {
	tx := r.DB

	var existingName *entity.Team
	err := tx.Model(&entity.Team{}).Where("name = ?", team.Name).Where("id != ?", team.ID).First(&existingName).Error
	if err == nil {
		return nil, errors.New("found same team name")
	}

	err = tx.Save(&team).Error
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (r TeamRepositoryImpl) Delete(teamId int) error {
	tx := r.DB

	err := tx.Delete(&entity.Team{}, "id = ?", teamId).Error
	if err != nil {
		return err
	}

	return nil
}
