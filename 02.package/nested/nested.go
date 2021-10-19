package main

import (
	"fmt"
	"github.com/asaikumaru/GoLangBasic/02.package/nested/hello"
	"github.com/asaikumaru/GoLangBasic/02.package/nested/say"
)

func main() {
	fmt.Println("Start program...")
	fmt.Printf(say.CallFromSay() + hello.CallFromHello())
}
