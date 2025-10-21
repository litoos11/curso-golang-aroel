package animal

import "fmt"

type Animal interface {
	Sonido()
}

type Perro struct {
	Nombre string
}

func (p *Perro) Sonido() {
	fmt.Printf("%s dice: Guau Guau\n", p.Nombre)
}

type Gato struct {
	Nombre string
}

func (g *Gato) Sonido() {
	fmt.Printf("%s dice: Miau Miau\n", g.Nombre)
}

func HacerSonido(a Animal) {
	a.Sonido()
}
