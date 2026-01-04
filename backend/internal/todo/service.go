package todo

import "errors"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTask(title string) (Task, error) {
	if title == "" {
		return Task{}, errors.New("title cannot be empty")
	}

	return s.repo.Create(title)
}
