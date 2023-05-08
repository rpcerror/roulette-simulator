package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var balance float64
	balance = 100 // starting balance
	var colorsstruct []string
	var reset int
	var roll int
	var together bool
	var bet bool
	var betamount float64
	for {
		var color string
		var black int
		if together == true {
			together = false
			bet = true
			if balance <= 10 {
				betamount = balance
			} else {
				betamount = balance / 10 //change the number depending on your total balance's percentage per bet. 10 = 10%, 20 = 5% etc.
			}
		}
		roll++
		fmt.Println("Roll:", roll, "Balance:", balance)
		if balance == 0 {
			break
		}
		//time.Sleep(time.Millisecond * 100)
		number := rand.Intn(15)
		if number == 0 {
			color = "green"
		}
		if number >= 1 && number <= 7 {
			color = "red"
		}
		if number >= 8 && number <= 14 {
			color = "black"
		}

		if bet {
			if color == "green" { //change to whatever color you want to bet on
				fmt.Println("W")
				balance = balance + betamount
			} else {
				fmt.Println("L")
				balance = balance - betamount
			}
			bet = false
		}
		colorsstruct = append(colorsstruct, color)
		x := 0
		for _, y := range colorsstruct {
			x++
			if y == "black" { //change the consecutive color
				black++
			}
			if x == 10 { //change the number depending on how many consecutive colors you want to bet after
				if black == 10 { //change the number depending on how many consecutive colors you want to bet after
					together = true
				}
				black = 0
				x = 0
				colorsstruct = nil
				reset++
			}
		}
	}
	time.Sleep(time.Second * 999)
}
