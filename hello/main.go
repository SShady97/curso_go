package main

import (
	"fmt"
	"log"

	"github.com/SShady97/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0) //Muestra fecha y hora en registros. Con 0 no los muestra, con 1 si.

	names := []string{"Jose", "Manuel", "Margarita"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	//message, err := greetings.Hello("Jose Alvarez")
	//if err != nil {
	//log.Fatal(err)
	//}
	fmt.Println(messages)
}
