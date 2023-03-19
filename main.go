package main

import (
	"flag"
	"log"
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
