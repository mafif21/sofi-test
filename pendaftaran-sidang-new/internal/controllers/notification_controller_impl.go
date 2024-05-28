package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
	"pendaftaran-sidang-new/internal/exception"
	"pendaftaran-sidang-new/internal/model/web"
	"pendaftaran-sidang-new/internal/services"
	"time"
)

type NotificationControllerImpl struct {
	Services services.NotificationService
}

func NewNotificationController(services services.NotificationService) NotificationController {
	return &NotificationControllerImpl{Services: services}
}

func (c NotificationControllerImpl) FindAll(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLSPR"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	filter := map[string]interface{}{}

	notification, err := c.Services.GetAllNotification(filter)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	response := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "success get all notification data",
		Data:   notification,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c NotificationControllerImpl) FindByUser(ctx *fiber.Ctx) error {
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

	filter := map[string]interface{}{
		"user_id": userId,
		"read_at": nil,
	}

	notification, err := c.Services.GetAllNotification(filter)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	response := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "success get all notification user data",
		Data:   notification,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c NotificationControllerImpl) Update(ctx *fiber.Ctx) error {
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

	notificationId := ctx.Params("notificationId")
	updateRequest := web.NotificationUpdateRequest{
		Id:     notificationId,
		ReadAt: time.Now(),
		UserId: userId,
	}

	updatedData, err := c.Services.Update(&updateRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	response := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "success read notification",
		Data:   updatedData,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c NotificationControllerImpl) Create(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC", "RLPBB", "RLPGJ", "RLDSN", "RLADM", "RLPPM", "RLSPR", "RLMHS"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	createRequest := &web.NotificationCreateRequest{}
	if err := ctx.BodyParser(&createRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "failed to parse json"))
	}

	newNotification, err := c.Services.Create(createRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "success add new notification",
		Data:   newNotification,
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)

}
