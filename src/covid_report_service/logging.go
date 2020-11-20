package covid_report_service

import (
	"context"
	"time"

	"github.com/esvm/if1015covidproject-api/src/covid_reports"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type loggingMiddleware struct {
	logger log.Logger
	next   CovidReportService
}

func (mw loggingMiddleware) InsertCovidReport(ctx context.Context, covidReport *covid_reports.CovidReport) (*covid_reports.CovidReport, error) {
	begin := time.Now()

	res, err := mw.next.InsertCovidReport(ctx, covidReport)

	arguments := []interface{}{"method", "InsertCovidReport", "err", err, "took", time.Since(begin)}

	level.Debug(mw.logger).Log(arguments...)

	return res, err
}

func (mw loggingMiddleware) GetCovidReports(ctx context.Context) ([]*covid_reports.CovidReport, error) {
	begin := time.Now()

	res, err := mw.next.GetCovidReports(ctx)

	arguments := []interface{}{"method", "GetCovidReports", "err", err, "took", time.Since(begin)}

	level.Debug(mw.logger).Log(arguments...)

	return res, err
}
