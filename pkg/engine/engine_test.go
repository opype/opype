package engine

import (
	"testing"
	"time"

	"github.com/opype/opype/pkg/cilog"
	"github.com/opype/opype/pkg/logstore"
)

func TestAnalyseLog(t *testing.T) {
	tn := time.Now()
	logToAnalyse := cilog.Log{
		Company: "Company Name or ID",
		Date:    &tn,
		Lines:   []string{"some", "log", "files"},
	}
	store := logstore.NewMemoryStore()
	eng := NewEngine(store)
	err := eng.AnalyseLog(logToAnalyse)
	if err != nil {
		t.Errorf("No error expected, but got: %s", err)
	}

	logsAfter, err := store.FindByCompany(logToAnalyse.Company)
	if err != nil {
		t.Errorf("No error expected when searching for logs, but got: %s", err)
	}

	if len(logsAfter) != 1 {
		t.Errorf("Expected to find just this one log in the store after analyse, but got %d logs", len(logsAfter))
	}
}
