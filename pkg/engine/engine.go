package engine

import (
	"github.com/opype/opype/pkg/cilog"
	"github.com/opype/opype/pkg/logstore"
)

//Engine represents an OPype engine.
//This is the core of OPype and this is where the logs analysis
//and resolution suggestion will happen
type Engine struct {
	LogStore logstore.LogStore
}

//NewEngine initializes and returns a new Engine
//using the provided log store
func NewEngine(logStore logstore.LogStore) *Engine {
	return &Engine{LogStore: logStore}
}

//AnalyseLog should be called when a log is received and a
//resolution is searched for that log
func (engine *Engine) AnalyseLog(log cilog.Log) error {
	_, err := engine.LogStore.Save(log)
	return err
}
