package covid_report_service

import (
	"context"

	"github.com/esvm/if1015covidproject-api/src/covid_report_service/store"
	"github.com/esvm/if1015covidproject-api/src/covid_reports"
	"github.com/go-kit/kit/log"
)

type CovidReportService interface {
	InsertCovidReports(context.Context, []*covid_reports.CovidReport) error
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

func (s basicService) InsertCovidReports(ctx context.Context, reports []*covid_reports.CovidReport) error {
	if err := ValidateCovidReports(reports); err != nil {
		return err
	}

	return s.store.InsertCovidReports(reports)
}

func (s basicService) GetCovidReports(ctx context.Context) ([]*covid_reports.CovidReport, error) {
	return s.store.GetCovidReports()
}
