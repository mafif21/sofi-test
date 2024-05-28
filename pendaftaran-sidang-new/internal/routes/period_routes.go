package routes

import (
	"github.com/gofiber/fiber/v2"
	"pendaftaran-sidang-new/internal/controllers"
	"pendaftaran-sidang-new/internal/middleware"
)

func PeriodRoutes(router fiber.Router, periodController controllers.PeriodController) {
	period := router.Group("/period").Use(middleware.UserAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))

	period.Get("/get", periodController.FindAll)
	period.Get("/check-period", periodController.CheckPeriod)
	period.Get("/get/:periodId", periodController.FindPeriodById)
	period.Post("/create", periodController.Create)
	period.Patch("/update/:periodId", periodController.Update)
	period.Delete("/delete/:periodId", periodController.Delete)
}
