package main

import (
	"flag"
	"log"

	"fastf1/ff1"
)

// data from
// https://livetiming.formula1.com/signalr
func main() {
	ptrName := flag.String("name", "/mnt/raid0/data/fastf1/20230105_fastf1_data.txt.gz", "livetiming.formula1.com data file")
	ptrDebug := flag.Bool("debug", false, "verbose debug inforamtions")
	flag.Parse()

	if err := ff1.ParseF1File(*ptrName, *ptrDebug); err != nil {
		log.Fatal(err)
	}
}
