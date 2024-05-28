package controller

import "github.com/gofiber/fiber/v2"

type ScheduleController interface {
	FindAll(ctx *fiber.Ctx) error
	FindAllAdmin(ctx *fiber.Ctx) error
	FindAllSuperadmin(ctx *fiber.Ctx) error
	FindAllPenguji(ctx *fiber.Ctx) error
	FindAllPembimbing(ctx *fiber.Ctx) error
	FindAllPic(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	FindPengajuanSchedules(ctx *fiber.Ctx) error
	FindAllMahasiswa(ctx *fiber.Ctx) error
	FindScheduleList(ctx *fiber.Ctx) error

	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	AddFlag(ctx *fiber.Ctx) error
	ChangeStatus(ctx *fiber.Ctx) error
}
