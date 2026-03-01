package main

import "fmt"

// Animal contains all the base fields for animals.
type Animal struct {
	Name     string
	IsMammal bool
}

// Speak provides generic behavior for all animals and
// how they speak.
func (a Animal) Speak() {
	fmt.Println("UGH!",
		"My name is", a.Name,
		", it is", a.IsMammal,
		"I am a mammal")
}

// Dog contains everything an Animal is but specific
// attributes that only a Dog has.
type Dog struct {
	Animal
	PackFactor int
}

// Speak knows how to speak like a dog.
func (d Dog) Speak() {
	fmt.Println("Woof!",
		"My name is", d.Name,
		", it is", d.IsMammal,
		"I am a mammal with a pack factor of", d.PackFactor)
}

// Cat contains everything an Animal is but specific
// attributes that only a Cat has.
type Cat struct {
	Animal
	ClimbFactor int
}

// Speak knows how to speak like a cat.
func (c Cat) Speak() {
	fmt.Println("Meow!",
		"My name is", c.Name,
		", it is", c.IsMammal,
		"I am a mammal with a climb factor of", c.ClimbFactor)
}

func main() {

	// Create a Dog by initializing its Animal parts
	// and then its specific Dog attributes.
	d := Dog{
		Animal: Animal{
			Name:     "Fido",
			IsMammal: true,
		},
		PackFactor: 5,
	}

	// Create a Cat by initializing its Animal parts
	// and then its specific Cat attributes.
	c := Cat{
		Animal: Animal{
			Name:     "Milo",
			IsMammal: true,
		},
		ClimbFactor: 4,
	}

	// Have the Dog and Cat speak.
	d.Speak()
	c.Speak()
}
