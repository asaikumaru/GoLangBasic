package main

import "fmt"

func main() {
	fmt.Println("Start program...")
	fmt.Printf(say.CallFromSay() + hello.CallFromHello())
}