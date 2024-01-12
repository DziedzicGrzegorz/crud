package server

import (
	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	*fiber.App
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(),
	}

	return server
}
func GroupPath(pathToGroup string, app *fiber.App) fiber.Router {
	group := app.Group(pathToGroup)

	return group
}
