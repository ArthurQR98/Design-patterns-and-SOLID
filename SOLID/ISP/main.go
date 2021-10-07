/*
“Los clientes no deberían estar obligados a depender de interfaces que no
utilicen”
Este principio viene a decir que una clase no
debería depender de métodos o propiedades que no necesita. Por lo tanto, cuando
definimos el contrato de una interfaz, debemos centrarnos en las clases que la van a
usar (las interfaces pertenecen a la clase cliente), no en las implementaciones que ya
tenemos desarrolladas.
*/

package main

import "fmt"

type Car interface {
	accelerate()
	brake()
	startEngine()
}

type Tesla interface {
	autoPilot()
	ludicrousSpeed()
}

type Mustang struct{}
type ModelS struct{}

func (m *Mustang) accelerate() {
	fmt.Println("Speeding up ...")
}

func (m *Mustang) brake() {
	fmt.Println("Stopping ...")
}

func (m *Mustang) startEngine() {
	fmt.Println("Starting engine")
}

func (m *Mustang) Run() {
	m.startEngine()
	m.brake()
	m.accelerate()
}

// For cars of Tesla
func (t *ModelS) accelerate() {
	fmt.Println("Speeding up ...")
}

func (t *ModelS) brake() {
	fmt.Println("Stopping ...")
}

func (t *ModelS) startEngine() {
	fmt.Println("Starting engine")
}

func (t *ModelS) autoPilot() {
	fmt.Println("Self driving ...")
}

func (t *ModelS) ludicrousSpeed() {
	fmt.Println("Woow!!")
}

func (t *ModelS) Run() {
	t.startEngine()
	t.brake()
	t.accelerate()
	t.autoPilot()
	t.ludicrousSpeed()
}

func main() {
	mustang := Mustang{}
	mustang.Run()
	fmt.Println("************************")
	modelS := ModelS{}
	modelS.Run()
}
