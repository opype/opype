package logstore

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/opype/opype/pkg/cilog"
)

func TestMemoryStoreSave(t *testing.T) {
	ms := NewMemoryStore()
	tn := time.Now()
	logToSave := cilog.Log{
		Company: uuid.New().String(),
		Lines:   []string{"just", "some", "lines"},
		Date:    &tn,
	}

	t.Run("Save", func(t *testing.T) {
		logSaved, err := ms.Save(logToSave)
		if err != nil {
			t.Errorf("Expected no error, but got: %s", err)
		}
		if logSaved.ID == "" {
			t.Error("Expected a log with an UUID, but got no UUID")
		}
		logToSave.ID = logSaved.ID

		if !reflect.DeepEqual(logToSave, *logSaved) {
			t.Errorf("Expected %v, but got %v", logToSave, *logSaved)
		}
	})

	t.Run("FindByID", func(t *testing.T) {
		t.Run("Existent", func(t *testing.T) {
			logFound, err := ms.FindByID(logToSave.ID)
			if err != nil {
				t.Errorf("Expected no error, but got: %s", err)
			}
			if logFound == nil {
				t.Error("Expected a log to be found, but got: nil")
			}
			if !reflect.DeepEqual(logToSave, *logFound) {
				t.Errorf("Expected %v, but got %v", logToSave, *logFound)
			}
		})
		t.Run("Not existent", func(t *testing.T) {
			logFound, err := ms.FindByID("NOT-EXISTENT-ID")
			if err != nil {
				t.Errorf("Expected no error, but got: %s", err)
			}
			if logFound != nil {
				t.Errorf("Expected no log to be found, but got: %v", *logFound)
			}
		})
	})

	t.Run("FindByCompany", func(t *testing.T) {
		t.Run("Existent", func(t *testing.T) {
			logsFound, err := ms.FindByCompany(logToSave.Company)
			if err != nil {
				t.Errorf("Expected no error, but got: %s", err)
			}
			if len(logsFound) != 1 {
				t.Errorf("Expected 1 log to be found, but got: %d", len(logsFound))
			}
			if !reflect.DeepEqual(logToSave, logsFound[0]) {
				t.Errorf("Expected %v, but got %v", logToSave, logsFound[0])
			}
		})

		t.Run("Not existent", func(t *testing.T) {
			logsFound, err := ms.FindByCompany("A not existent company")
			if err != nil {
				t.Errorf("Expected no error, but got: %s", err)
			}
			if len(logsFound) != 0 {
				t.Errorf("Expected 0 logs to be found, but got: %v", logsFound)
			}
		})
	})

	t.Run("FindByCompanyAndDates", func(t *testing.T) {
		tnStart := tn.Add(-1 * 60 * 60 * time.Second)
		tnEnd := tn.Add(1 * 60 * 60 * time.Second)
		t.Run("Existent", func(t *testing.T) {
			logsFound, err := ms.FindByCompanyAndDates(logToSave.Company, tnStart, tnEnd)
			if err != nil {
				t.Errorf("Expected no error, but got: %s", err)
			}
			if len(logsFound) != 1 {
				t.Errorf("Expected 1 log to be found, but got: %d", len(logsFound))
			}
			if !reflect.DeepEqual(logToSave, logsFound[0]) {
				t.Errorf("Expected %v, but got %v", logToSave, logsFound[0])
			}
		})

		t.Run("Not existent - good times", func(t *testing.T) {
			logsFound, err := ms.FindByCompanyAndDates("A not existent company", tnStart, tnEnd)
			if err != nil {
				t.Errorf("Expected no error, but got: %s", err)
			}
			if len(logsFound) != 0 {
				t.Errorf("Expected 0 logs to be found, but got: %v", logsFound)
			}
		})

		t.Run("Not existent - good company", func(t *testing.T) {
			logsFound, err := ms.FindByCompanyAndDates(logToSave.Company, tnEnd, tnStart)
			if err != nil {
				t.Errorf("Expected no error, but got: %s", err)
			}
			if len(logsFound) != 0 {
				t.Errorf("Expected 0 logs to be found, but got: %v", logsFound)
			}
		})
	})

	t.Run("Delete", func(t *testing.T) {
		err := ms.Delete(logToSave.ID)
		if err != nil {
			t.Errorf("Expected no error, but got: %s", err)
		}

		logFound, err := ms.FindByID(logToSave.ID)
		if err != nil {
			t.Errorf("Expected no error, but got: %s", err)
		}

		if logFound != nil {
			t.Errorf("Expected to find no log after deletion, but found: %v", logFound)
		}
	})

}
