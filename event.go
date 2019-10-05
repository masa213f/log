package log

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	// ISO8601 + nano second
	TimestampFormat = "2006-01-02T15:04:05.999999"
)

type eventTime time.Time

func (t *eventTime) String() string {
	return time.Time(*t).Format(TimestampFormat)
}

func (t *eventTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

type event struct {
	Timestamp eventTime `json:"ts"`
	Level     string    `json:"lv"`
	Message   string    `json:"msg,omitempty"`
}

func newEvent(t time.Time, lv Level, msg string) *event {
	return &event{
		Level:     lv.String(),
		Timestamp: eventTime(t),
		Message:   msg,
	}
}

func (e *event) json() string {
	byte, _ := json.Marshal(e)
	return string(byte)
}

func (e *event) text() string {
	return fmt.Sprintf("%s: [%s] %s", e.Timestamp.String(), e.Level, e.Message)
}
