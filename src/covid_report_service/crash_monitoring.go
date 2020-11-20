package covid_report_service

import (
	"context"

	"github.com/esvm/if1015covidproject-api/src/covid_reports"
	"github.com/rollbar/rollbar-go"
)

type crashMonitoringMiddleware struct {
	next CovidReportService
}

func (mw crashMonitoringMiddleware) InsertCovidReport(ctx context.Context, covidReport *covid_reports.CovidReport) (*covid_reports.CovidReport, error) {
	res, err := mw.next.InsertCovidReport(ctx, covidReport)
	if err != nil {
		rollbar.Error(err)
	}

	return res, err
}

func (mw crashMonitoringMiddleware) GetCovidReports(ctx context.Context) ([]*covid_reports.CovidReport, error) {
	res, err := mw.next.GetCovidReports(ctx)
	if err != nil {
		rollbar.Error(err)
	}

	return res, err
}
