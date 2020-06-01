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
	exactTime, err := getExactTime()
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}

	fmt.Printf("current time: %s\n", localTime)
	fmt.Printf("exact time: %s\n", exactTime)
}

func getLocalTime() string {
	return time.Now().Round(0).Format(layout)
}

func getExactTime() (string, error) {
	exactTime, err := ntp.Time("1.beevik-ntp.pool.ntp.org")
	if err != nil {
		return "", err
	}

	return exactTime.Format(layout), err
}
