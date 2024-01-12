package cpuHandlers

import (
	cpuPresenter "crud/internal/presenter"
	"crud/pkg/cpu"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GetCpus is handler/controller which lists all Cpus from the CpuShop
func GetCpus(service cpu.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.ReadCpus()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(cpuPresenter.CpuErrorResponse(err))
		}
		return c.JSON(cpuPresenter.CpuSuccessResponse(fetched))
	}
}
