package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Seen Channels
//  0  RPM        0<->13244
//  2  Speed km/h 0<->340
//  3  nGear      0<->8
//  4  Throttle   0<->100 & 104
//  45 DRS        0-14 (Odd DRS is Disabled, Even DRS is Enabled?)
//                0 =  Off
//                1 =  Off
//                2 =  (?)
//                3 =  (?)
//                8 =  Detected, Eligible once in Activation Zone (Noted Sometimes)
//                9 = ???
//                10 = On (Unknown Distinction)
//                11 = ???
//                12 = On (Unknown Distinction)
//                13 = ???
//                14 = On (Unknown Distinction)
//  5  Brake      0 100 104

type Car struct {
	Channels map[int]int
}

func (c Car) String() string {
	var b strings.Builder
	for k, v := range c.Channels {
		fmt.Fprintf(&b, "%d:%#v ", k, v)
	}
	return b.String()
}

type Entry struct {
	Utc  time.Time
	Cars map[int]Car
}

func (e Entry) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "Utc=%s ", e.Utc)
	for k, v := range e.Cars {
		fmt.Fprintf(&b, "%d:%T ", k, v)
	}
	return b.String()
}

type CarData struct {
	Entries []Entry
}

func (c CarData) String() string {
	var b strings.Builder
	b.WriteString("[")
	for i, entry := range c.Entries {
		fmt.Fprintf(&b, "%d:%s ,", i, entry)
	}
	b.WriteString("]")
	return b.String()
}

func ParseCarData(message string) (c CarData, err error) {
	err = json.Unmarshal([]byte(message), &c)
	return
}
