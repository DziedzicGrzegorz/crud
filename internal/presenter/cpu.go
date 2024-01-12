package cpuPresenter

import (
	"crud/pkg/cpu"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status bool            `json:"status"`
	Data   []cpu.CpuEntity `json:"data"`
	Error  string          `json:"error,omitempty"`
}

var jsonStr = `{
    "status": true,
    "data": [
        {
            "id": 1,
            "name": "CpuEntity 1",
            "maxClockSpeed": 3500,
            "cores": 4,
            "threads": 8,
            "createdAt": "2022-01-01T00:00:00Z"
        },
        {
            "id": 2,
            "name": "CpuEntity 2",
            "maxClockSpeed": 4000,
            "cores": 6,
            "threads": 12,
            "createdAt": "2022-01-01T00:00:00Z"
        }
    ],
    "error": null
}`

// cpuSuccessResponse is the singular SuccessResponse that will be passed in the response by
// Handler
func CpuSuccessResponse(data []cpu.CpuEntity) *fiber.Map {

	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// cpuErrorResponse is the ErrorResponse that will be passed in the response by Handler
func CpuErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
