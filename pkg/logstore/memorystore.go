package logstore

import (
	"fmt"
	"sync"
	"time"

	"github.com/opype/opype/pkg/cilog"

	"github.com/google/uuid"
)

//MemoryStore implements a simple store with the logs kept in memory
type MemoryStore struct {
	LogStore
	Logs  []cilog.Log
	mutex sync.Mutex
}

//NewMemoryStore creates a MemoryStore prepared
//for further operation
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		Logs: []cilog.Log{},
	}
}

//Save saves the passed log in the memory store
func (memStore *MemoryStore) Save(log cilog.Log) (*cilog.Log, error) {
	log.ID = uuid.New().String()
	memStore.Logs = append(memStore.Logs, log)
	return &log, nil
}

//Delete removes the log with the passed id from the memory store
func (memStore *MemoryStore) Delete(id string) error {

	for i, log := range memStore.Logs {
		if log.ID == id {
			memStore.mutex.Lock()
			if log.ID == id {
				memStore.Logs = append(memStore.Logs[:i], memStore.Logs[i+1:]...)
			}
			memStore.mutex.Unlock()
			return nil
		}
	}

	return fmt.Errorf("No log with the specified ID found")
}

//FindByID will return the log with the specified ID from the memory store
func (memStore *MemoryStore) FindByID(id string) (*cilog.Log, error) {
	for _, log := range memStore.Logs {
		if log.ID == id {
			return &log, nil
		}
	}

	return nil, nil
}

//FindByCompany will return all the logs for the specified company from the memory store
func (memStore *MemoryStore) FindByCompany(company string) ([]cilog.Log, error) {
	logs := []cilog.Log{}

	for _, log := range memStore.Logs {
		if log.Company == company {
			logs = append(logs, log)
		}
	}

	return logs, nil
}

//FindByCompanyAndDates will return all the logs between the specified dates for the specified company
func (memStore *MemoryStore) FindByCompanyAndDates(company string,
	startDate, endDate time.Time) ([]cilog.Log, error) {
	logs := []cilog.Log{}

	for _, log := range memStore.Logs {
		if log.Company == company &&
			log.Date.After(startDate) &&
			log.Date.Before(endDate) {
			logs = append(logs, log)
		}

	}

	return logs, nil
}
