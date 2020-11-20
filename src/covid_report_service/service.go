package covid_report_service

import (
	"context"

	"github.com/esvm/if1015covidproject-api/src/covid_report_service/store"
	"github.com/esvm/if1015covidproject-api/src/covid_reports"
	"github.com/go-kit/kit/log"
)

type CovidReportService interface {
	InsertCovidReport(context.Context, *covid_reports.CovidReport) (*covid_reports.CovidReport, error)
	GetCovidReports(context.Context) ([]*covid_reports.CovidReport, error)
}

type basicService struct {
	store store.Store
}

func NewCovidReportService(logger log.Logger) CovidReportService {
	var service CovidReportService
	service = basicService{store.New(logger)}
	service = loggingMiddleware{logger, service}
	service = instrumentingMiddleware{service}
	service = crashMonitoringMiddleware{service}
	return service
}

func (s basicService) InsertCovidReport(ctx context.Context, report *covid_reports.CovidReport) (*covid_reports.CovidReport, error) {
	if err := ValidateCovidReport(report); err != nil {
		return nil, err
	}

	return s.store.InsertCovidReport(report)
}

func (s basicService) GetCovidReports(ctx context.Context) ([]*covid_reports.CovidReport, error) {
	return s.store.GetCovidReports()
}
