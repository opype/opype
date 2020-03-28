package logstore

import (
	"time"

	"github.com/opype/opype/pkg/cilog"
)

//LogStore defines the methods needed for a log store
//so that it can be used with OPype engine
type LogStore interface {
	Save(cilog.Log) (*cilog.Log, error)
	Delete(id string) error
	FindById(id string) (*cilog.Log, error)
	FindByCompany(company string) ([]cilog.Log, error)
	FindByCompanyAndDates(company string, startDate time.Time, endDate time.Time) ([]cilog.Log, error)
}
