package viewsessionsreport

import (
	"github.com/TristanSch1/flow/internal/application"
	"github.com/TristanSch1/flow/internal/domain/sessionsreport"
)

type UseCase struct {
	sessionRepository application.SessionRepository
}

func (s UseCase) Execute() (sessionsreport.SessionsReport, error) {
	sessions, err := s.sessionRepository.FindAllSessions()
	if err != nil {
		return sessionsreport.SessionsReport{}, err
	}

	return sessionsreport.SessionsReport{
		Sessions: sessions,
	}, nil
}

func NewViewSessionsReportUseCase(sessionRepository application.SessionRepository) UseCase {
	return UseCase{
		sessionRepository: sessionRepository,
	}
}
