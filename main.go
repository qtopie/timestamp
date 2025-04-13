package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	timeFlag := flag.String("time", "", "Time value of unix epoch seconds/milliseconds")
	durationFlag := flag.String("duration", "", "Duration of time interval")
	flag.Parse()

	if len(*durationFlag) > 0 {
		d := parseDuration(*durationFlag)
		printDuration(d)
	} else {
		// default to print current time
		t := parseTime(*timeFlag)
		printTime(t)
	}
}

func parseTime(s string) time.Time {
	if len(s) == 0 {
		return time.Now()
	}

	epoch, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	if epoch < 1<<31 {
		return time.UnixMilli(int64(epoch) * 1000)
	} else {
		return time.UnixMilli(int64(epoch))
	}
}

func printTime(t time.Time) {
	fmt.Println(t)
	fmt.Println("Time Epoch(ms):", t.UnixMilli())
}

func parseDuration(s string) time.Duration {
	if strings.Contains(s, ",") {
		values := strings.Split(s, ",")
		if len(values) < 2 {
			panic("required two time values for an interval")
		}

		t1 := parseTime((values[0]))
		t2 := parseTime(values[1])

		if t1.After(t2) {
			return t1.Sub(t2)
		}

		return t2.Sub(t1)
	} else {
		d, err := time.ParseDuration(s)
		if err != nil {
			panic(err)
		}
		return d
	}

}

func printDuration(d time.Duration) {
	// s := ""
	// if d.Abs().Hours() > 24 {
	// 	s += string(d.Abs().Hours() / 24)
	// }
	fmt.Println("Duration:", d)
}
