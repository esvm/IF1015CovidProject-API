package covid_report

import (
	"time"

	"github.com/labstack/echo"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func instrumentingMiddleware(endpoint string, newrelicApp *newrelic.Application) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			begin := time.Now()

			err := next(c)

			CovidReportAPIRequestsTotal.With("endpoint", endpoint).Add(1)
			CovidReportAPIRequestsDuration.With("endpoint", endpoint).Observe(time.Since(begin).Seconds())
			txn := newrelicApp.StartTransaction(endpoint)
			defer txn.End()

			return err
		}
	}
}
