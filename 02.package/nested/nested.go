package main

import (
	"fmt"
	"github.com/GoLangBasic/02.package/nested/hello"
	"github.com/GoLangBasic/02.package/nested/say"
)

func main() {
	fmt.Println("Start program...")
	fmt.Printf(say.CallFromSay() + hello.CallFromHello())
}
