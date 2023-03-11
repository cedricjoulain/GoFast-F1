package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

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

// BaseMessage common description
type BaseMessage []any

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
		message := make(BaseMessage, 0)
		line := CorrectedLine(scanner.Text())
		if err = json.Unmarshal([]byte(line), &message); err != nil {
			log.Println("error", err, "parsing line", line)
			continue
		}
		if err = ParseMessage(message); err != nil {
			log.Println("error", err, "parsing line", line)
			continue
		}
	}

	return scanner.Err()
}

// CorrectedLine, make it json unmashable
func CorrectedLine(line string) string {
	// wrong string escape
	line = strings.ReplaceAll(line, "'", `"`)
	// wrong boolean
	line = strings.ReplaceAll(line, "True", "true")
	line = strings.ReplaceAll(line, "False", "false")
	return line
}

// ParseMessage parse "digested" F1 line
func ParseMessage(message BaseMessage) (err error) {
	if len(message) < 3 {
		return fmt.Errorf("message too short (len = %d)", len(message))
	}
	var (
		value string
		topic string
		tms   time.Time
	)
	if topic, err = AsString(message[0]); err != nil {
		return
	}
	if value, err = AsString(message[len(message)-1]); err != nil {
		return
	}
	if tms, err = time.ParseInLocation(time.RFC3339Nano, value, time.UTC); err != nil {
		return
	}

	switch topic {
	default:
		log.Println(tms, topic)
	}
	return
}

// AsString ensure it's a strings
func AsString(in any) (s string, err error) {
	switch e := in.(type) {
	case string:
		s = e
	default:
		err = fmt.Errorf("%#v is a %T not a string", e, in)
	}
	return
}
