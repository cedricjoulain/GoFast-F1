package main

import (
	"bufio"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/ulikunitz/xz"
)

// data from
// https://livetiming.formula1.com/signalr
func main() {
	ptrName := flag.String("name", "/mnt/raid0/data/fastf1/20230105_fastf1_data.txt.gz", "livetiming.formula1.com data file")
	flag.Parse()

	if err := ParseF1File(*ptrName); err != nil {
		log.Fatal(err)
	}
}

// ParseF1File read log file and parse line by line
func ParseF1File(filename string) (err error) {
	var (
		rc *os.File
		r  io.Reader
	)
	if rc, err = os.Open(filename); err != nil {
		return
	}
	defer rc.Close()

	switch filepath.Ext(filename) {
	case ".gz":
		if rz, rzerr := gzip.NewReader(rc); err != nil {
			return rzerr
		} else {
			r = rz
			defer rz.Close()
		}
	case ".xz":
		// xz is just a reader
		if r, err = xz.NewReader(rc); err != nil {
			return
		}
	default:
		// keep same file reader
		r = rc
	}

	scanner := bufio.NewScanner(r)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return scanner.Err()
}
