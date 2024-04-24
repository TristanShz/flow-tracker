package list

import "github.com/TristanSch1/flow/internal/application"

type UseCase struct {
	sessionRepository application.SessionRepository
}

func (s UseCase) Execute() ([]string, error) {
	projects := s.sessionRepository.FindAllProjects()

	return projects, nil
}

func NewListProjectsUseCase(sessionRepository application.SessionRepository) UseCase {
	return UseCase{
		sessionRepository: sessionRepository,
	}
}
