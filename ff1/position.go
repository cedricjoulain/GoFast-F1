package ff1

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type PositionStatus int

const (
	StatusUnknown  PositionStatus = iota
	StatusOnTrack                 // 1 "OnTrack"
	StatusOffTrack                // 2 "OffTrack"

	OnTrack  = "OnTrack"
	OffTrack = "OffTrack"
)

func (s PositionStatus) String() string {
	switch s {
	case StatusOnTrack:
		return OnTrack
	case StatusOffTrack:
		return OffTrack
	default:
		return "unknown"
	}
}

func (s *PositionStatus) MarshalJSON() ([]byte, error) {
	// TODO error if unknow ?
	return []byte(s.String()), nil
}

func (s *PositionStatus) UnmarshalJSON(data []byte) (err error) {
	switch string(data) {
	case OnTrack:
		*s = StatusOnTrack
	case OffTrack:
		*s = StatusOffTrack
	default:
		err = fmt.Errorf("unknown position status %s", data)
	}
	return
}

// From https://github.com/theOehrly/Fast-F1.git api.py
// Status (str): 'OnTrack' or 'OffTrack'
// X, Y, Z (int): Position coordinates; starting from 2020 the coordinates are given in 1/10 meter

type OnePosition struct {
	Status PositionStatus // TODO enum: seen OnTrack only
	X      int
	Y      int
	Z      int
}

func (p OnePosition) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "Status:%s ", p.Status)
	fmt.Fprintf(&b, "X:%d ", p.X)
	fmt.Fprintf(&b, "Y:%d ", p.Y)
	fmt.Fprintf(&b, "Z:%d ", p.Z)
	return b.String()
}

type PositionInfo struct {
	Entries   map[int]OnePosition
	Timestamp time.Time
}
type Position struct {
	Position []PositionInfo
}

func ParsePosition(message string) (p Position, err error) {
	err = json.Unmarshal([]byte(message), &p)
	return
}
