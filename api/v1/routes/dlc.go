package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/atharvbhadange/go-api-template/api/v1/controllers"
	mw "github.com/atharvbhadange/go-api-template/api/v1/middleware"
	C "github.com/atharvbhadange/go-api-template/constants"
)

func SetupDlcRoutes(router fiber.Router) {

	router.Get("/dlc", mw.RateLimit(C.Tier3, 0), controllers.GetDlc)

}
