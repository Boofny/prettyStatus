package main

import (
	"fmt"
	"time"
	"unsafe"
)

var slime = `
    █████    
  ███▒▒▒███  
 ███   ▒▒███ 
▒███    ▒███ 
▒███    ▒███ 
▒▒███   ███  
 ▒▒▒█████▒   
   ▒▒▒▒▒▒    
`
func main() {
	var s struct{}
	var one byte
	var two byte
	one = 1
	two = 1
	three := one + two
	fmt.Println(unsafe.Sizeof(three))
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(three)
	fmt.Println(unsafe.Sizeof('0'))

	// fmt.Println(day, weekDay, month)
	f := time.Now().Month()
	fmt.Println(f)
	now := time.Now()

	// Format the date as "YYYY-MM-DD"
	// The layout "2006-01-02" uses the reference time's year, month, and day
	formattedDate := now.Format("2006-01-02")
	fmt.Println("Formatted Date:", formattedDate)

	// Format the date and time as "MM/DD/YYYY HH:MM AM/PM"
	// The layout "01/02/2006 03:04PM" uses the numeric month (01), day (02),
	// year (2006), 12-hour hour (03), minute (04), and AM/PM (PM)
	formattedDateTime := now.Format("01/02/2006 03:04PM")
	fmt.Println("Formatted Date/Time:", formattedDateTime)

	dayName := time.Now().Local().Format("Mon")
	currentTime := time.Now()
	dayNumber := currentTime.Local().Day()

	fmt.Printf("Day: %s, Day Num: %d\n", dayName, dayNumber)
}




