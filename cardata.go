package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Car map[string]any

func (c Car) String() string {
	var b strings.Builder
	for k, v := range c {
		fmt.Fprintf(&b, "%s:%T ", k, v)
	}
	return b.String()
}

type Entry struct {
	Utc  time.Time
	Cars map[int]any
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
