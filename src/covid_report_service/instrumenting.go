package covid_report_service

import (
	"context"
	"time"

	"github.com/esvm/if1015covidproject-api/src/covid_reports"
)

type instrumentingMiddleware struct {
	next CovidReportService
}

func (mw instrumentingMiddleware) InsertCovidReport(ctx context.Context, covidReport *covid_reports.CovidReport) (*covid_reports.CovidReport, error) {
	begin := time.Now()

	res, err := mw.next.InsertCovidReport(ctx, covidReport)

	InsertCovidReportTotal.Add(1)
	InsertCovidReportDuration.Observe(time.Since(begin).Seconds())
	return res, err
}

func (mw instrumentingMiddleware) GetCovidReports(ctx context.Context) ([]*covid_reports.CovidReport, error) {
	begin := time.Now()

	res, err := mw.next.GetCovidReports(ctx)

	GetCovidReportsTotal.Add(1)
	GetCovidReportsDuration.Observe(time.Since(begin).Seconds())
	return res, err
}
