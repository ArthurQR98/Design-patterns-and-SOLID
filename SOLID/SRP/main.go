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
	"log"
	"net/url"
	"os"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {}

// separation of concerns
// God Object

func (j *Journal) Save(filename string) {
	err := os.WriteFile(filename, []byte(j.String()), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func (j *Journal) Load(filename string) {}

func (j *Journal) LoadFromWeb(url *url.URL) {}

var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	err := os.WriteFile(filename, []byte(strings.Join(j.entries, LineSeparator)), 0664)
	if err != nil {
		log.Fatal(err)
	}
}

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	err := os.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0664)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	j := Journal{}
	j.AddEntry("I cried today")
	j.AddEntry("I ate a bug")

	fmt.Println(j.String())

	// Bad
	j.Save("bad.txt")

	// Good
	SaveToFile(&j, "journal.txt")

	// Good
	p := Persistence{lineSeparator: "\r\n"}
	p.SaveToFile(&j, "journal.txt")
}
