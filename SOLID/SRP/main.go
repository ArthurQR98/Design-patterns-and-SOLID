/*
Tener más de una responsabilidad en nuestras clases o módulos hace que nuestro
código sea difícil de leer, de testear y mantener. Es decir, hace que el código sea
menos flexible, más rígido y, en definitiva, menos tolerante al cambio.
El principio de responsabilidad única no se basa en crear clases con un solo método,
sino en diseñar componentes que solo estén expuestos a una fuente de cambio
*/
package main

import (
	"fmt"

	"github.com/ArthurQR98/patterns_design_go/SOLID/SRP/notification"
	"github.com/ArthurQR98/patterns_design_go/SOLID/SRP/repository"
)

type UseCase struct {
	repo     repository.Repo
	notifier notification.Notifier
}

func NewUseCase(repo repository.Repo, notifier notification.Notifier) UseCase {
	return UseCase{repo, notifier}
}

func (u *UseCase) doSomethingWithTaxes() {
	fmt.Println("Do something related with taxes ...")
}

func (u *UseCase) saveChange() {
	u.repo.Update()
}

func (u *UseCase) notify() {
	u.notifier.Notify("Te esperamos en la fiesta!!")
}

func main() {
	repo := repository.Repo{}
	notify := notification.Notifier{}
	myUseCase := NewUseCase(repo, notify)
	myUseCase.doSomethingWithTaxes()
	myUseCase.notify()
	myUseCase.saveChange()
}
