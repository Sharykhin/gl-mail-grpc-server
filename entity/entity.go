package entity

import (
	"time"

	"github.com/Sharykhin/gl-mail-grpc"
)

// JSONTime represents time format that should be returned to a client
type JSONTime time.Time

// String implements general interface to change represent time in string format
func (t JSONTime) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

// FailMail internal entity that with nullable DeletedAt field
type FailMail struct {
	api.FailMailEntity
	DeletedAt *JSONTime
}
