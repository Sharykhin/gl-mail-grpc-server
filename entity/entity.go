package entity

import (
	"encoding/json"
	"time"
)

type FailMail struct {
	ID        int64
	Action    string
	Payload   Payload
	Reason    string
	CreatedAt JSONTime
	DeletedAt *JSONTime
}

// JSONTime represents time format that should be returned to a client
type JSONTime time.Time

func (t JSONTime) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

// Payload is a specific time for json struct of a payload
type Payload json.RawMessage

// MarshalJSON implements general interface for providing data in preferable way
func (p Payload) MarshalJSON() ([]byte, error) {
	return []byte(string(p)), nil
}
