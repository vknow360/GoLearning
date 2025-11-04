package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	specificTime := time.Date(2025, time.January, 30, 12, 0, 0, 0, time.UTC)
	fmt.Println(specificTime)

	parsedTime, _ := time.Parse("2006-01-02", "2020-05-01")
	fmt.Println(parsedTime)

	t := time.Now()
	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println(oneDayLater)
	fmt.Println(oneDayLater.Weekday())

	loc, _ := time.LoadLocation("Asia/Kolkata")
	t = time.Date(2025, time.July, 8, 14, 16, 40, 00, time.UTC)

	tLocal := t.In(loc)

	roundedTime := t.Round(time.Hour)
	roundedTimeLocal := roundedTime.In(loc)

	fmt.Println("Original Time (UTC):", t)
	fmt.Println("Original Time (Local):", tLocal)
	fmt.Println("Rounded Time (UTC):", roundedTime)
	fmt.Println("Rounded Time (Local):", roundedTimeLocal)

	fmt.Println(time.Now().Format("2006-01-02"))

}
