package main

import (
	"encoding/json"
	"time"
)

// From https://github.com/theOehrly/Fast-F1.git api.py
// Status (str): 'OnTrack' or 'OffTrack'
// X, Y, Z (int): Position coordinates; starting from 2020 the coordinates are given in 1/10 meter

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
