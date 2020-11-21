package covid_report_service

import (
	"github.com/esvm/if1015covidproject-api/src/covid_reports"
	"github.com/esvm/if1015covidproject-api/src/validation"
)

func ValidateCovidReports(reports []*covid_reports.CovidReport) error {
	for _, report := range reports {
		if err := validation.Validate(report); err != nil {
			return err
		}
	}
	return nil
}
