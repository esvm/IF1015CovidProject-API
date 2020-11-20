package covid_report

import (
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

var (
	CovidReportAPIRequestsDuration = kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Name:       "ge_api_gateway_covid_report_api_requests_duration_seconds",
		Help:       "Covid Report api requests duration in seconds",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01},
	}, []string{"endpoint"})

	CovidReportAPIRequestsTotal = kitprometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
		Name: "ge_api_gateway_covid_report_api_requests_total",
		Help: "Covid Report api requests count",
	}, []string{"endpoint"})
)
