package main

import (
	// "encoding/csv"
	"flag"
	// "fmt"
	// "io"
	"log"
	"os"
)

import (
// "golang.org/x/text/encoding/japanese"
// "golang.org/x/text/transform"
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

	csvsample.Src_to_out(
		csvsample.CreateReadCSV(finn),
		csvsample.CreateWriteCSV(fout))
}
