package main

import (
	"fmt"
	"time"
)

func switch_examples() {

	var now = time.Now()

	switch now.Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Today Is Weekend, Just Enjoy")
	default:
		fmt.Println("Please Complete The Work")
	}

	switch {
	case now.Day() > 2 :
		fmt.Println("Last 28 days are really difficult, with the remaining Salary")
	default:
		fmt.Println("Please Spend Less (-_-)")
	}
	
}
