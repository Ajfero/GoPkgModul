package greet

import "fmt"

// Para crear variables globales del paquete se debe realizar con la palabra reservada VAR, no funcion con :=
var Emoji = "😘" // variable no exportadas

// greet.GreetEnglish() //* Redundancia
func English() string {
	hello := "Hi"
	fmt.Println("The word 'Hello' in English is: ", hello)
	return "Ok" + Emoji

}

func Spanish() string {
	hello := "Hola"
	fmt.Println("The word 'Hello' in Spanish is: ", hello)
	return "Ok" + Emoji

}

func Italian() string {
	hello := "Chiao"
	fmt.Println("The word 'Hello' in Spanish is: ", hello)
	return "Ok" + Emoji
}
