package viewsessionsreport

import (
	"github.com/TristanSch1/flow/internal/application"
	"github.com/TristanSch1/flow/internal/domain/sessionsreport"
)

type UseCase struct {
	sessionRepository application.SessionRepository
}

func (s UseCase) Execute(
	command Command,
	presenter application.SessionsReportPresenter,
) error {
	sessions, err := s.sessionRepository.FindAllSessions()
	if err != nil {
		return err
	}

	presenter.Show(sessionsreport.SessionsReport{
		Sessions: sessions,
	})

	return nil
}

func NewViewSessionsReportUseCase(sessionRepository application.SessionRepository) UseCase {
	return UseCase{
		sessionRepository: sessionRepository,
	}
}
