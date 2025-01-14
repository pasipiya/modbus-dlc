package routes

import (
	"modbus-dlc/api/v1/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controllers.Health)

	v1API := app.Group("/api/v1")

	SetupProductsRoutes(v1API)
	SetupDlcRoutes(v1API)
}
