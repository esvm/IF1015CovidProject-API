package covid_report_service

import (
	"github.com/esvm/if1015covidproject-api/src/covid_reports"
	"github.com/esvm/if1015covidproject-api/src/validation"
)

func ValidateCovidReport(report *covid_reports.CovidReport) error {
	return validation.Validate(report)
}
