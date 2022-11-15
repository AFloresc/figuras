package figuras

import (
	"fmt"
)

type Geometria interface {
	area() float64
	perimetro() float64
}

func Medidas(g Geometria) {
	fmt.Println(g)
	fmt.Printf("Area: %.2f\n", g.area())
	fmt.Printf("Perímetro: %.2f\n", g.perimetro())
}
