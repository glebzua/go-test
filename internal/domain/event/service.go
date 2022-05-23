package event

type Service interface {
	FindAll() ([]Event, error)
	FindOne(id uint64) (*Event, error)
}

type service struct {
	repo *Repository
}

func NewService(r *Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FindAll() ([]Event, error) {
	return (*s.repo).FindAll()
}

func (s *service) FindOne(id uint64) (*Event, error) {
	return (*s.repo).FindOne(id)
}
