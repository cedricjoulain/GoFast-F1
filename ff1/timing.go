package ff1

import (
	"fmt"
	"strconv"
)

const (
	Lines = "Lines"

	Sectors                 = "Sectors"
	GapToLeader             = "GapToLeader"
	IntervalToPositionAhead = "IntervalToPositionAhead"
	Speeds                  = "Speeds"
)

type Lap struct {
}

type Timing struct {
	Car map[int]Lap // index is pilot number
}

func parseLap(message map[string]any) (l Lap, err error) {
	for k, v := range message {
		switch k {
		case Sectors, GapToLeader, IntervalToPositionAhead, Speeds:
		default:
			err = fmt.Errorf("unhandle lap info %s :%#v", k, v)
			return
		}

	}
	return
}

func ParseTimingData(message map[string]any) (t Timing, err error) {
	if len(message) != 1 {
		err = fmt.Errorf("unknown timing data message of size %d", len(message))
		return
	}
	value, ok := message[Lines]
	if !ok {
		err = fmt.Errorf("no %s key", Lines)
		return
	}
	message, ok = value.(map[string]any)
	if !ok {
		err = fmt.Errorf("wrong type for %s:%T", Lines, value)
		return
	}
	// this is car number with lap info
	t.Car = make(map[int]Lap)
	var car int
	for k, value := range message {
		message, ok = value.(map[string]any)
		if !ok {
			err = fmt.Errorf("wrong type for line %s:%T", k, value)
			return
		}
		// there should be only one
		if car, err = strconv.Atoi(k); err != nil {
			return
		}
		if t.Car[car], err = parseLap(message); err != nil {
			return
		}
	}
	return
}
