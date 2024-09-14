package controllers

import (
	"github.com/gofiber/fiber/v2"

	S "github.com/atharvbhadange/go-api-template/api/v1/services"
	H "github.com/atharvbhadange/go-api-template/handler"
	// U "github.com/atharvbhadange/go-api-template/utils"
)

func GetDlc(ctx *fiber.Ctx) error {
	// dbTrx, txErr := U.StartNewPGTrx(ctx)

	// if txErr != nil {
	// 	return H.BuildError(ctx, "Unable to get transaction", fiber.StatusInternalServerError, txErr)
	// }

	// idInt, err := ctx.ParamsInt("id")

	// if err != nil {
	// 	return H.BuildError(ctx, "Invalid product id", fiber.StatusBadRequest, err)
	// }

	dlc, serviceErr := S.GetDlc()

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, fiber.Map{
		"ok":      1,
		"data": dlc,
	})
}