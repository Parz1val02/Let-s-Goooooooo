package main

import (
	"fmt"
	"time"
)

func example1() {
	start := time.Now()
	fmt.Println("The script started at: " + start.Format(time.ANSIC))
	sum := 0
	for i := 1; i < 10000000000; i++ {
		sum += i
	}
	end := time.Now()
	duration := end.Sub(start)
	fmt.Println("The script ended at: " + end.Format(time.ANSIC))
	fmt.Println("The task took", duration.Hours(), "hour(s) to complete!")
	fmt.Println("The task took", duration.Minutes(), "minutes(s) to complete!")
	fmt.Println("The task took", duration.Seconds(), "seconds(s) to complete!")
	fmt.Println("The task took", duration.Milliseconds(), "millisecond(s) to complete!")
	fmt.Println("The task took", duration.Microseconds(), "microsecond(s) to complete!")
	fmt.Println("The task took", duration.Nanoseconds(), "nanosecond(s) to complete!")
}

func example2() {
	deadlineSeconds := time.Duration((600 * 10) * time.Millisecond)
	start := time.Now()
	fmt.Println("Deadline for the transaction is", deadlineSeconds)
	fmt.Println("The transaction has started at:", start)
	sum := 0
	for i := 1; i < 25000000000; i++ {
		sum += i
	}
	end := time.Now()
	duration := end.Sub(start)
	transactionTime := time.Duration(duration.Nanoseconds()) * time.Nanosecond
	fmt.Println("The transaction has completed at:", end, duration)
	if transactionTime <= deadlineSeconds {
		fmt.Println("Performance is OK transaction completed in",
			transactionTime)
	} else {
		fmt.Println("Performance problem, transaction completed in",
			transactionTime, "second(s)!")
	}
}

func example3() {
	t1, err := time.Parse(time.RFC3339, "2019-09-27T22:18:11+00:00")
	if err != nil {
		fmt.Println(err)
	}
	t2, err := time.Parse(time.UnixDate, "2019-09-27T22:18:11+00:00")
	if err != nil {
		fmt.Println(err)
	}
	t3, err := time.Parse(time.ANSIC, "2019-09-27T22:18:11+00:00")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("RFC3339:", t1)
	fmt.Println("UnixDate", t2)
	fmt.Println("ANSIC", t3)
}

func example4() {
	date := time.Date(2019, 9, 27, 18, 50, 48, 324359102, time.UTC)
	nextDate := date.AddDate(1, 2, 3)
	fmt.Println(nextDate)

	current := time.Now()
	losAngeles, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The local current time is:", current.Format(time.ANSIC))
	fmt.Println("The time in Los Angeles is:", current.In(losAngeles).Format(time.ANSIC))
}

func timeDiff(timezone string) {
	current := time.Now()
	timeZone, err := time.LoadLocation(timezone)
	if err != nil {
		fmt.Println(err)
	}
	remoteTime := current.In(timeZone)
	fmt.Println("The current time is: ", current.Format(time.ANSIC))
	fmt.Println("The timezone ", timezone, "current time is: ", remoteTime.Format(time.ANSIC))
}

func main() {
	example1()
	example2()
	example3()
	example4()
	timeDiff("America/Los_Angeles")
}
