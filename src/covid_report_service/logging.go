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

func (mw loggingMiddleware) InsertCovidReportsBrazil(ctx context.Context, covidReport []*covid_reports.CovidReportBrazilState) error {
	begin := time.Now()

	err := mw.next.InsertCovidReportsBrazil(ctx, covidReport)

	arguments := []interface{}{"method", "InsertCovidReportsBrazil", "err", err, "took", time.Since(begin)}

	level.Debug(mw.logger).Log(arguments...)

	return err
}

func (mw loggingMiddleware) InsertCovidReportsCountries(ctx context.Context, covidReport []*covid_reports.CovidReportCountry) error {
	begin := time.Now()

	err := mw.next.InsertCovidReportsCountries(ctx, covidReport)

	arguments := []interface{}{"method", "InsertCovidReportsCountries", "err", err, "took", time.Since(begin)}

	level.Debug(mw.logger).Log(arguments...)

	return err
}

func (mw loggingMiddleware) GetCovidReportsBrazil(ctx context.Context) ([]*covid_reports.CovidReportBrazilState, error) {
	begin := time.Now()

	res, err := mw.next.GetCovidReportsBrazil(ctx)

	arguments := []interface{}{"method", "GetCovidReportsBrazil", "err", err, "took", time.Since(begin)}

	level.Debug(mw.logger).Log(arguments...)

	return res, err
}
