package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
	"pendaftaran-sidang-new/internal/exception"
	"pendaftaran-sidang-new/internal/model/web"
	"pendaftaran-sidang-new/internal/services"
	"strconv"
)

type PeriodControllerImpl struct {
	PeriodService services.PeriodService
}

func NewPeriodController(periodService services.PeriodService) PeriodController {
	return &PeriodControllerImpl{
		PeriodService: periodService,
	}
}

func (c PeriodControllerImpl) FindAll(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLPPM", "RLMHS", "RLSPR"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	filter := map[string]interface{}{}

	allPeriods, err := c.PeriodService.GetAllPeriod(filter)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all period data",
		Data:   allPeriods,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c PeriodControllerImpl) FindPeriodById(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLPPM", "RLSPR"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	periodId, err := ctx.ParamsInt("periodId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "period id is not valid"))
	}

	foundData, err := c.PeriodService.GetPeriodById(periodId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "period has found",
		Data:   foundData,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c PeriodControllerImpl) Delete(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLPPM", "RLSPR"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	periodId, err := ctx.ParamsInt("periodId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "period id is not valid"))
	}

	err = c.PeriodService.Delete(periodId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "period with id " + strconv.Itoa(periodId) + " has been deleted",
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c PeriodControllerImpl) Create(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLPPM", "RLSPR"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	periodRequest := &web.PeriodCreateRequest{}
	if err := ctx.BodyParser(&periodRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "failed to parse json"))
	}

	newPeriod, err := c.PeriodService.Create(periodRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "new period has been created",
		Data:   newPeriod,
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c PeriodControllerImpl) Update(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLPPM", "RLSPR"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	periodId, err := ctx.ParamsInt("periodId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "period id is not valid"))
	}

	periodRequest := &web.PeriodUpdateRequest{}
	if err := ctx.BodyParser(&periodRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "failed to parse json"))
	}

	periodRequest.Id = periodId
	newPeriod, err := c.PeriodService.Update(periodRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "period has been updated",
		Data:   newPeriod,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c PeriodControllerImpl) CheckPeriod(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC", "RLPBB", "RLPGJ", "RLDSN", "RLMHS", "RLADM", "RLPPM", "RLSPR"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	period, err := c.PeriodService.GetPeriodByTime()
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "time is valid",
		Data:   period,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}
