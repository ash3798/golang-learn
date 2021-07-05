package main

import "fmt"

type Animal struct {
	legs int
}

func (a *Animal) Eat() {
	fmt.Println("eat func")
}

func (a *Animal) Sleep() {
	fmt.Println("sleep func")
}

func (a *Animal) Run() {
	fmt.Println("run func")
}

type Dog struct {
	animal Animal //adding animal object into dog to get properties of animal COMPOSITION
	name   string
}

func TestInheritance() {
	dog := Dog{animal: Animal{legs: 10}, name: "bruno"}
	fmt.Printf("dog object: %+v \n", dog)

	fmt.Println("name :", dog.name)
	fmt.Println("legs :", dog.animal.legs)
}
