package main

import (
	"library/animal"
)

func main() {
	//var myBook = book.NewBook("Moby Dick", "Herman Melville", 300)

	//myBook.PrintInfo()

	//myBook.SetTitle("Moby Dick (Edición especial)")
	//fmt.Println(myBook.GetTitle())

	//var myBook2 = book.NewTexBook("Comunicación", "Jaime gamarra", 261, "Santillana SAC", "Secundaria")
	//myBook2.PrintInfo()

	//Usando el metodo de la interface
	//book.Print(myBook)
	//book.Print(myBook2)

	//miPerro := animal.Perro{Nombre: "Luna"}
	//miGato := animal.Gato{Nombre: "Lulu"}

	//animal.HacerSonido(&miPerro)
	//animal.HacerSonido(&miGato)

	animales := []animal.Animal{
		&animal.Perro{Nombre: "Max"},
		&animal.Gato{Nombre: "Tom"},
		&animal.Perro{Nombre: "Buddy"},
		&animal.Gato{Nombre: "Luna"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}

}
