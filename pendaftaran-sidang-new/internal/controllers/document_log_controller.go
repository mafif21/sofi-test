package controllers

import "github.com/gofiber/fiber/v2"

type DocumentLogController interface {
	FindAll(ctx *fiber.Ctx) error
	FindDocumentLogById(ctx *fiber.Ctx) error
	FindLatestDocumentLog(ctx *fiber.Ctx) error
	CreateSlide(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
