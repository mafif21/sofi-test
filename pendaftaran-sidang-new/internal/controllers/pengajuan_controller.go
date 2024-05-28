package controllers

import "github.com/gofiber/fiber/v2"

type PengajuanController interface {
	FindAll(ctx *fiber.Ctx) error
	FindAllPembimbing(ctx *fiber.Ctx) error
	FindAllPic(ctx *fiber.Ctx) error
	FindAllPengajuanKk(ctx *fiber.Ctx) error
	FindPengajuanById(ctx *fiber.Ctx) error
	FindPengajuanByUser(ctx *fiber.Ctx) error

	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	PengajuanRejected(ctx *fiber.Ctx) error
	PengajuanApprove(ctx *fiber.Ctx) error
	ChangeStatus(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
