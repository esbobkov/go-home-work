package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

const layout = "2006-01-02 03:04:05 -0700 MST"

func main() {
	localTime := getLocalTime()
	exactTime := getExactTime()

	fmt.Printf("current time: %s\n", localTime)
	fmt.Printf("exact time: %s\n", exactTime)
}

func getLocalTime() string {
	return time.Now().Round(0).Format(layout)
}

func getExactTime() string {
	exactTime, err := ntp.Time("1.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}

	return exactTime.Format(layout)
}
