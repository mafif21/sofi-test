package routes

import (
	"github.com/gofiber/fiber/v2"
	"pendaftaran-sidang-new/internal/controllers"
	"pendaftaran-sidang-new/internal/middleware"
)

func DocumentLogRoutes(router fiber.Router, documentLogController controllers.DocumentLogController) {
	documentLog := router.Group("/documentlog").Use(middleware.UserAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))

	documentLog.Get("/get", documentLogController.FindAll)
	documentLog.Get("/get/:documentId", documentLogController.FindDocumentLogById)
	documentLog.Get("/pengajuan/get/:pengajuanId", documentLogController.FindLatestDocumentLog)
	documentLog.Post("/create/slide", documentLogController.CreateSlide) // done
	documentLog.Patch("/update/:documentId", documentLogController.Update)
	documentLog.Delete("/delete/:documentId", documentLogController.Delete)
}
