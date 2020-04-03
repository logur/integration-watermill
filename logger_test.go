package watermill

import (
	"errors"
	"testing"

	"github.com/ThreeDotsLabs/watermill"
	"logur.dev/logur"
	"logur.dev/logur/logtesting"
)

func TestLogger(t *testing.T) {
	var logger interface{} = New(nil)

	if _, ok := logger.(watermill.LoggerAdapter); !ok {
		t.Error("Logger does not implement the watermill.LoggerAdapter interface")
	}
}

func TestLogger_Error_NilFields(t *testing.T) {
	testLogger := &logur.TestLogger{}
	logger := New(testLogger)

	err := errors.New("error")
	logger.Error("something went wrong", err, nil)

	lastEvent := testLogger.LastEvent()

	if lastEvent == nil {
		t.Fatal("failed to get last event")
	}

	logEvent := logur.LogEvent{
		Line:   "something went wrong",
		Level:  logur.Error,
		Fields: map[string]interface{}{"err": err},
	}

	logtesting.AssertLogEventsEqual(t, logEvent, *lastEvent)
}
