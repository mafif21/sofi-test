package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
	"penjadwalan-sidang-new/internal/config"
	"penjadwalan-sidang-new/internal/controller"
	"penjadwalan-sidang-new/internal/repositories"
	"penjadwalan-sidang-new/internal/routes"
	"penjadwalan-sidang-new/internal/services"
)

func StartApp() {
	app := fiber.New()

	app.Use(cors.New())
	validator := validator.New()

	db, err := config.OpenConnection()
	if err != nil {
		panic(err)
	}

	scheduleRepositories := repositories.NewScheduleRepository(db)
	scheduleServices := services.NewScheduleService(scheduleRepositories, validator)
	scheduleController := controller.NewScheduleController(scheduleServices)

	api := app.Group("/api")
	routes.NewRoutes(api, scheduleController)

	err = app.Listen(":" + os.Getenv("APP_PORT"))
	if err != nil {
		panic(err)
	}
}
