package state

import "time"

type Poll struct {
	Title     string
	Options   []string
	Votes     map[string]int
	EndTime   time.Time
	MessageID string
}
