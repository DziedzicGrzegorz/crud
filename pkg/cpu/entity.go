package cpu

import "time"

type CpuEntity struct {
	Id            int       `json:"id"`
	Name          string    `json:"name"`
	MaxClockSpeed int       `json:"maxClockSpeed"`
	Cores         int       `json:"cores"`
	Threads       int       `json:"threads"`
	CreatedAt     time.Time `json:"createdAt"`
}
