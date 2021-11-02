package main

import "fmt"

type Edocument struct {
	Number        string
	Date          string
	NumberOfPages int
	Footer
}

type Footer struct {
	Author string
}

func main() {

	doc1 := Edocument{

		Number:        "001-A",
		Date:          "02.11.2021",
		NumberOfPages: 10,
		Footer: Footer{
			Author: "Author-1",
		},
	}

	var doc2 Edocument
	doc2.Number = "002-A"
	doc2.Date = "02.11.2021"
	doc2.NumberOfPages = 5
	doc2.Author = "Author-2" // == doc2.Footer.Author= "Author-2"

	fmt.Printf("%T-%v\n", doc1, doc1)
	fmt.Printf("%T-%v\n", doc2, doc2)
}
