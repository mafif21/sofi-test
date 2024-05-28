package routes

import (
	"github.com/gofiber/fiber/v2"
	"penjadwalan-sidang-new/internal/controller"
	"penjadwalan-sidang-new/internal/middleware"
)

func NewRoutes(router fiber.Router, scheduleController controller.ScheduleController) {
	schedule := router.Group("/schedule").Use(middleware.UserAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    fiber.StatusUnauthorized,
				"message": "Unauthorized",
			})
		},
	}))

	schedule.Get("/get", scheduleController.FindAll)                                       //
	schedule.Get("/get/:scheduleId", scheduleController.FindById)                          //
	schedule.Get("/pengajuan/get/:pengajuanId", scheduleController.FindPengajuanSchedules) // success
	schedule.Get("/admin/get", scheduleController.FindAllAdmin)                            //
	schedule.Get("/admin-before/get", scheduleController.FindAllAdmin)                     //
	schedule.Get("/superadmin/get", scheduleController.FindAllSuperadmin)                  //
	schedule.Get("/penguji/get", scheduleController.FindAllPenguji)                        //
	schedule.Get("/pembimbing/get", scheduleController.FindAllPembimbing)                  //
	schedule.Get("/mahasiswa/get", scheduleController.FindAllMahasiswa)                    //
	schedule.Get("/list-embed/get/:pengajuanId", scheduleController.FindScheduleList)

	schedule.Post("/create", scheduleController.Create)
	schedule.Patch("/update/:scheduleId", scheduleController.Update)
	schedule.Patch("/change-flag/:scheduleId", scheduleController.AddFlag)
	schedule.Patch("/change-status/:scheduleId", scheduleController.ChangeStatus)
	schedule.Delete("/delete/:scheduleId", scheduleController.Delete)
}
