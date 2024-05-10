package app

import (
	"github.com/TristanShz/flow/internal/application"
	abortsession "github.com/TristanShz/flow/internal/application/usecases/flowsession/abort"
	"github.com/TristanShz/flow/internal/application/usecases/flowsession/sessionstatus"
	startsession "github.com/TristanShz/flow/internal/application/usecases/flowsession/start"
	"github.com/TristanShz/flow/internal/application/usecases/flowsession/stopsession"
	"github.com/TristanShz/flow/internal/application/usecases/flowsession/viewsessionsreport"
	"github.com/TristanShz/flow/internal/application/usecases/project/list"
)

type App struct {
	SessionRepository         application.SessionRepository
	DateProvider              application.DateProvider
	StartFlowSessionUseCase   startsession.UseCase
	StopFlowSessionUseCase    stopsession.UseCase
	AbortFlowSessionUseCase   abortsession.UseCase
	FlowSessionStatusUseCase  sessionstatus.UseCase
	ListProjectsUseCase       list.UseCase
	ViewSessionsReportUseCase viewsessionsreport.UseCase
}

func NewApp(
	sessionRepository application.SessionRepository,
	dateProvider application.DateProvider,
	startFlowSessionUseCase startsession.UseCase,
	stopFlowSessionUseCase stopsession.UseCase,
	abortFlowSessionUseCase abortsession.UseCase,
	flowSessionStatusUseCase sessionstatus.UseCase,
	listProjectsUseCase list.UseCase,
	viewSessionsReportUseCase viewsessionsreport.UseCase,
) *App {
	return &App{
		SessionRepository:         sessionRepository,
		DateProvider:              dateProvider,
		StartFlowSessionUseCase:   startFlowSessionUseCase,
		StopFlowSessionUseCase:    stopFlowSessionUseCase,
		AbortFlowSessionUseCase:   abortFlowSessionUseCase,
		FlowSessionStatusUseCase:  flowSessionStatusUseCase,
		ListProjectsUseCase:       listProjectsUseCase,
		ViewSessionsReportUseCase: viewSessionsReportUseCase,
	}
}
