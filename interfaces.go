package figuras

import (
	"fmt"
)

type geometria interface {
	area() float64
	perimetro() float64
}

func Medidas(g geometria) {
	fmt.Println("Medidas:", g)
	fmt.Printf("Area: %.2f\n", g.area())
	fmt.Printf("Perímetro: %.2f\n", g.perimetro())
}
