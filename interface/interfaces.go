package main

import "fmt"

type animal interface {
	mover() string
}

type perro struct{}
type pez struct{}
type pajaro struct{}

func (perro) mover() string {
	return "Perro moviendose"
}

func (pez) mover() string {
	return "Pez moviendose"
}

func (pajaro) mover() string {
	return "pajaro moviendose"
}

func moverAnimal(a animal) {
	fmt.Println(a.mover())
}

func main() {
	dog := perro{}
	fish := pez{}
	bird := pajaro{}
	moverAnimal(dog)
	moverAnimal(fish)
	moverAnimal(bird)
}
