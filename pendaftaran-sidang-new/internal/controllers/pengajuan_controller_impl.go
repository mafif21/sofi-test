package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
	"path/filepath"
	"pendaftaran-sidang-new/internal/exception"
	"pendaftaran-sidang-new/internal/helper"
	"pendaftaran-sidang-new/internal/model/web"
	"pendaftaran-sidang-new/internal/services"
)

type PengajuanControllerImpl struct {
	PengajuanService services.PengajuanService
}

func NewPengajuanController(pengajuanService services.PengajuanService) PengajuanController {
	return &PengajuanControllerImpl{PengajuanService: pengajuanService}
}

func (c PengajuanControllerImpl) FindAll(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLSPR"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	filter := map[string]interface{}{
		"status": []string{"ditolak oleh admin", "belum disetujui admin", "telah disetujui admin", "belum dijadwalkan", "tidak lulus (sudah update dokumen)", "sudah dijadwalkan", "tidak luluc (sudah dijadwalkan)"},
	}

	pengajuanDatas, err := c.PengajuanService.GetAllPengajuan(filter)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all pengajuan data",
		Data:   pengajuanDatas,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c PengajuanControllerImpl) FindAllPembimbing(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPBB"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	userLoggedIn := 12

	filter := map[string]interface{}{
		"pembimbing": userLoggedIn,
	}

	pengajuanDatas, err := c.PengajuanService.GetAllPengajuan(filter)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all pengajuan data",
		Data:   pengajuanDatas,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c PengajuanControllerImpl) FindAllPic(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)

	rolesToCheck := []string{"RLPIC"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	filter := map[string]interface{}{
		"kk":     "Cybernetics",
		"status": []string{"belum dijadwalkan", "tidak lulus (belum dijadwalkan)"},
	}

	pengajuanDatas, err := c.PengajuanService.GetAllPengajuan(filter)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all pengajuan data",
		Data:   pengajuanDatas,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c PengajuanControllerImpl) FindAllPengajuanKk(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC", "RLPBB", "RLPGJ", "RLDSN", "RLADM", "RLPPM", "RLSPR"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	filter := map[string]interface{}{
		"kk": "Cybernetics",
	}

	pengajuanDatas, err := c.PengajuanService.GetAllPengajuan(filter)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all pengajuan data",
		Data:   pengajuanDatas,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c PengajuanControllerImpl) FindPengajuanById(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLPBB", "RLSPR", "RLPIC"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	pengajuanId, err := ctx.ParamsInt("pengajuanId")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "pengajuan id is not valid",
		})
	}

	foundData, err := c.PengajuanService.GetPengajuanById(pengajuanId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "pengajuan has found",
		Data:   foundData,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c PengajuanControllerImpl) FindPengajuanByUser(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLMHS"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	userLoggedIn := ctx.Locals("user_id")
	userId := userLoggedIn.(int)

	foundData, err := c.PengajuanService.GetPengajuanByUserId(userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "pengajuan has found",
		Data:   foundData,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c PengajuanControllerImpl) Create(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLMHS"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	studentLoggedIn := ctx.Locals("user_id")
	userId := studentLoggedIn.(int)

	pengajuanRequest := web.PengajuanCreateRequest{}

	pengajuanRequest.UserId = userId
	if err := ctx.BodyParser(&pengajuanRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "failed to parse JSON",
		})
	}

	docTa, errDocTa := ctx.FormFile("doc_ta")
	makalah, errMakalah := ctx.FormFile("makalah")

	const maxFileSize = 5 * 1024 * 1024
	if docTa.Size > maxFileSize || makalah.Size > maxFileSize {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "file size exceeds 5MB"))
	}

	if docTa == nil || makalah == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "must upload document",
		})
	}

	allowedExtensions := map[string]bool{
		".pdf":  true,
		".doc":  true,
		".docx": true,
	}

	docTaExt := filepath.Ext(docTa.Filename)
	makalahExt := filepath.Ext(makalah.Filename)

	if !allowedExtensions[docTaExt] || !allowedExtensions[makalahExt] {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Only .pdf, .doc, or .docx files are allowed",
		})
	}

	docTaFileName, errDocTaFileName := helper.FileHandler(errDocTa, docTa)
	makalahFileName, errMakalahFileName := helper.FileHandler(errMakalah, makalah)

	if errDocTaFileName != nil || docTaFileName == "" || errMakalahFileName != nil || makalahFileName == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "file not valid",
		})
	}

	docTa.Filename = docTaFileName
	pengajuanRequest.DocTa = docTa.Filename

	makalah.Filename = makalahFileName
	pengajuanRequest.Makalah = makalah.Filename

	newPengajuan, err := c.PengajuanService.Create(&pengajuanRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	_ = ctx.SaveFile(docTa, fmt.Sprintf("./public/doc_ta/%s", docTa.Filename))
	_ = ctx.SaveFile(makalah, fmt.Sprintf("./public/makalah/%s", makalah.Filename))

	webResponse := web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "pengajuan has been created",
		Data:   newPengajuan,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (c PengajuanControllerImpl) Update(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLMHS", "RLADM", "RLPIC", "RLPBB", "RLPGJ", "RLSPR"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	pengajuanId, err := ctx.ParamsInt("pengajuanId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "id is not valid",
		})
	}

	pengajuanRequest := web.PengajuanUpdateRequest{}
	pengajuanRequest.Id = pengajuanId

	allowedExtensions := map[string]bool{
		".pdf":  true,
		".doc":  true,
		".docx": true,
	}

	if err := ctx.BodyParser(&pengajuanRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "failed to parse JSON",
		})
	}

	docTa, docTaErr := ctx.FormFile("doc_ta")

	if docTa != nil {
		const maxFileSize = 5 * 1024 * 1024
		if docTa.Size > maxFileSize {
			return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "file size exceeds 5MB"))
		}

		docTaExt := filepath.Ext(docTa.Filename)
		if !allowedExtensions[docTaExt] {
			return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "Only .pdf, .doc, or .docx files are allowed",
			})
		}

		docTaNewFilename, err := helper.FileHandler(docTaErr, docTa)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "file not valid",
			})
		}

		docTa.Filename = docTaNewFilename
		pengajuanRequest.DocTa = docTa.Filename
	}

	makalah, makalahErr := ctx.FormFile("makalah")

	if makalah != nil {
		const maxFileSize = 5 * 1024 * 1024
		if makalah.Size > maxFileSize {
			return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "file size exceeds 5MB"))
		}

		makalahExt := filepath.Ext(makalah.Filename)
		if !allowedExtensions[makalahExt] {
			return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "Only .pdf, .doc, or .docx files are allowed",
			})
		}

		makalahNewFileName, err := helper.FileHandler(makalahErr, makalah)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "file not valid",
			})
		}

		makalah.Filename = makalahNewFileName
		pengajuanRequest.Makalah = makalah.Filename
	}

	updatedData, err := c.PengajuanService.Update(&pengajuanRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	if makalah != nil {
		_ = ctx.SaveFile(makalah, fmt.Sprintf("./public/makalah/%s", makalah.Filename))
	}

	if docTa != nil {
		_ = ctx.SaveFile(docTa, fmt.Sprintf("./public/doc_ta/%s", docTa.Filename))
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "pengajuan has been updated",
		Data:   updatedData,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (c PengajuanControllerImpl) PengajuanRejected(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	adminLoggedIn := ctx.Locals("user_id")
	adminId := adminLoggedIn.(int)
	pengajuanId, err := ctx.ParamsInt("pengajuanId")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "pengajuan id is not valid",
		})
	}

	statusRequest := &web.StatusAdminUpdate{
		Id:     pengajuanId,
		UserId: adminId,
		Status: "rejected",
	}

	if err := ctx.BodyParser(&statusRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "failed to parse JSON",
		})
	}

	updatedAdminStatus, err := c.PengajuanService.AdminStatus(statusRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "pengajuan has been rejected by admin",
		Data:   updatedAdminStatus,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (c PengajuanControllerImpl) PengajuanApprove(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	adminLoggedIn := ctx.Locals("user_id")
	adminId := adminLoggedIn.(int)
	pengajuanId, err := ctx.ParamsInt("pengajuanId")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "pengajuan id is not valid",
		})
	}

	statusRequest := &web.StatusAdminUpdate{
		Id:     pengajuanId,
		UserId: adminId,
		Status: "accept",
	}

	if err := ctx.BodyParser(&statusRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "failed to parse JSON",
		})
	}

	updatedAdminStatus, err := c.PengajuanService.AdminStatus(statusRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "pengajuan has been approve by admin",
		Data:   updatedAdminStatus,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (c PengajuanControllerImpl) ChangeStatus(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC", "RLPBB", "RLPGJ", "RLDSN", "RLADM", "RLPPM", "RLSPR", "RLMHS"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	userLoggedIn := ctx.Locals("user_id")
	userId := userLoggedIn.(int)

	pengajuanId, err := ctx.ParamsInt("pengajuanId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "pengajuan id is not valid",
		})
	}

	changeStatusRequest := &web.ChangeStatusRequest{}
	if err := ctx.BodyParser(&changeStatusRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "failed to parse JSON",
		})
	}

	changeStatusRequest.PengajuanId = pengajuanId
	changeStatusRequest.CreatedBy = userId

	updatedData, err := c.PengajuanService.ChangeStatus(changeStatusRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "pengajuan status has been update",
		Data:   updatedData,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (c PengajuanControllerImpl) Delete(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLSPR"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	pengajuanId, err := ctx.ParamsInt("pengajuanId")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "pengajuan id is not valid",
		})
	}

	err = c.PengajuanService.Delete(pengajuanId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "success to delete pengajuan",
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
