/*
“Todas las entidades software deberían estar abiertas a extensión, pero cerradas
a modificación”

Nos recomienda que, en los casos en los que se introduzcan nuevos comportamientos
en sistemas existentes, en lugar de modificar los componentes antiguos, se deben
crear componentes nuevos. La razón es que si esos componentes o clases están
siendo usadas en otra parte (del mismo proyecto o de otros) estaremos alterando
su comportamiento y provocando efectos indeseados.
*/

package main

import "fmt"

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct{}

// Bad
func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// Good
type Specification interface {
	IsSatisfed(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfed(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfed(p *Product) bool {
	return p.size == s.size
}

type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfed(p *Product) bool {
	return a.first.IsSatisfed(p) && a.second.IsSatisfed(p)
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfed(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	// Bad
	fmt.Printf("Green products (old):\n")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is Green\n", v.name)
	}
	// Good
	fmt.Printf("Green products (new):\n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is Green\n", v.name)
	}

	largeSpec := SizeSpecification{large}
	lgSpec := AndSpecification{greenSpec, largeSpec}
	fmt.Printf("Large green products\n")
	for _, v := range bf.Filter(products, lgSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}

}
