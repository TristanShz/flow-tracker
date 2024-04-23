package start

import (
	"errors"
	"strconv"

	"github.com/TristanSch1/flow/internal/application"
	"github.com/TristanSch1/flow/internal/domain/session"
)

type UseCase struct {
	sessionRepository application.SessionRepository
	dateProvider      application.DateProvider
	idProvider        application.IDProvider
}

func (s UseCase) Execute(command Command) error {
	lastSession := s.sessionRepository.FindLastSession()

	if lastSession != nil && lastSession.EndTime.IsZero() {
		return ErrSessionAlreadyStarted
	}

	startTime := s.dateProvider.GetNow()
	session := session.Session{
		Id:        strconv.FormatInt(startTime.Unix(), 10),
		StartTime: startTime,
		Project:   command.Project,
		Tags:      command.Tags,
	}

	s.sessionRepository.Save(session)

	return nil
}

var ErrSessionAlreadyStarted = errors.New("there is already a session in progress")

func NewStartFlowSessionUseCase(
	sessionRepository application.SessionRepository,
	dateProvider application.DateProvider,
	idProvider application.IDProvider,
) UseCase {
	return UseCase{
		sessionRepository: sessionRepository,
		dateProvider:      dateProvider,
		idProvider:        idProvider,
	}
}
