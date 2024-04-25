package app

import (
	"github.com/TristanSch1/flow/internal/application"
	startsession "github.com/TristanSch1/flow/internal/application/usecases/flowsession/start"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/status"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/stopsession"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/viewsessionsreport"
	"github.com/TristanSch1/flow/internal/application/usecases/project/list"
)

type App struct {
	DateProvider              application.DateProvider
	StartFlowSessionUseCase   startsession.UseCase
	StopFlowSessionUseCase    stopsession.UseCase
	FlowSessionStatusUseCase  status.UseCase
	ListProjectsUseCase       list.UseCase
	ViewSessionsReportUseCase viewsessionsreport.UseCase
}

func NewApp(
	dateProvider application.DateProvider,
	startFlowSessionUseCase startsession.UseCase,
	stopFlowSessionUseCase stopsession.UseCase,
	flowSessionStatusUseCase status.UseCase,
	listProjectsUseCase list.UseCase,
	viewSessionsReportUseCase viewsessionsreport.UseCase,
) *App {
	return &App{
		DateProvider:              dateProvider,
		StartFlowSessionUseCase:   startFlowSessionUseCase,
		StopFlowSessionUseCase:    stopFlowSessionUseCase,
		FlowSessionStatusUseCase:  flowSessionStatusUseCase,
		ListProjectsUseCase:       listProjectsUseCase,
		ViewSessionsReportUseCase: viewSessionsReportUseCase,
	}
}
