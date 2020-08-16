package main

import (
	"fmt"

	"github.com/Ajfero/GoPkgModul.git/greet"
)

func main() {
	regards()
	greet.English()
	greet.Italian()
	greet.Spanish()
}

func regards() {
	title := "Diferents languajes regards: "
	fmt.Println(title)
}
