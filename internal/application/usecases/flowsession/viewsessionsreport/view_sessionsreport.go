package viewsessionsreport

import (
	"github.com/TristanShz/flow/internal/application"
	"github.com/TristanShz/flow/internal/domain/session"
	"github.com/TristanShz/flow/internal/domain/sessionsreport"
	"github.com/TristanShz/flow/pkg/timerange"
)

type UseCase struct {
	sessionRepository application.SessionRepository
}

func (s UseCase) Execute(
	command Command,
	presenter application.SessionsReportPresenter,
) error {
	var sessions []session.Session

	if !command.Since.IsZero() || !command.Until.IsZero() {
		sessionsInTimeRange := s.sessionRepository.FindInTimeRange(timerange.TimeRange{
			Since: command.Since,
			Until: command.Until,
		})
		sessions = sessionsInTimeRange
	} else if command.Project != "" {
		sessionsForProject := s.sessionRepository.FindAllByProject(command.Project)

		sessions = sessionsForProject
	} else {
		allSessions := s.sessionRepository.FindAllSessions()

		sessions = allSessions
	}

	sessionsReport := sessionsreport.SessionsReport{
		Sessions: sessions,
	}

	if command.Format == sessionsreport.FormatByProject {
		presenter.ShowByProject(sessionsReport)
	} else {
		presenter.ShowByDay(sessionsReport)
	}

	return nil
}

func NewViewSessionsReportUseCase(sessionRepository application.SessionRepository) UseCase {
	return UseCase{
		sessionRepository: sessionRepository,
	}
}
