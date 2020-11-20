package covid_report_service

import (
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

var (
	InsertCovidReportTotal = kitprometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
		Name: "ge_covid_report_service_insert_covid_report_total",
		Help: "Insert Covid Report requests count",
	}, []string{})

	InsertCovidReportDuration = kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Name:       "ge_covid_report_service_insert_covid_report_duration_seconds",
		Help:       "Insert Covid Report duration in seconds",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01},
	}, []string{})

	GetCovidReportsTotal = kitprometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
		Name: "ge_covid_report_service_get_covid_reports_total",
		Help: "Get Covid Reports requests count",
	}, []string{})

	GetCovidReportsDuration = kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Name:       "ge_covid_report_service_get_covid_reports_duration_seconds",
		Help:       "Get Covid Reports duration in seconds",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01},
	}, []string{})
)
