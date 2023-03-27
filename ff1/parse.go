package ff1

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ulikunitz/xz"
)

// BaseMessage common description
type BaseMessage []any

// ParseF1File read log file and parse line by line
func ParseF1File(filename string, debug bool) (err error) {
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
		if _, err = ParseMessage(message, debug); err != nil {
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
func ParseMessage(message BaseMessage, debug bool) (parsed any, err error) {
	if len(message) < 3 {
		err = fmt.Errorf("message too short (len = %d)", len(message))
		return
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
	if tms.IsZero() {
		// nothing ?
		return
	}
	switch topic {
	case TopicDriverList, TopicExtrapolatedClock, TopicHeartbeat, TopicLapCount, TopicRaceControlMessages, TopicSessionData, TopicTimingAppData, TopicTimingStats, TopicTopThree, TopicTrackStatus, TopicWeatherData:
		// TODO
	case TopicTimingData:
		if m, ok := (message[1]).(map[string]any); !ok {
			err = fmt.Errorf("timing data should be map[string]any is a %T", message[1])
			return
		} else {
			if parsed, err = ParseTimingData(m); err != nil {
				return
			}
		}
	case TopicCarData, TopicPosition:
		if value, err = AsString(message[1]); err != nil {
			return
		}
		if value, err = DecodeZInfo(value); err != nil {
			return
		}
		switch topic {
		case TopicCarData:
			if parsed, err = ParseCarData(value); err != nil {
				return
			}
			if debug {
				for _, entry := range parsed.(CarData).Entries {
					for i, car := range entry.Cars {
						fmt.Printf("%s car:%d %s\n", entry.Utc, i, car)
					}
				}
			}
		case TopicPosition:
			if parsed, err = ParsePosition(value); err != nil {
				return
			}
			if debug {
				for _, pinfo := range parsed.(Position).Position {
					for i, onepos := range pinfo.Entries {
						fmt.Printf("%s car:%d %s\n", pinfo.Timestamp, i, onepos)
					}
				}
			}
		}
	default:
		err = fmt.Errorf("unknown topic %s", topic)
	}
	return
}

// DecodeZInfo decode, decompress, return as string
func DecodeZInfo(coded string) (decoded string, err error) {
	var (
		in, out []byte
		r       io.ReadCloser
	)
	// Decode base64-encoded input
	if in, err = base64.StdEncoding.DecodeString(coded); err != nil {
		return
	}
	// Decompress using zlib
	buffer := bytes.NewReader(in)
	if r, err = NewReader(buffer); err != nil {
		return
	}
	defer r.Close()
	if out, err = io.ReadAll(r); err != nil {
		return
	}
	decoded = string(out)
	return
}

// AsString ensure it's a string
func AsString(in any) (s string, err error) {
	switch e := in.(type) {
	case string:
		s = e
	default:
		err = fmt.Errorf("%#v is a %T not a string", e, in)
	}
	return
}
