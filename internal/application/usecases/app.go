package app

import (
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/allsessionsreport"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/projectsessionsreport"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/start"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/status"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/stop"
	"github.com/TristanSch1/flow/internal/application/usecases/project/list"
)

type App struct {
	StartFlowSessionUseCase      start.UseCase
	StopFlowSessionUseCase       stop.UseCase
	FlowSessionStatusUseCase     status.UseCase
	ListProjectsUseCase          list.UseCase
	AllSessionsReportUseCase     allsessionsreport.UseCase
	ProjectSessionsReportUseCase projectsessionsreport.UseCase
}

func NewApp(
	startFlowSessionUseCase start.UseCase,
	stopFlowSessionUseCase stop.UseCase,
	flowSessionStatusUseCase status.UseCase,
	listProjectsUseCase list.UseCase,
	allSessionsReportUseCase allsessionsreport.UseCase,
	projectSessionsReportUseCase projectsessionsreport.UseCase,
) *App {
	return &App{
		StartFlowSessionUseCase:      startFlowSessionUseCase,
		StopFlowSessionUseCase:       stopFlowSessionUseCase,
		FlowSessionStatusUseCase:     flowSessionStatusUseCase,
		ListProjectsUseCase:          listProjectsUseCase,
		AllSessionsReportUseCase:     allSessionsReportUseCase,
		ProjectSessionsReportUseCase: projectSessionsReportUseCase,
	}
}
