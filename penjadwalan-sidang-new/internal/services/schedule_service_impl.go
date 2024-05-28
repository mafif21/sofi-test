package services

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"penjadwalan-sidang-new/internal/exception"
	"penjadwalan-sidang-new/internal/helper"
	"penjadwalan-sidang-new/internal/model/entity"
	"penjadwalan-sidang-new/internal/model/web"
	"penjadwalan-sidang-new/internal/repositories"
	"strconv"
	"time"
)

type ScheduleServiceImpl struct {
	ScheduleRepository repositories.ScheduleRepository
	Validator          *validator.Validate
}

func NewScheduleService(scheduleRepository repositories.ScheduleRepository, validator *validator.Validate) ScheduleService {
	return &ScheduleServiceImpl{ScheduleRepository: scheduleRepository, Validator: validator}
}

func (s ScheduleServiceImpl) GetAll(filter map[string]interface{}, page int, token string) ([]web.ScheduleDetailResponse, error) {
	var allSchedulesResponse []web.ScheduleDetailResponse

	schedules, err := s.ScheduleRepository.FindAll(filter, page)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusInternalServerError, "internal server error")
	}

	for _, schedule := range schedules {
		foundPengajuan := helper.GetPengajuanById(schedule.PengajuanId, token)
		allSchedulesResponse = append(allSchedulesResponse, helper.ToScheduleDetailResponse(&schedule, &foundPengajuan.Data))
	}

	return allSchedulesResponse, nil
}

func (s ScheduleServiceImpl) GetPengajuanSchedules(pengajuanId int, filter map[string]interface{}, order string) ([]web.ScheduleResponse, error) {
	schedules, err := s.ScheduleRepository.FindPengajuanSchedules(pengajuanId, filter, order)
	if err != nil || len(schedules) < 1 {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, "data not found")
	}

	var response []web.ScheduleResponse
	for _, schedule := range schedules {
		response = append(response, helper.ToScheduleResponse(&schedule))
	}
	return response, nil
}

func (s ScheduleServiceImpl) GetById(scheduleId int, token string) (*web.ScheduleDetailResponse, error) {
	schedule, err := s.ScheduleRepository.FindById(scheduleId)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, "data not found")
	}

	foundPengajuan := helper.GetPengajuanById(schedule.PengajuanId, token)
	response := helper.ToScheduleDetailResponse(schedule, &foundPengajuan.Data)
	return &response, nil
}

func (s ScheduleServiceImpl) GetByPengajuan(pengajuan *web.Pengajuan) (*web.ScheduleDetailResponse, error) {
	schedule, err := s.ScheduleRepository.FindByPengajuanId(pengajuan.Id)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, "data not found")
	}

	response := helper.ToScheduleDetailResponse(schedule, pengajuan)
	return &response, nil
}

func (s ScheduleServiceImpl) Create(request *web.ScheduleCreateRequest, token string) (*[]web.ScheduleDetailResponse, error) {
	if err := s.Validator.Struct(request); err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	var pengajuansId []int

	if time.Now().After(request.DateTime) {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "time has passed")
	}

	validTime := helper.GetValidSidangTime()
	if request.DateTime.Before(validTime) || request.DateTime.Equal(validTime) {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "scheduling must be exactly 2 hours in advance")
	}

	if request.Penguji1.Id == request.Penguji2.Id {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "penguji 1 and penguji 2 must be different")
	}

	if request.Penguji1.Jfa != "NJFA" {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "penguji 1 must have jfa")
	}

	for _, pengajuanMember := range request.Members {
		if pengajuanMember.Kk != request.Penguji1.Kk && pengajuanMember.Kk != request.Penguji2.Kk {
			return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "have member with different kk (kompetensi keahlian)")
		}

		if pengajuanMember.Pembimbing1Id == request.Penguji1.Id || pengajuanMember.Pembimbing2Id == request.Penguji2.Id {
			return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "pengajuan with id "+strconv.Itoa(pengajuanMember.PengajuanId)+" penguji 1 and penguji 2 cant same with pembimbing 1")
		}

		if pengajuanMember.Pembimbing2Id == request.Penguji1.Id || pengajuanMember.Pembimbing2Id == request.Penguji2.Id {
			return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "pengajuan with id "+strconv.Itoa(pengajuanMember.PengajuanId)+" penguji 1 and penguji 2 cant same with pembimbing 2")
		}

		pengajuansId = append(pengajuansId, pengajuanMember.PengajuanId)
	}

	rooms, _ := s.ScheduleRepository.CheckAvailRoom(request.DateTime, request.Room, pengajuansId, "create")
	if len(rooms) > 0 {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "room is not available")
	}

	foundDataPenguji1, _ := s.ScheduleRepository.CheckAvailUser(request.DateTime, request.Penguji1.Id, "penguji1_id", pengajuansId, "create")
	if len(foundDataPenguji1) > 0 {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "penguji 1 is not available")
	}

	foundDataPenguji2, _ := s.ScheduleRepository.CheckAvailUser(request.DateTime, request.Penguji2.Id, "penguji2_id", pengajuansId, "create")
	if len(foundDataPenguji2) > 0 {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "penguji 2 is not available")
	}

	var scheduleCreateData []entity.Schedule
	var usersId []int
	for _, id := range pengajuansId {
		newSchedule := entity.Schedule{
			PengajuanId: id,
			DateTime:    request.DateTime,
			Room:        request.Room,
			Penguji1Id:  request.Penguji1.Id,
			Penguji2Id:  request.Penguji2.Id,
		}

		scheduleCreateData = append(scheduleCreateData, newSchedule)
	}

	save, err := s.ScheduleRepository.Save(scheduleCreateData)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "cant make a new schedule")
	}

	var allSchedulesResponse []web.ScheduleDetailResponse
	for _, schedule := range save {
		foundPengajuan := helper.GetPengajuanById(schedule.PengajuanId, token)
		updatedStatus := "sudah dijadwalkan"
		if foundPengajuan.Data.Status == "tidak lulus (belum dijadwalkan)" {
			updatedStatus = "tidak lulus (sudah dijadwalkan)"
		}

		updatedData := helper.ChangePengajuanStatus("-", updatedStatus, "sidang", foundPengajuan.Data.Id, token)
		allSchedulesResponse = append(allSchedulesResponse, helper.ToScheduleDetailResponse(&schedule, &updatedData.Data))
		usersId = append(usersId, foundPengajuan.Data.UserId)
	}

	helper.CreateNotification(usersId, "Penjadwalan Sidang", "Jadwal sidang anda sudah ditetapkan!", "/schedule/get/mahasiswa", token)
	return &allSchedulesResponse, nil
}

func (s ScheduleServiceImpl) Update(request *web.ScheduleUpdateRequest, token string) (*[]web.ScheduleDetailResponse, error) {
	if err := s.Validator.Struct(request); err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	var pengajuansId []int

	foundSchedule, err := s.ScheduleRepository.FindById(request.Id)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	memberSchedules, err := s.ScheduleRepository.FindMemberSchedule(foundSchedule)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	if time.Now().After(request.DateTime) {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "time has passed")
	}

	validTime := helper.GetValidSidangTime()
	if request.DateTime.Before(validTime) || request.DateTime.Equal(validTime) {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "scheduling must be exactly 2 hours in advance")
	}

	if request.Penguji1.Id == request.Penguji2.Id {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "penguji 1 and penguji 2 must be different")
	}

	if request.Penguji1.Jfa != "NJFA" {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "penguji 1 must have jfa")
	}

	for _, pengajuanMember := range request.Members {
		if pengajuanMember.Kk != request.Penguji1.Kk && pengajuanMember.Kk != request.Penguji2.Kk {
			return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "have member with different kk (kompetensi keahlian)")
		}

		if pengajuanMember.Pembimbing1Id == request.Penguji1.Id || pengajuanMember.Pembimbing2Id == request.Penguji2.Id {
			return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "pengajuan with id "+strconv.Itoa(pengajuanMember.PengajuanId)+" penguji 1 and penguji 2 cant same with pembimbing 1")
		}

		if pengajuanMember.Pembimbing2Id == request.Penguji1.Id || pengajuanMember.Pembimbing2Id == request.Penguji2.Id {
			return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "pengajuan with id "+strconv.Itoa(pengajuanMember.PengajuanId)+" penguji 1 and penguji 2 cant same with pembimbing 2")
		}

		pengajuansId = append(pengajuansId, pengajuanMember.PengajuanId)
	}

	rooms, _ := s.ScheduleRepository.CheckAvailRoom(request.DateTime, request.Room, pengajuansId, "update")
	if len(rooms) > 0 {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "room is not available")
	}

	foundDataPenguji1, _ := s.ScheduleRepository.CheckAvailUser(request.DateTime, request.Penguji1.Id, "penguji1_id", pengajuansId, "update")
	if len(foundDataPenguji1) > 0 {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "penguji 1 is not available")
	}

	foundDataPenguji2, _ := s.ScheduleRepository.CheckAvailUser(request.DateTime, request.Penguji2.Id, "penguji2_id", pengajuansId, "update")
	if len(foundDataPenguji2) > 0 {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "penguji 2 is not available")
	}

	var scheduleUpdateData []entity.Schedule
	var usersId []int

	for _, schedule := range memberSchedules {
		schedule.DateTime = request.DateTime
		schedule.Room = request.Room
		schedule.Penguji1Id = request.Penguji1.Id
		schedule.Penguji2Id = request.Penguji2.Id

		scheduleUpdateData = append(scheduleUpdateData, schedule)
	}

	save, err := s.ScheduleRepository.UpdateMany(scheduleUpdateData)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	var allSchedulesResponse []web.ScheduleDetailResponse
	for _, schedule := range save {
		foundPengajuan := helper.GetPengajuanById(schedule.PengajuanId, token)
		allSchedulesResponse = append(allSchedulesResponse, helper.ToScheduleDetailResponse(&schedule, &foundPengajuan.Data))

		usersId = append(usersId, foundPengajuan.Data.UserId)
	}

	helper.CreateNotification(usersId, "Penjadwalan Sidang", "Jadwal sidang anda sudah berubah!", "/schedule/get/mahasiswa", token)
	return &allSchedulesResponse, nil
}

func (s ScheduleServiceImpl) Delete(scheduleId int, token string) error {
	foundSchedule, err := s.ScheduleRepository.FindById(scheduleId)
	if err != nil {
		return exception.NewErrorResponse(fiber.StatusNotFound, "data not found")
	}

	memberSchedules, err := s.ScheduleRepository.FindMemberSchedule(foundSchedule)
	if err != nil {
		return exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	err = s.ScheduleRepository.Delete(memberSchedules)
	if err != nil {
		return exception.NewErrorResponse(fiber.StatusBadRequest, "failed to delete schedule")
	}

	for _, schedule := range memberSchedules {
		foundPengajuan := helper.GetPengajuanById(schedule.PengajuanId, token)

		updatedStatus := "belum dijadwalkan"
		if foundPengajuan.Data.Status == "tidak lulus (sudah dijadwalkan)" {
			updatedStatus = "tidak lulus (belum dijadwalkan)"
		}

		_ = helper.ChangePengajuanStatus("-", updatedStatus, "penjadwalan", foundPengajuan.Data.Id, token)
	}

	return nil
}

func (s ScheduleServiceImpl) AddFlag(code string, scheduleId int) (*web.ScheduleResponse, error) {
	foundData, err := s.ScheduleRepository.FindById(scheduleId)
	if err != nil {
		fmt.Println(err)
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, "data not found")
	}

	switch code {
	case "rev":
		foundData.FlagAddRevision = true
	case "scr":
		foundData.FlagChangeScores = true
	default:
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "flag param only 'scr' and 'rev'")
	}

	update, err := s.ScheduleRepository.Update(foundData)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "failed to update flag")
	}

	response := helper.ToScheduleResponse(update)

	return &response, nil

}

func (s ScheduleServiceImpl) ChangeStatus(request *web.ScheduleUpdateStatusRequest) (*web.ScheduleResponse, error) {
	if err := s.Validator.Struct(request); err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	foundData, err := s.ScheduleRepository.FindById(request.Id)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, "data not found")
	}

	foundData.Status = request.Status
	update, err := s.ScheduleRepository.Update(foundData)
	if err != nil {
		return nil, err
	}

	response := helper.ToScheduleResponse(update)

	return &response, nil

}
