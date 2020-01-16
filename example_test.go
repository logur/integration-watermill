package watermill_test

import (
	"logur.dev/logur"

	watermillintegration "logur.dev/integration/watermill"
)

func ExampleNew() {
	logger := watermillintegration.New(logur.NoopLogger{})

	// Output:
	_ = logger
}
