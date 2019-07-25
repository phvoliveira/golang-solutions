package main

import (
	"fmt"
)

func main() {
	specialYear := "26.09.1918"
	normalYear := "13.09."
	leapYear := "12.09."
	var year int
	fmt.Scanf("%d\n", &year)

	if year == 1918 {
		fmt.Printf("%s\n", specialYear)
	} else if (year < 1918 && year%4 == 0) || (year > 1918 && year%4 == 0 && year%100 != 0) || (year > 1918 && year%400 == 0) {
		fmt.Printf("%s%d\n", leapYear, year)
	} else {
		fmt.Printf("%s%d\n", normalYear, year)
	}
}
