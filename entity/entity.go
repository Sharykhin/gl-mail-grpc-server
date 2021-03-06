package entity

import (
	"time"

	"encoding/json"
)

// JSONTime represents time format that should be returned to a client
type JSONTime time.Time

// String implements general interface to change represent time in string format
func (t JSONTime) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

// FailMail internal entity that with nullable DeletedAt field
type FailMail struct {
	ID        int64
	Action    string
	Payload   json.RawMessage
	Reason    string
	CreatedAt JSONTime
	DeletedAt *JSONTime
}
