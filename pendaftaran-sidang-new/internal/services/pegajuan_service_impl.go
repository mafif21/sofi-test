package services

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"os"
	"path"
	"pendaftaran-sidang-new/internal/exception"
	"pendaftaran-sidang-new/internal/helper"
	"pendaftaran-sidang-new/internal/model/entity"
	"pendaftaran-sidang-new/internal/model/web"
	"pendaftaran-sidang-new/internal/repositories"
	"strconv"
)

type PengajuanServiceImpl struct {
	PengajuanRepository repositories.PengajuanRepository
	Validator           *validator.Validate
}

func NewPengajuanService(pengajuanRepository repositories.PengajuanRepository, validator *validator.Validate) PengajuanService {
	return &PengajuanServiceImpl{
		PengajuanRepository: pengajuanRepository,
		Validator:           validator,
	}
}

func (s PengajuanServiceImpl) GetAllPengajuan(filter map[string]interface{}) ([]web.PengajuanResponse, error) {
	allPengajuan, err := s.PengajuanRepository.FindAll(filter)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusInternalServerError, "internal server error")
	}

	var pengajuanResponse []web.PengajuanResponse
	for _, pengajuan := range allPengajuan {
		pengajuanResponse = append(pengajuanResponse, helper.ToPengajuanResponse(&pengajuan))
	}

	return pengajuanResponse, nil
}

func (s PengajuanServiceImpl) GetPengajuanById(pengajuanId int) (*web.PengajuanDetailResponse, error) {
	found, err := s.PengajuanRepository.FindPengajuanById(pengajuanId)

	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, "data not found")
	}

	response := helper.ToPengajuanDetailResponse(found)
	return &response, nil
}

func (s PengajuanServiceImpl) GetPengajuanByUserId(userId int) (*web.PengajuanDetailResponse, error) {
	found, err := s.PengajuanRepository.FindPengajuanByUserId(userId)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, "user dont have pengajuan data")
	}

	response := helper.ToPengajuanDetailResponse(found)
	return &response, nil
}

func (s PengajuanServiceImpl) Create(request *web.PengajuanCreateRequest) (*web.PengajuanResponse, error) {
	if err := s.Validator.Struct(request); err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	if request.Pembimbing1Id == request.Pembimbing2Id {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "cant same pembimbing 1 and pembimbing 2")
	}

	newPengajuan := &entity.Pengajuan{
		UserId:         request.UserId,
		Nim:            request.Nim,
		Pembimbing1Id:  request.Pembimbing1Id,
		Pembimbing2Id:  request.Pembimbing2Id,
		Judul:          request.Judul,
		Eprt:           request.Eprt,
		Tak:            request.Tak,
		PeriodID:       request.PeriodID,
		FormBimbingan1: request.FormBimbingan1,
		FormBimbingan2: request.FormBimbingan2,
	}

	newPengajuan.Kk = "Cybernetics"

	docTaUrl := fmt.Sprintf("/public/doc_ta/%s", request.DocTa)
	newPengajuan.DocTa = docTaUrl

	makalahUrl := fmt.Sprintf("/public/makalah/%s", request.Makalah)
	newPengajuan.Makalah = makalahUrl

	savedPengajuan, err := s.PengajuanRepository.Save(newPengajuan)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	response := helper.ToPengajuanResponse(savedPengajuan)
	return &response, nil
}

func (s PengajuanServiceImpl) Update(request *web.PengajuanUpdateRequest) (*web.PengajuanResponse, error) {
	if err := s.Validator.Struct(request); err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	oldData, err := s.PengajuanRepository.FindPengajuanById(request.Id)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	if request.Pembimbing1Id == request.Pembimbing2Id {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "cant same pembimbing 1 and pembimbing 2")
	}

	if request.DocTa != "" {
		_ = os.Remove("./public/doc_ta/" + path.Base(oldData.DocTa))
		oldData.DocTa = fmt.Sprintf("/public/doc_ta/%s", request.DocTa)
	} else {
		oldData.DocTa = oldData.DocTa
	}

	if request.Makalah != "" {
		_ = os.Remove("./public/makalah/" + path.Base(oldData.Makalah))
		oldData.Makalah = fmt.Sprintf("/public/makalah/%s", request.Makalah)
	} else {
		oldData.Makalah = oldData.Makalah
	}

	oldData.Nim = request.Nim
	oldData.Pembimbing1Id = request.Pembimbing1Id
	oldData.Pembimbing2Id = request.Pembimbing2Id
	oldData.Judul = request.Judul
	oldData.Eprt = request.Eprt
	oldData.Tak = request.Tak

	oldData.PeriodID = request.PeriodID
	oldData.FormBimbingan1 = request.FormBimbingan1
	oldData.FormBimbingan2 = request.FormBimbingan2

	updatedData, err := s.PengajuanRepository.Update(oldData)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "failed to update data")
	}

	response := helper.ToPengajuanResponse(updatedData)
	return &response, nil

}

func (s PengajuanServiceImpl) AdminStatus(request *web.StatusAdminUpdate) (*web.PengajuanResponse, error) {
	found, err := s.PengajuanRepository.FindPengajuanById(request.Id)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	var statusLogMessage string
	var notificationMessage string
	var title string
	var url string

	if request.Status == "accept" {
		statusLogMessage = "telah disetujui admin"
		found.Status = "telah disetujui admin"
		title = "Approved"
		notificationMessage = "Sidang anda telah di approve admin, silahkan membuat team"
		url = "team/get-team"
	} else {
		statusLogMessage = "ditolak oleh admin"
		found.Status = "ditolak oleh admin"
		notificationMessage = "Berkas Anda Ditolak Admin, Silahkan Perbaiki Berkas Anda"
		url = "pengajuan/update/" + strconv.Itoa(found.ID)
	}

	if request.Feedback == "" {
		request.Feedback = "-"
	}

	if request.IsEnglish {
		found.IsEnglish = true
	} else {
		found.IsEnglish = false
	}

	newStatusLog := &entity.StatusLog{
		Feedback:     request.Feedback,
		CreatedBy:    request.UserId,
		WorkFlowType: "pengajuan",
		Name:         statusLogMessage,
		PengajuanID:  found.ID,
	}

	newNotification := &entity.Notification{
		UserId:  found.UserId,
		Title:   title,
		Message: notificationMessage,
		Url:     url,
	}

	updated, err := s.PengajuanRepository.UpdateStatus(found, newStatusLog, newNotification)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	response := helper.ToPengajuanResponse(updated)
	return &response, nil
}

func (s PengajuanServiceImpl) ChangeStatus(request *web.ChangeStatusRequest) (*web.PengajuanResponse, error) {
	found, err := s.PengajuanRepository.FindPengajuanById(request.PengajuanId)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	found.Status = request.Status

	newStatusLog := &entity.StatusLog{
		Feedback:     request.Feedback,
		CreatedBy:    request.CreatedBy,
		WorkFlowType: request.WorkFlowType,
		Name:         request.Status,
		PengajuanID:  found.ID,
	}

	if request.Feedback == "" {
		request.Feedback = "-"
	}

	updated, err := s.PengajuanRepository.UpdateStatus(found, newStatusLog, nil)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	response := helper.ToPengajuanResponse(updated)
	return &response, nil
}

func (s PengajuanServiceImpl) Delete(pengajuanId int) error {
	found, err := s.PengajuanRepository.FindPengajuanById(pengajuanId)

	if err != nil {
		return exception.NewErrorResponse(fiber.StatusNotFound, "data not found")
	}

	err = s.PengajuanRepository.Delete(found)
	if err != nil {
		return exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	return nil
}
