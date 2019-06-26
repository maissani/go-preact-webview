package main

import "fmt"

var tata string

func main() {
	var tata string = "tata"
	fmt.Println("toto", tata, &tata)
	chagepourtiti(&tata)
	fmt.Println("toto", tata, &tata)
	tata = "tutu"
	fmt.Println("toto", tata, &tata)
}

func chagepourtiti(t *string) {
	*t = "titi"
}
