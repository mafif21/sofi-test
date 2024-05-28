package controllers

import "github.com/gofiber/fiber/v2"

type StatusLogController interface {
	FindAll(ctx *fiber.Ctx) error
	FindAllByPengajuanId(ctx *fiber.Ctx) error
	FindStatusLogById(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
