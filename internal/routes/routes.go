package routes

import (
	"fetch-saldo/internal/handler"
	"fetch-saldo/internal/middleware"
	"fetch-saldo/internal/model"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/api/get-balances",
		middleware.ParseBodyRequest[model.BalanceRequest],
		handler.GetBalance,
	)
}
