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
	//
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

	//	for z := 2; z > 1; z++ {
	//
	//	if z == 100 {
	//		break
	//	}
	//	fmt.Println(z)
	//}

	for x, y := uint64(1), uint64(2); func() bool { return x < 100 && y < 184467440737095516 }(); func() { x++; y *= x }() {
		if x%10 == 0 {
			fmt.Println("%10", x)
			continue
		}
		fmt.Println(x, y)
	}
}
