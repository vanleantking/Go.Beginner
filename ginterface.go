package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "LaLLamaQueLLama!"
}

func main() {
	animals := []Animal{Dog{}, Cat{}, Llama{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak()) // method dispatch via jmp-table
	}
}
