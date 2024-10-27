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

