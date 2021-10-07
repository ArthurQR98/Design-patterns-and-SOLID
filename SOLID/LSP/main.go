/*
“Las funciones que utilicen punteros o referencias a clases base deben ser
capaces de usar objetos de clases derivadas sin saberlo”
Si una clase A es extendida por una clase
B, debemos de ser capaces de sustituir cualquier instancia de A por cualquier objeto
de B sin que el sistema deje de funcionar o se den comportamientos inesperados
*/
package main

import "fmt"

type area interface {
	getArea() int
}
type Rectangle struct {
	width  int
	height int
}

type Square struct {
	value int
}

func (r *Rectangle) getArea() int {
	return r.width * r.height
}

func (s *Square) getArea() int {
	return s.value * s.value
}

func main() {
	rectangle := Rectangle{4, 5}
	square := Square{5}
	fmt.Println(rectangle.getArea())
	fmt.Println(square.getArea())
}
