package ff1

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
	fmt.Fprintf(&b, "RPM:%d ", c.Channels[0])
	fmt.Fprintf(&b, "Speed:%d ", c.Channels[2])
	fmt.Fprintf(&b, "nGear:%d ", c.Channels[3])
	fmt.Fprintf(&b, "Throttle:%d ", c.Channels[4])
	fmt.Fprintf(&b, "DRS:%d ", c.Channels[45])
	fmt.Fprintf(&b, "Brake:%d ", c.Channels[5])
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
