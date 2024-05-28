package routes

import (
	"github.com/gofiber/fiber/v2"
	"pendaftaran-sidang-new/internal/controllers"
	"pendaftaran-sidang-new/internal/middleware"
)

func PengajuanRoutes(router fiber.Router, pengajuanController controllers.PengajuanController) {
	pengajuan := router.Group("/pengajuan").Use(middleware.UserAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))

	pengajuan.Get("/get", pengajuanController.FindAll)
	pengajuan.Get("/pembimbing/get", pengajuanController.FindAllPembimbing)
	pengajuan.Get("/pic/get", pengajuanController.FindAllPic)
	pengajuan.Get("/kk/get", pengajuanController.FindAllPengajuanKk)
	pengajuan.Get("/get/:pengajuanId", pengajuanController.FindPengajuanById)
	pengajuan.Get("/user/get", pengajuanController.FindPengajuanByUser)

	pengajuan.Post("/create", pengajuanController.Create)
	pengajuan.Patch("/update/:pengajuanId", pengajuanController.Update)
	pengajuan.Patch("/rejected/:pengajuanId", pengajuanController.PengajuanRejected)
	pengajuan.Patch("/approve/:pengajuanId", pengajuanController.PengajuanApprove)
	pengajuan.Patch("/change-status/:pengajuanId", pengajuanController.ChangeStatus)
	pengajuan.Delete("/delete/:pengajuanId", pengajuanController.Delete)

}
