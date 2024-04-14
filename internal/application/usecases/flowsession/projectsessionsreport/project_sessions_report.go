package projectsessionsreport

import (
	"time"

	"github.com/TristanSch1/flow/internal/application"
	"github.com/TristanSch1/flow/internal/domain/sessionsreport"
)

type UseCase struct {
	sessionRepository application.SessionRepository
}

type ProjectSessionReport struct {
	Total            time.Duration
	NumberOfSessions int
}

func (s UseCase) Execute(command Command) (ProjectSessionReport, error) {
	sessions, err := s.sessionRepository.FindAllByProject(command.Project)
	if err != nil {
		return ProjectSessionReport{}, err
	}

	sessionsReport := sessionsreport.SessionsReport{
		Sessions: sessions,
	}

	return ProjectSessionReport{
		Total:            sessionsReport.TotalDuration(),
		NumberOfSessions: len(sessions),
	}, nil
}

func NewProjectSessionsReportUseCase(sessionRepository application.SessionRepository) UseCase {
	return UseCase{
		sessionRepository: sessionRepository,
	}
}
