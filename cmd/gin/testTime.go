package main

import (
	"fmt"
	"reflect"
	"time"
)

func aboutTime() {

	//Time.time to string
	currentTime := time.Now()
	currentTimeFormat := currentTime.Format(time.RFC3339) // with predef layout
	fmt.Println(currentTimeFormat)

	//Custom Format
	fmt.Println("Custom Format:", currentTime.Format("Monday, January 2, 2006 15:04:05"))

	//Parse str to time.Time
	dateStr := "2024-02-13 12:34:56"
	parsedTime, err := time.Parse("2006-01-02 15:04:05", dateStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	fmt.Println("Parsed Time:", parsedTime, " ", reflect.TypeOf(parsedTime))

}
