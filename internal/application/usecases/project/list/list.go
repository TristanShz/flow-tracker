package list

import "github.com/TristanSch1/flow/internal/application"

type UseCase struct {
	sessionRepository application.SessionRepository
}

func (s UseCase) Execute() ([]string, error) {
	projects, err := s.sessionRepository.FindAllProjects()
	if err != nil {
		return []string{}, err
	}

	return projects, nil
}

func NewListProjectsUseCase(sessionRepository application.SessionRepository) UseCase {
	return UseCase{
		sessionRepository: sessionRepository,
	}
}
