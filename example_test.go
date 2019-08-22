package template_test

import (
	"github.com/goph/logur"

	templateintegration "logur.dev/integration/template"
)

func ExampleNew() {
	logger := templateintegration.New(logur.NewNoopLogger())

	// Output:
	_ = logger
}
