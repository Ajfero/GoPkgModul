package main

import "fmt"

func main() {
	regards()
	fmt.Println(greet.English)
	fmt.Println(greet.Italian)
	fmt.Println(greet.Spanish)
}

func regards() {
	title := "Diferents languajes regards"
	fmt.Println("Hello in English is. ", title)
}
