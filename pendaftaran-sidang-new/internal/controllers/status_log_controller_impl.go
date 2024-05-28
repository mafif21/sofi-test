package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
	"pendaftaran-sidang-new/internal/exception"
	"pendaftaran-sidang-new/internal/model/web"
	"pendaftaran-sidang-new/internal/services"
)

type StatusLogControllerImpl struct {
	StatusLogService services.StatusLogService
	PengajuanService services.PengajuanService
}

func NewStatusLogController(statusLogService services.StatusLogService, pengajuanService services.PengajuanService) StatusLogController {
	return &StatusLogControllerImpl{StatusLogService: statusLogService, PengajuanService: pengajuanService}
}

func (c StatusLogControllerImpl) FindAll(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLSPR"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	filter := map[string]interface{}{}

	statusLogs, err := c.StatusLogService.GetAllStatusLogs(filter)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	response := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "success get all status logs",
		Data:   statusLogs,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c StatusLogControllerImpl) FindAllByPengajuanId(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLMHS"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	pengajuanId, err := ctx.ParamsInt("pengajuanId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "pengajuan id is not valid"))
	}

	userLoggedIn := ctx.Locals("user_id")
	userId := userLoggedIn.(int)

	userPengajuan, err := c.PengajuanService.GetPengajuanByUserId(userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	if userPengajuan.Id != pengajuanId {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "pengajuan id is not valid with user"))
	}

	filter := map[string]interface{}{
		"pengajuan_id": pengajuanId,
	}

	statusLogs, err := c.StatusLogService.GetAllStatusLogs(filter)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	response := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "success get all status logs",
		Data:   statusLogs,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c StatusLogControllerImpl) FindStatusLogById(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLSPR"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	statusId, err := ctx.ParamsInt("statusId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "status id is not valid"))
	}

	found, err := c.StatusLogService.GetStatusLogsById(statusId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	response := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "success get status log data",
		Data:   found,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c StatusLogControllerImpl) Create(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC", "RLPBB", "RLPGJ", "RLDSN", "RLADM", "RLPPM", "RLSPR", "RLMHS"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	userLoggedIn := ctx.Locals("user_id")
	userId := userLoggedIn.(int)

	statusLogRequest := &web.StatusLogCreateRequest{}
	if err := ctx.BodyParser(&statusLogRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "failed to parse json"))
	}

	statusLogRequest.CreatedBy = userId

	newStatusLogs, err := c.StatusLogService.Create(statusLogRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	response := &web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "success create new status log",
		Data:   newStatusLogs,
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c StatusLogControllerImpl) Delete(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLSPR"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	statusId, err := ctx.ParamsInt("statusId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "status id is not valid"))
	}

	err = c.StatusLogService.Delete(statusId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	response := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "success delete status logs",
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
