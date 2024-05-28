package controllers

import "github.com/gofiber/fiber/v2"

type TeamController interface {
	FindAll(ctx *fiber.Ctx) error
	FindTeamById(ctx *fiber.Ctx) error
	FindTeamByUserId(ctx *fiber.Ctx) error
	CreateTeam(ctx *fiber.Ctx) error
	CreatePersonal(ctx *fiber.Ctx) error
	AddMember(ctx *fiber.Ctx) error
	LeaveTeam(ctx *fiber.Ctx) error
	AvailableMember(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
