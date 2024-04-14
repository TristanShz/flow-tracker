package allsessionsreport

import (
	"time"

	"github.com/TristanSch1/flow/internal/application"
	"github.com/TristanSch1/flow/internal/domain/sessionsreport"
)

type UseCase struct {
	sessionRepository application.SessionRepository
}

type AllSessionsReport struct {
	Projects         map[string]time.Duration
	Total            time.Duration
	NumberOfSessions int
}

func (s UseCase) Execute(command Command) (AllSessionsReport, error) {
	sessions, err := s.sessionRepository.FindAllSessions()
	if err != nil {
		return AllSessionsReport{}, err
	}

	sessionsReport := sessionsreport.SessionsReport{
		Sessions: sessions,
	}

	return AllSessionsReport{
		Projects:         sessionsReport.ProjectsReport(),
		Total:            sessionsReport.TotalDuration(),
		NumberOfSessions: len(sessions),
	}, nil
}

func NewFlowSessionsReportUseCase(sessionRepository application.SessionRepository) UseCase {
	return UseCase{
		sessionRepository: sessionRepository,
	}
}
