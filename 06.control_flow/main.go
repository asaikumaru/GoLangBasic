package main

import "fmt"

func main() {
	// if	// else

	//a, b, c := 10, 20, 30
	//if a>b {
	//	fmt.Println("a")
	//}else if b>c{
	//	fmt.Println("b")
	//}else {
	//	fmt.Println("c")
	//}

	//var text = "Admin"
	//
	//switch {
	//case text == "Admin" || text == "admin":
	//	fmt.Println("ok")
	//default:
	//	fmt.Println("not ok")
	//}

	//switch text {
	//case "admin":
	//	fmt.Println("Error not admin")
	//case "odmin":
	//	fmt.Println("Error not admin")
	//case "Admin":
	//	fmt.Println("Yeah - Admin!")
	//fallthrough
	//case "Some Amdin":
	//	fmt.Println("Key")
	//default:
	//	fmt.Println("Some login")
	//}

	for x := 1; x < 100; x++ {
		if x%10 == 0 {
			fmt.Println("%10", x)
			break
		}
		fmt.Println(x)
	}
}
