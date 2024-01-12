package cpu

import "time"

//here we handle with database
//now we use mock data
// Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	ReadCpuEntity() ([]CpuEntity, error)
}
type repository struct {
	Collection []CpuEntity
}

func (r *repository) ReadCpuEntity() ([]CpuEntity, error) {
	return r.Collection, nil
}

func NewRepository() Repository {
	return &repository{
		Collection: exampleCpuEntity(),
	}
}

func exampleCpuEntity() []CpuEntity {
	return []CpuEntity{
		{
			Id:            1,
			Name:          "CpuEntity 1",
			MaxClockSpeed: 3500,
			Cores:         4,
			Threads:       8,
			CreatedAt:     time.Now(),
		},
		{
			Id:            2,
			Name:          "CpuEntity 2",
			MaxClockSpeed: 4000,
			Cores:         6,
			Threads:       12,
			CreatedAt:     time.Now(),
		},
		{
			Id:            3,
			Name:          "CpuEntity 3",
			MaxClockSpeed: 4500,
			Cores:         8,
			Threads:       16,
			CreatedAt:     time.Now(),
		},
		{
			Id:            4,
			Name:          "CpuEntity 4",
			MaxClockSpeed: 5000,
			Cores:         10,
			Threads:       20,
			CreatedAt:     time.Now(),
		},
		{
			Id:            5,
			Name:          "CpuEntity 5",
			MaxClockSpeed: 5500,
			Cores:         12,
			Threads:       24,
			CreatedAt:     time.Now(),
		},
	}
}
