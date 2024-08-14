package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func activity1() {
	current := time.Now()

	day := strconv.Itoa(int(current.Day()))
	month := strconv.Itoa(int(current.Month()))
	year := strconv.Itoa(int(current.Year()))
	hour := strconv.Itoa(int(current.Hour()))
	minute := strconv.Itoa(int(current.Minute()))
	second := strconv.Itoa(int(current.Second()))

	fmt.Println(hour + ":" + minute + ":" + second + " " + year + "/" + month + "/" + day)
}

func activity2() {
	date := time.Date(2023, 1, 31, 2, 49, 21, 324359102, time.UTC)
	fmt.Println(date)
}

func activity3() {
	start := time.Now()
	duration := 2 * time.Second
	time.Sleep(duration)
	end := time.Now()
	diff := end.Sub(start)
	fmt.Println("The execution took exactly", diff.Seconds(), "seconds!")
}

func activity4() {
	current := time.Now()
	fmt.Println("The current time:", current.Format(time.ANSIC))
	hours := 6 * 60 * 60
	minutes := 6 * 60
	seconds := 6
	total := hours + minutes + seconds
	future := current.Add(time.Duration(total) * time.Second)
	fmt.Println("6 hours, 6 minutes and 6 seconds from now the time will be:", future.Format(time.ANSIC))
}

func activity5() {
	currentTime := time.Now()
	eastCoast, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Println(err)
	}
	westCoast, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		log.Println(err)
	}
	eastTime := currentTime.In(eastCoast)
	westTime := currentTime.In(westCoast)
	fmt.Println("The current time is:", currentTime.Format(time.ANSIC))
	fmt.Println("The timezone", eastCoast, "current time is:", eastTime.Format(time.ANSIC))
	fmt.Println("The timezone", westCoast, "current time is:", westTime.Format(time.ANSIC))
}

func main() {
	activity1()
	activity2()
	activity3()
	activity4()
	activity5()
}
