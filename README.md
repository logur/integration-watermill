# Logur integration for [Watermill](https://github.com/ThreeDotsLabs/watermill)

[![CircleCI](https://circleci.com/gh/logur/integration-watermill.svg?style=svg)](https://circleci.com/gh/logur/integration-watermill)
[![Coverage](https://gocover.io/_badge/logur.dev/integration/watermill)](https://gocover.io/logur.dev/integration/watermill)
[![Go Report Card](https://goreportcard.com/badge/logur.dev/integration/watermill?style=flat-square)](https://goreportcard.com/report/logur.dev/integration/watermill)
[![GolangCI](https://golangci.com/badges/github.com/logur/integration-watermill.svg)](https://golangci.com/r/github.com/logur/integration-watermill)
[![Go Version](https://img.shields.io/badge/go%20version-%3E=1.11-61CFDD.svg?style=flat-square)](https://github.com/logur/integration-watermill)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/logur.dev/integration/watermill)


## Installation

```bash
go get logur.dev/integration/watermill
```


## Usage

```go
package main

import (
	"github.com/goph/logur"
	watermillintegration "logur.dev/integration/watermill"
)

func main() {
	logger := watermillintegration.New(logur.NewNoopLogger())
}
```


## Development

When all coding and testing is done, please run the test suite:

``` bash
$ make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
