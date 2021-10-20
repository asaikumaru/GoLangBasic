package main

import "fmt"

var car = "ford"

func main() {

	fmt.Println(car)
	CheckSpareWheel()

}

func CheckSpareWheel() {
	if car == "ford" {
		fmt.Println("Spare wheel exist")
	}
}
