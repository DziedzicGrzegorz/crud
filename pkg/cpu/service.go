package cpu

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	ReadCpus() ([]CpuEntity, error)
}

type service struct {
	repository Repository
}

// NewService is used to create a single instance of the service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// FetchBooks is a service layer that helps fetch all books in BookShop
func (s *service) ReadCpus() ([]CpuEntity, error) {
	return s.repository.ReadCpuEntity()
}
