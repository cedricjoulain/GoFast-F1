package main

import (
	"encoding/json"
	"time"
)

type OnePosition struct {
	Status string // TODO enum: seen OnTrack only
	X      int
	Y      int
	Z      int
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
