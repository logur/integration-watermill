# Logur integration for TEMPLATE

[![CircleCI](https://circleci.com/gh/logur/integration-template.svg?style=svg)](https://circleci.com/gh/logur/integration-template)
[![Coverage](https://gocover.io/_badge/logur.dev/integration/template)](https://gocover.io/logur.dev/integration/template)
[![Go Report Card](https://goreportcard.com/badge/logur.dev/integration/template?style=flat-square)](https://goreportcard.com/report/logur.dev/integration/template)
[![GolangCI](https://golangci.com/badges/github.com/logur/integration-template.svg)](https://golangci.com/r/github.com/logur/integration-template)
[![Go Version](https://img.shields.io/badge/go%20version-%3E=1.11-61CFDD.svg?style=flat-square)](https://github.com/logur/integration-template)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/logur.dev/integration/template)


## Installation

```bash
go get logur.dev/integration/template
```


## Usage

```go
package main

import (
	"github.com/goph/logur"
	templateintegration "logur.dev/integration/template"
)

func main() {
	logger := templateintegration.New(logur.NewNoopLogger())
}
```


## Development

When all coding and testing is done, please run the test suite:

``` bash
$ make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
