# pflag-date #
[![Go Reference](https://pkg.go.dev/badge/github.com/gbarr/pflag-date.svg)](https://pkg.go.dev/github.com/gbarr/pflag-date)
[![codecov](https://codecov.io/github/gbarr/pflag-date/graph/badge.svg?token=WCV5JUZHFY)](https://codecov.io/github/gbarr/pflag-date)
[![Go ReportCard](https://goreportcard.com/badge/gbarr/pflag-date)](http://goreportcard.com/report/gbarr/pflag-date)
[![golangci-lint](https://github.com/gbarr/pflag-date/actions/workflows/lint.yml/badge.svg)](https://github.com/gbarr/pflag-date/actions/workflows/lint.yml)

[`pflag-date`](https://github.com/gbarr/pflag-date) implements a Golang [`pflag.Value`](https://pkg.go.dev/github.com/spf13/pflag#Value) interface for `YYYY-MM-DD`-specified dates.
Combining [github.com/spf13/pflag](https://pkg.go.dev/github.com/spf13/pflag) with [github.com/hardfinhq/go-date](https://pkg.go.dev/github.com/hardfinhq/go-date).
This facilitiates command-line argument handling of date parameters such  `--start-date=2024-10-01`.

## Documentation ##

```
go get github.com/gbarr/pflag-date
```

https://pkg.go.dev/github.com/gbarr/pflag-date

## Example ##

```go
package main

import (
	"fmt"
	pfdate "github.com/gbarr/pflag-date"
	"github.com/spf13/pflag"
)

func main() {
	var dt pfdate.Date
	pflag.VarP(&dt, "start-date", "d", "YYYY-MM-DD date")
	pflag.Parse()
	fmt.Println("date:", dt.String())
}

```
----

## Credits and License

Copyright (c) 2024 [Graham Barr](https://github.com/gbarr).

Released under the [MIT License](https://en.wikipedia.org/wiki/MIT_License), see [LICENSE](./LICENSE).
