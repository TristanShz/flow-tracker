package app

import (
	"github.com/TristanSch1/flow/internal/application"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/sessionstatus"
	startsession "github.com/TristanSch1/flow/internal/application/usecases/flowsession/start"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/stopsession"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/viewsessionsreport"
	"github.com/TristanSch1/flow/internal/application/usecases/project/list"
)

type App struct {
	SessionRepository         application.SessionRepository
	DateProvider              application.DateProvider
	StartFlowSessionUseCase   startsession.UseCase
	StopFlowSessionUseCase    stopsession.UseCase
	FlowSessionStatusUseCase  sessionstatus.UseCase
	ListProjectsUseCase       list.UseCase
	ViewSessionsReportUseCase viewsessionsreport.UseCase
}

func NewApp(
	sessionRepository application.SessionRepository,
	dateProvider application.DateProvider,
	startFlowSessionUseCase startsession.UseCase,
	stopFlowSessionUseCase stopsession.UseCase,
	flowSessionStatusUseCase sessionstatus.UseCase,
	listProjectsUseCase list.UseCase,
	viewSessionsReportUseCase viewsessionsreport.UseCase,
) *App {
	return &App{
		SessionRepository:         sessionRepository,
		DateProvider:              dateProvider,
		StartFlowSessionUseCase:   startFlowSessionUseCase,
		StopFlowSessionUseCase:    stopFlowSessionUseCase,
		FlowSessionStatusUseCase:  flowSessionStatusUseCase,
		ListProjectsUseCase:       listProjectsUseCase,
		ViewSessionsReportUseCase: viewSessionsReportUseCase,
	}
}
