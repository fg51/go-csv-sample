package main

import (
	"flag"
	"log"
	"os"
)

import "github.com/kflange/go-csv-sample"

func failOnError(err error) {
	log.Fatal("error:", err)
}

func main() {
	flag.Parse()

	finn, iErr := os.Open(flag.Arg(0))
	if iErr != nil {
		failOnError(iErr)
	}
	defer finn.Close()

	fout, oErr := os.Create(flag.Arg(1))
	if oErr != nil {
		failOnError(oErr)
	}
	defer fout.Close()

	csvsample.SrcToOut(
		csvsample.CreateReadCSV(finn),
		csvsample.CreateWriteCSV(fout),
	)
}
