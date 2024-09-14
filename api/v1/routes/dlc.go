package routes

import (
	"github.com/gofiber/fiber/v2"

	"modbus-dlc/api/v1/controllers"
	mw "modbus-dlc/api/v1/middleware"
	C "modbus-dlc/constants"
)

func SetupDlcRoutes(router fiber.Router) {

	router.Get("/dlc", mw.RateLimit(C.Tier3, 0), controllers.GetDlc)

}
