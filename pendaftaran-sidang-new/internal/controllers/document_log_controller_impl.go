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
	"strconv"
)

type DocumentLogControllerImpl struct {
	DocumentLogService services.DocumentLogService
	PengajuanService   services.PengajuanService
}

func NewDocumentLogController(documentLogService services.DocumentLogService, pengajuanService services.PengajuanService) DocumentLogController {
	return &DocumentLogControllerImpl{
		DocumentLogService: documentLogService,
		PengajuanService:   pengajuanService,
	}
}

func (c DocumentLogControllerImpl) FindAll(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLSPR"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	order := ctx.Query("order", "desc")
	page := ctx.Query("pages", "1")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "value is invalid"))
	}

	limit := ctx.Query("limit", "10")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "value is invalid"))
	}

	if pageInt < 1 {
		pageInt = 1
	}

	var filter = make(map[string]interface{})

	pengajuanId := ctx.Query("pengajuan_id")
	if pengajuanId != "" {
		filter["pengajuan_id"] = pengajuanId
	}

	fileType := ctx.Query("type")
	if fileType != "" {
		filter["type"] = fileType
	}

	documents, err := c.DocumentLogService.GetAllDocumentLog(filter, order, limitInt, pageInt)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all document logs",
		Data:   documents,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)

}

func (c DocumentLogControllerImpl) FindDocumentLogById(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLSPR", "RLMHS"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	documentId, err := ctx.ParamsInt("documentId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "document id is not valid"))
	}

	foundDocument, err := c.DocumentLogService.GetDocumentLogById(documentId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	response := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "success get document",
		Data:   foundDocument,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c DocumentLogControllerImpl) FindLatestDocumentLog(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLSPR"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	fileType := ctx.Query("type", "slide")

	pengajuanId, err := ctx.ParamsInt("pengajuanId")
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	filter := map[string]interface{}{
		"type": fileType,
	}

	document, err := c.DocumentLogService.GetLatestDocument(filter, pengajuanId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	response := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "success get latest document",
		Data:   document,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c DocumentLogControllerImpl) CreateSlide(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLMHS"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	studentLoggedIn := ctx.Locals("user_id")
	userId := studentLoggedIn.(int)

	foundPengajuan, err := c.PengajuanService.GetPengajuanByUserId(userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	if foundPengajuan.Status == "belum disetujui admin" || foundPengajuan.Status == "ditolak oleh admin" {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "make sure student status 'telah disetujui admin'"))
	}

	requestSlide := &web.DocumentLogCreateRequest{
		PengajuanId: foundPengajuan.Id,
		Type:        "slide",
		CreatedBy:   foundPengajuan.UserId,
	}

	slide, err := ctx.FormFile("slide")
	slideName, errSlideName := helper.FileHandler(err, slide)

	if errSlideName != nil || slideName == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "file not valid"))
	}

	ext := filepath.Ext(slide.Filename)
	if ext != ".ppt" && ext != ".pptx" {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "extension file is not valid"))
	}

	slide.Filename = slideName
	requestSlide.FileName = slide.Filename
	requestSlide.FileUrl = fmt.Sprintf("/public/slides/%s", slide.Filename)

	savedDocument, err := c.DocumentLogService.Create(requestSlide)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	response := &web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "slide has been upload",
		Data:   savedDocument,
	}

	_ = ctx.SaveFile(slide, fmt.Sprintf("./public/slides/%s", slide.Filename))

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c DocumentLogControllerImpl) Update(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLSPR"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	documentId, err := ctx.ParamsInt("documentId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "period id is not valid"))
	}

	documentRequest := &web.DocumentLogUpdateRequest{}
	documentRequest.Id = documentId

	if err := ctx.BodyParser(&documentRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "failed to parse json"))
	}

	updatedData, err := c.DocumentLogService.Update(documentRequest)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "document type has been updated",
		Data:   updatedData,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)

}

func (c DocumentLogControllerImpl) Delete(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLSPR"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	documentId, err := ctx.ParamsInt("documentId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "period id is not valid"))
	}

	err = c.DocumentLogService.Delete(documentId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "success delete document",
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
