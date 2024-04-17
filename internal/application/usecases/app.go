package app

import (
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/start"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/status"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/stop"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/viewsessionsreport"
	"github.com/TristanSch1/flow/internal/application/usecases/project/list"
)

type App struct {
	StartFlowSessionUseCase   start.UseCase
	StopFlowSessionUseCase    stop.UseCase
	FlowSessionStatusUseCase  status.UseCase
	ListProjectsUseCase       list.UseCase
	ViewSessionsReportUseCase viewsessionsreport.UseCase
}

func NewApp(
	startFlowSessionUseCase start.UseCase,
	stopFlowSessionUseCase stop.UseCase,
	flowSessionStatusUseCase status.UseCase,
	listProjectsUseCase list.UseCase,
	viewSessionsReportUseCase viewsessionsreport.UseCase,
) *App {
	return &App{
		StartFlowSessionUseCase:   startFlowSessionUseCase,
		StopFlowSessionUseCase:    stopFlowSessionUseCase,
		FlowSessionStatusUseCase:  flowSessionStatusUseCase,
		ListProjectsUseCase:       listProjectsUseCase,
		ViewSessionsReportUseCase: viewSessionsReportUseCase,
	}
}
