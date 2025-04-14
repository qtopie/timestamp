package main

import (
	"flag"
	"fmt"
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
		loc, _ := time.LoadLocation("Asia/Shanghai")
		t, err := time.ParseInLocation(time.DateOnly, s, loc)
		if err != nil {
			panic(err)
		}

		return t
	}

	// 3000 AD (公元3000年)
	if epoch < 32503651201 {
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

		t1 := parseTime(strings.TrimSpace(values[0]))
		t2 := parseTime(strings.TrimSpace(values[1]))

		fmt.Println("First Moment:", t1)
		fmt.Println("Second Moment:", t2)

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
	fmt.Println("Total Time:", d)

	fmt.Print("Human-Readable Time:")
	ms := d.Milliseconds()
	if ms > 10*365*24*3600*1000 {
		years := ms / (365 * 24 * 3600 * 1000)
		fmt.Print(years, "Y")
		ms -= years * (365 * 24 * 3600 * 1000)
	}

	if ms > 24*3600*1000 {
		days := ms / (24 * 3600 * 1000)
		fmt.Print(days, "D")
		ms -= days * (24 * 3600 * 1000)
	}
	d, _ = time.ParseDuration(strconv.Itoa(int(ms)) + "ms")
	fmt.Println(d)
}
