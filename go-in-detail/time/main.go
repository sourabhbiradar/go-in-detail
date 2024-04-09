package main

import (
	"fmt"
	"time" // time package
)

func main() {
	// time
	fmt.Println("\ttime")

	// time.Duration
	fmt.Println("time.Duration")

	timeNow := time.Now()
	fmt.Println(timeNow) // date time ...
	
	fT:=timeNow.Add(2*time.Hour)
	fmt.Println("two hours from now :",fT)

	t := 2*time.Hour + 34*time.Minute + 12*time.Second
	fmt.Println(t) // 2h34m12s

	customFormat := timeNow.Format("2006-01-02 Firday, 3:04 PM")
	fmt.Println("Custom formatted time:", customFormat)

	// time.Date(year, month, day, hour, min, sec, nsec, loc)
	date:=time.Date(2021,time.May,11,5,23,9,55,time.Local)
	fmt.Println("Date :",date)

	fmt.Println("year only :",date.Year())

}

	

  

