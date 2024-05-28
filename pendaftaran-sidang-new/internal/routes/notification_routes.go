package routes

import (
	"github.com/gofiber/fiber/v2"
	"pendaftaran-sidang-new/internal/controllers"
	"pendaftaran-sidang-new/internal/middleware"
)

func NotificationRoutes(router fiber.Router, controller controllers.NotificationController) {
	notification := router.Group("/notification").Use(middleware.UserAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))

	notification.Get("/get", controller.FindAll)
	notification.Get("/user/get", controller.FindByUser)
	notification.Patch("/update/:notificationId", controller.Update)
	notification.Post("/create", controller.Create)

}
