package cilog

import (
	"time"
)

//Log is the basic log we can receive from a CI / CD system
type Log struct {
	ID      string
	Lines   []string
	Company string
	Date    *time.Time
}
