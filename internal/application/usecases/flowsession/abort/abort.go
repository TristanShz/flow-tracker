package abortsession

import (
	"errors"

	"github.com/TristanSch1/flow/internal/application"
)

type UseCase struct {
	sessionRepository application.SessionRepository
}

func (s UseCase) Execute() error {
	lastSession := s.sessionRepository.FindLastSession()

	if lastSession == nil || !lastSession.EndTime.IsZero() {
		return ErrNoActiveSession
	}

	s.sessionRepository.Delete(lastSession.Id)

	return nil
}

var ErrNoActiveSession = errors.New("no active session")

func NewAbortFlowSessionUseCase(sessionRepository application.SessionRepository) UseCase {
	return UseCase{
		sessionRepository: sessionRepository,
	}
}
