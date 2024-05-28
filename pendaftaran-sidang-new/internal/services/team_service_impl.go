package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"pendaftaran-sidang-new/internal/exception"
	"pendaftaran-sidang-new/internal/helper"
	"pendaftaran-sidang-new/internal/model/entity"
	"pendaftaran-sidang-new/internal/model/web"
	"pendaftaran-sidang-new/internal/repositories"
)

type TeamServiceImpl struct {
	TeamRepository        repositories.TeamRepository
	PengajuanRepository   repositories.PengajuanRepository
	DocumentLogRepository repositories.DocumentLogRepository
	Validator             *validator.Validate
}

func NewTeamService(teamRepository repositories.TeamRepository, pengajuanRepository repositories.PengajuanRepository, documentLogRepository repositories.DocumentLogRepository, validator *validator.Validate) TeamService {
	return &TeamServiceImpl{TeamRepository: teamRepository, PengajuanRepository: pengajuanRepository, DocumentLogRepository: documentLogRepository, Validator: validator}
}

func (s TeamServiceImpl) GetAllTeam(filter map[string]interface{}, page int, limit int) ([]web.TeamResponse, error) {
	teams, err := s.TeamRepository.FindAll(filter, page, limit)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusInternalServerError, "internal server error")
	}

	var teamResponse []web.TeamResponse
	for _, team := range teams {
		teamResponse = append(teamResponse, helper.ToTeamResponse(&team))
	}

	return teamResponse, nil
}

func (s TeamServiceImpl) GetTeamById(teamId int) (*web.TeamResponse, error) {
	found, err := s.TeamRepository.FindTeamById(teamId)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	response := helper.ToTeamResponse(found)
	return &response, nil
}

func (s TeamServiceImpl) GetTeamByUserId(userId int) (*web.TeamResponseDetail, error) {
	student, err := helper.GetDetailStudent(userId)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	if student.Data.TeamId < 1 {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, "user dont have team id")
	}

	found, err := s.TeamRepository.FindTeamById(student.Data.TeamId)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	members, err := helper.GetAllTeamMember(found.ID)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	for i := range members.Data {
		pengajuanData, err := s.PengajuanRepository.FindPengajuanByUserId(members.Data[i].UserId)
		if err == nil {
			pengajuanResponse := helper.ToPengajuanResponse(pengajuanData)
			members.Data[i].Pengajuan = &pengajuanResponse
		}
	}

	response := helper.ToTeamDetailResponse(found, members)
	return &response, nil
}

func (s TeamServiceImpl) GetAvailableMember() ([]web.AvailableMember, error) {
	//validPengajuan, err := s.PengajuanRepository.FindAll(map[string]interface{}{
	//	"status": []string{"telah disetujui admin", "tidak lulus (sudah update dokumen)"},
	//})
	//
	//if err != nil {
	//	return nil, exception.NewErrorResponse(fiber.StatusInternalServerError, "internal server error")
	//}
	//
	//var availableMember []web.AvailableMember
	//
	//for _, pengajuan := range validPengajuan {
	//	student, err := helper.GetDetailStudent(pengajuan.UserId)
	//	if err == nil {
	//		if student.Data.TeamId == 0 {
	//			memberData := web.AvailableMember{
	//				UserId: pengajuan.UserId,
	//				TeamId: student.Data.TeamId,
	//				Nim:    student.Data.Nim,
	//				Name:   student.Data.User.Name,
	//			}
	//
	//			availableMember = append(availableMember, memberData)
	//		}
	//	}
	//}
	//
	//return availableMember, nil
	panic("")

}

func (s TeamServiceImpl) Create(request *web.TeamCreateRequest) (*web.TeamResponse, error) {
	if err := s.Validator.Struct(request); err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	foundPengajuan, err := s.PengajuanRepository.FindPengajuanByUserId(request.UserId)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	filter := map[string]interface{}{
		"type": "slide",
	}

	_, err = s.DocumentLogRepository.FindLatestDocument(filter, foundPengajuan.ID)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, "make sure student already upload slide")
	}

	if foundPengajuan.Status == "ditolak oleh admin" || foundPengajuan.Status == "belum disetujui admin" {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "created team is invalid")
	}

	statusLogs := &entity.StatusLog{
		Feedback:     "-",
		CreatedBy:    request.UserId,
		WorkFlowType: "penjadwalan",
		Name:         "",
		PengajuanID:  foundPengajuan.ID,
	}

	if foundPengajuan.Status == "tidak lulus (sudah update dokumen)" {
		foundPengajuan.Status = "tidak lulus (belum dijadwalkan)"
		statusLogs.Name = "tidak lulus (belum dijadwalkan)"

	} else {
		foundPengajuan.Status = "belum dijadwalkan"
		statusLogs.Name = "belum dijadwalkan"
	}

	newTeam := &entity.Team{
		Name: request.Name,
	}

	savedTeam, err := s.TeamRepository.Save(newTeam, foundPengajuan, statusLogs)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	//_, err = helper.UpdateUserTeamID(request.UserId, savedTeam.ID)
	//if err != nil {
	//	return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	//}

	response := helper.ToTeamResponse(savedTeam)
	return &response, nil
}

func (s TeamServiceImpl) Update(request *web.TeamUpdateRequest) (*web.TeamResponse, error) {
	if err := s.Validator.Struct(request); err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	foundTeam, err := s.TeamRepository.FindTeamById(request.Id)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, "team not found")
	}

	foundTeam.Name = request.Name

	updatedData, err := s.TeamRepository.Update(foundTeam)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	response := helper.ToTeamResponse(updatedData)
	return &response, nil
}

func (s TeamServiceImpl) Delete(teamId int, userId int) error {
	foundTeam, err := s.TeamRepository.FindTeamById(teamId)
	if err != nil {
		return exception.NewErrorResponse(fiber.StatusNotFound, "team not found")
	}

	err = s.TeamRepository.Delete(foundTeam.ID)
	if err != nil {
		return exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	return nil
}

func (s TeamServiceImpl) AddMember(request *web.MemberRequest) (*string, error) {
	if err := s.Validator.Struct(request); err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	_, err := s.TeamRepository.FindTeamById(request.TeamId)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	foundPengajuan, err := s.PengajuanRepository.FindPengajuanByUserId(request.UserId)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	filter := map[string]interface{}{
		"type": "slide",
	}

	_, err = s.DocumentLogRepository.FindLatestDocument(filter, foundPengajuan.ID)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, "make sure student already upload slide")
	}

	if foundPengajuan.Status == "ditolak oleh admin" || foundPengajuan.Status == "belum disetujui admin" {
		_, err = helper.UpdateUserTeamID(foundPengajuan.UserId, 0)
		if err != nil {
			return nil, exception.NewErrorResponse(fiber.StatusNotFound, "student not in database")
		}

		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "created team is invalid")
	}

	//_, err = helper.UpdateUserTeamID(request.UserId, request.TeamId)
	//if err != nil {
	//	return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	//}

	statusLogs := &entity.StatusLog{
		Feedback:     "-",
		CreatedBy:    request.UserId,
		WorkFlowType: "penjadwalan",
		PengajuanID:  foundPengajuan.ID,
	}

	notification := &entity.Notification{
		UserId:  request.UserId,
		Title:   "Invite Team",
		Message: "Sukses diundang ke dalam team sidang",
		Url:     "/team/get-team",
	}

	if foundPengajuan.Status == "tidak lulus (sudah update dokumen)" {
		foundPengajuan.Status = "tidak lulus (belum dijadwalkan)"
		statusLogs.Name = "tidak lulus (belum dijadwalkan)"

		_, err = s.PengajuanRepository.UpdateStatus(foundPengajuan, statusLogs, notification)
		if err != nil {
			return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
		}

		return &foundPengajuan.Nim, nil
	}

	foundPengajuan.Status = "belum dijadwalkan"
	statusLogs.Name = "belum dijadwalkan"

	_, err = s.PengajuanRepository.UpdateStatus(foundPengajuan, statusLogs, notification)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	return &foundPengajuan.Nim, nil

}

func (s TeamServiceImpl) LeaveTeam(request *web.MemberRequest) (*string, error) {
	if err := s.Validator.Struct(request); err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	_, err := s.TeamRepository.FindTeamById(request.TeamId)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	foundPengajuan, err := s.PengajuanRepository.FindPengajuanByUserId(request.UserId)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	//_, err = helper.UpdateUserTeamID(request.UserId, 0)
	//if err != nil {
	//	return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	//}

	//member, err := helper.GetAllTeamMember(request.TeamId)
	//if err != nil {
	//	return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	//}

	//if len(member.Data) < 1 {
	//	err = s.TeamRepository.Delete(request.TeamId)
	//	if err != nil {
	//		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	//	}
	//}

	statusLogs := &entity.StatusLog{
		Feedback:     "-",
		CreatedBy:    request.UserId,
		WorkFlowType: "pengajuan",
		PengajuanID:  foundPengajuan.ID,
	}

	notification := &entity.Notification{
		UserId:  request.UserId,
		Title:   "Leave Team",
		Message: "User telah keluar dari team",
		Url:     "/team/get-team",
	}

	if foundPengajuan.Status == "tidak lulus (belum dijadwalkan)" {
		foundPengajuan.Status = "tidak lulus (sudah update dokumen)"
		statusLogs.Name = "tidak lulus (sudah update dokumen)"

		_, err = s.PengajuanRepository.UpdateStatus(foundPengajuan, statusLogs, notification)
		if err != nil {
			return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
		}

		return &foundPengajuan.Nim, nil
	}

	foundPengajuan.Status = "telah disetujui admin"
	statusLogs.Name = "telah disetujui admin"

	_, err = s.PengajuanRepository.UpdateStatus(foundPengajuan, statusLogs, notification)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	return &foundPengajuan.Nim, nil
}
