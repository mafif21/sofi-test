package routes

import (
	"github.com/gofiber/fiber/v2"
	"pendaftaran-sidang-new/internal/controllers"
	"pendaftaran-sidang-new/internal/middleware"
)

func StatusLogRoutes(router fiber.Router, controller controllers.StatusLogController) {
	status := router.Group("/statuslog").Use(middleware.UserAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))

	status.Get("/get", controller.FindAll)
	status.Get("/get/:statusId", controller.FindStatusLogById)
	status.Get("/pengajuan/get/:pengajuanId", controller.FindAllByPengajuanId)
	status.Post("/create", controller.Create)
	status.Delete("/delete/:statusId", controller.Delete)

}
