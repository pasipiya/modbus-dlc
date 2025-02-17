package routes

import (
	"github.com/gofiber/fiber/v2"

	"modbus-dlc/api/v1/controllers"
	mw "modbus-dlc/api/v1/middleware"
	C "modbus-dlc/constants"
)

func SetupProductsRoutes(router fiber.Router) {

	router.Get("/products", mw.RateLimit(C.Tier3, 0), controllers.GetProducts)
	router.Get("/products/:id", mw.RateLimit(C.Tier3, 0), controllers.GetProduct)

	router.Post("/products", mw.RateLimit(C.Tier2, 0), controllers.CreateProduct)

	router.Patch("/products/:id", mw.RateLimit(C.Tier2, 0), controllers.UpdateProduct)

	router.Delete("/products/:id", mw.RateLimit(C.Tier3, 0), controllers.DeleteProduct)

}
