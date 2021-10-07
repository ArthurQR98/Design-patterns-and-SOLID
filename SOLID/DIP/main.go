/*
“Los módulos de alto nivel no deben depender de módulos de bajo nivel. Ambos
deben depender de abstracciones. Las abstracciones no deben depender de
concreciones. Los detalles deben depender de abstracciones”
*/

// HLM should not dependend on LLM
// Both should dependend on abstractions

package main

import "fmt"

type Relationship int

const (
	Parent Relationship = iota
	Child
	Silbling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// Low-level module
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}
type RelationShips struct {
	relations []Info
}

func (r *RelationShips) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

func (r *RelationShips) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// High-level module
type Research struct {
	// break DIP
	// relationship RelationShips
	browser RelationshipBrowser
}

func (r *Research) Investigate() {
	// relations := r.relationship.relations
	// for _, rel := range relations {
	// 	if rel.from.name == "John" && rel.relationship == Parent {
	// 		fmt.Println("John has a child called ", rel.to.name)
	// 	}
	// }
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.name)
	}
}

func main() {
	parent := Person{"John"}
	chil1 := Person{"Chris"}
	chil2 := Person{"Matt"}
	relationships := RelationShips{}
	relationships.AddParentAndChild(&parent, &chil1)
	relationships.AddParentAndChild(&parent, &chil2)
	r := Research{&relationships}
	r.Investigate()
}
