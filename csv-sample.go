package csvsample

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
)

import (
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// CreateReadCSV is reading csv text.
func CreateReadCSV(finn io.Reader) *csv.Reader {
	// reader := csv.NewReader(finn) // utf-8
	reader := csv.NewReader(transform.NewReader(finn, japanese.ShiftJIS.NewDecoder()))
	// reader := csv.NewReader(transform.NewReader(finn, japanese.EUCJP.NewDecoder()))
	reader.LazyQuotes = true
	return reader
}

// CreateWriteCSV is output csv text.
func CreateWriteCSV(fout io.Writer) *csv.Writer {
	// writer := csv.NewWriter(fout) // utf-8
	writer := csv.NewWriter(transform.NewWriter(fout, japanese.ShiftJIS.NewDecoder()))
	// writer := csv.NewWriter(transform.NewWriter(fout, japanese.EUCJP.NewDecoder()))
	writer.UseCRLF = true
	// writer.Comma = '\t'
	return writer
}

// SrcToOut is the function for command-line.
func SrcToOut(reader *csv.Reader, writer *csv.Writer) {
	log.Printf("Start")
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("error:", err)
		}

		var newRecord []string
		for i, v := range line {
			if i > 0 {
				newRecord = append(newRecord, fmt.Sprint(i)+":"+v)
			}
			writer.Write(newRecord)
			log.Printf("%#v", line[0]+","+line[1])
		}
	}
	writer.Flush()
	log.Printf("Finish !")
}
