package watermill_test

import (
	"github.com/goph/logur"

	watermillintegration "logur.dev/integration/watermill"
)

func ExampleNew() {
	logger := watermillintegration.New(logur.NewNoopLogger())

	// Output:
	_ = logger
}
