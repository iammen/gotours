package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	firstDay := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, 0).Add(-time.Nanosecond)

	fmt.Printf("First day = %v\n", firstDay)
	fmt.Printf("Last day = %v\n", lastDay)
	fmt.Printf("Today = %v\n", t)

	t2, err := time.Parse("2006-01-02", "2019-03-01")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Time parse with format = %v", t2.Format("2006-01-02"))
}
