package routes

import (
	cpuHandlers "crud/internal/handlers"
	"crud/pkg/cpu"

	"github.com/gofiber/fiber/v2"
)

// here decalre all http method
func CpuRouter(app fiber.Router, service cpu.Service) {
	app.Get("/", cpuHandlers.GetCpus(service))
}
