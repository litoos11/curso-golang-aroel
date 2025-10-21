package main

import "library/animal"

func main() {
	// This is just a placeholder main function.
	// var book = book.Book{
	// 	Title:  "Go Programming",
	// 	Author: "John Doe",
	// 	Pages:  350,
	// }

	// tbook := book.NewBook("Go Programming", "John Doe", 350)
	// tbook.PrintInfo()

	// tbook.SetPages(400)
	// fmt.Printf("Updated Pages: %d\n", tbook.GetPages())

	// ebook := book.NewEBook("Learning Go", "Jane Smith", 250, 5.5, "PDF")
	// ebook.PrintInfo()

	// book.Print(tbook)
	// book.Print(ebook)

	// perro := animal.Perro{Nombre: "Rex"}
	// gato := animal.Gato{Nombre: "Miau"}

	// animal.HacerSonido(&perro)
	// animal.HacerSonido(&gato)

	animales := []animal.Animal{
		&animal.Perro{Nombre: "Rex"},
		&animal.Gato{Nombre: "Miau"},
		&animal.Perro{Nombre: "Buddy"},
		&animal.Gato{Nombre: "Whiskers"},
	}

	for _, a := range animales {
		animal.HacerSonido(a)
	}
}
