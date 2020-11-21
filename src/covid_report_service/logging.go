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

func (mw loggingMiddleware) InsertCovidReports(ctx context.Context, covidReport []*covid_reports.CovidReport) error {
	begin := time.Now()

	err := mw.next.InsertCovidReports(ctx, covidReport)

	arguments := []interface{}{"method", "InsertCovidReports", "err", err, "took", time.Since(begin)}

	level.Debug(mw.logger).Log(arguments...)

	return err
}

func (mw loggingMiddleware) GetCovidReports(ctx context.Context) ([]*covid_reports.CovidReport, error) {
	begin := time.Now()

	res, err := mw.next.GetCovidReports(ctx)

	arguments := []interface{}{"method", "GetCovidReports", "err", err, "took", time.Since(begin)}

	level.Debug(mw.logger).Log(arguments...)

	return res, err
}
