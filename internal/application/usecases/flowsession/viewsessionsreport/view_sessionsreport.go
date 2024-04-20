package viewsessionsreport

import (
	"github.com/TristanSch1/flow/internal/application"
	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/domain/sessionsreport"
)

type UseCase struct {
	sessionRepository application.SessionRepository
}

func (s UseCase) Execute(
	command Command,
	presenter application.SessionsReportPresenter,
) error {
	var sessions []session.Session

	if command.Project == "" {
		allSessions, err := s.sessionRepository.FindAllSessions()
		if err != nil {
			return err
		}

		sessions = allSessions
	} else {
		sessionsForProject, err := s.sessionRepository.FindAllByProject(command.Project)
		if err != nil {
			return err
		}

		sessions = sessionsForProject
	}

	sessionsReport := sessionsreport.SessionsReport{
		Sessions: sessions,
	}

	if command.Format == sessionsreport.FormatByDay {
		presenter.ShowByDay(sessionsReport)
	} else {
		presenter.ShowByProject(sessionsReport)
	}

	return nil
}

func NewViewSessionsReportUseCase(sessionRepository application.SessionRepository) UseCase {
	return UseCase{
		sessionRepository: sessionRepository,
	}
}
