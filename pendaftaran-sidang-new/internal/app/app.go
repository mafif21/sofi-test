package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
	"pendaftaran-sidang-new/internal/config"
	"pendaftaran-sidang-new/internal/controllers"
	"pendaftaran-sidang-new/internal/repositories"
	"pendaftaran-sidang-new/internal/routes"
	"pendaftaran-sidang-new/internal/services"
)

func StartApp() {
	app := fiber.New()

	app.Static("/public/doc_ta", "./public/doc_ta")
	app.Static("/public/makalah", "./public/makalah")
	app.Static("/public/slides", "./public/slides")

	app.Use(cors.New())
	validator := validator.New()

	db, err := config.OpenConnection()
	if err != nil {
		panic(err)
	}

	periodRepositories := repositories.NewPeriodRepository(db)
	periodServices := services.NewPeriodService(periodRepositories, validator)
	periodControllers := controllers.NewPeriodController(periodServices)

	pengajuanRepositories := repositories.NewPengajuanRepository(db)
	pengajuanServices := services.NewPengajuanService(pengajuanRepositories, validator)
	pengajuanControllers := controllers.NewPengajuanController(pengajuanServices)

	documentLogRepositories := repositories.NewDocumentLogRepository(db)
	documentLogServices := services.NewDocumentLogService(documentLogRepositories, pengajuanRepositories, validator)
	documentLogControllers := controllers.NewDocumentLogController(documentLogServices, pengajuanServices)

	teamRepositories := repositories.NewTeamRepository(db)
	teamServices := services.NewTeamService(teamRepositories, pengajuanRepositories, documentLogRepositories, validator)
	teamControllers := controllers.NewTeamContoller(teamServices, pengajuanServices)

	statusLogRepositories := repositories.NewStatusLogRepository(db)
	statusLogServices := services.NewStatusLogService(statusLogRepositories, pengajuanRepositories, validator)
	statusLogControllers := controllers.NewStatusLogController(statusLogServices, pengajuanServices)

	notificationRepositories := repositories.NewNotificationRepository(db)
	notificationServices := services.NewNotificationService(notificationRepositories, validator)
	notificationControllers := controllers.NewNotificationController(notificationServices)

	api := app.Group("/api")
	routes.PeriodRoutes(api, periodControllers)
	routes.DocumentLogRoutes(api, documentLogControllers)
	routes.PengajuanRoutes(api, pengajuanControllers)
	routes.TeamRoutes(api, teamControllers)
	routes.StatusLogRoutes(api, statusLogControllers)
	routes.NotificationRoutes(api, notificationControllers)

	err = app.Listen(":" + os.Getenv("APP_PORT"))
	if err != nil {
		panic(err)
	}
}
