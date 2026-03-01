There are aspects of Go that don’t allow type hierarchies to provide the same level of functionality they do in other object oriented programming languages. Specifically, **the concepts of base types and subtyping don’t exist in Go** so type reuse requires a different way of thinking.

We will see why type hierarchies are not always the best pattern to use in Go. It is better to group concrete types together not by a common state but by a common behavior. We'll see how to leverage interfaces to group and decouple concrete types, and lastly, we will look at some guidelines around declaring types.

# Part 1
Let’s start with a program we see way too often from those trying to learn Go. This program uses a traditional type hierarchy pattern that would be commonly seen in an object oriented program.

```go
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
```

In above code we see the beginning of our traditional object oriented program. We have the declaration of the concrete type **Animal** and it has two fields, **Name** and **IsMammal**. Then we have a method named **Speak** that allows an Animal to talk. 

Since an **Animal** is a base type for all animals, the implementation of the **Speak** method is generic and can’t represent any given animal very well beyond this base state.

Then we have a new concrete type named **Dog** which embeds a value of type **Animal** and has a unique field named **PackFactor**. 
We see the use of composition to reuse the fields and methods of the Animal type. 

In this case, composition is providing some of the same benefits of inheritance, with respect to type reuse. The **Dog** type also implements its own version of the **Speak** method, which allows the **Dog** to bark like a dog. This method is overriding the implementation of the **Animal** type’s **Speak** method.

Next we have a third concrete type named **Cat** that also embeds a value of type **Animal** and has a field named **ClimbFactor**. 
We see the use of composition again for the same reasons, and **Cat** has a method named **Speak** that allows the **Cat** to meow like a cat. Again, this method is overriding the implementation of the **Animal** type’s **Speak** method.

Then we have the **main** function where we put everything together. We create a value of type **Dog** using a struct literal and initialize the embedded **Animal** value and the **PackFactor** field. We create a value of type **Cat** using a struct literal and initialize the embedded **Animal** value and the **ClimbFactor** field. Then, finally, we call the **Speak** method against the **Dog** and **Cat** values.

This works in Go, and you can see how the use of embedding types provides familiar type hierarchy functionality. However there are some flaws with doing this in Go, and one is that Go does not support the idea of subtyping. This means you can’t use the Animal type as a base type like you can in other object oriented languages.

What is important to understand is that, in Go, the **Dog and Cat types can’t be used as a value of type Animal**. What we have is an embedded value of type Animal inside a value of type Dog and Cat. **You can’t pass a Dog or Cat to any function that accepts values of type Animal**. This also means that there is no way to group a set of Cat and Dog values together in the same list via the Animal type.

```go
// Attempt to use Animal as a base type.
animals := []Animal{
    Dog{},
    Cat{},
}

: cannot use Dog literal (type Dog) as type Animal in array or slice literal
: cannot use Cat literal (type Cat) as type Animal in array or slice literal
```

Above code shows what happens in Go when you try to use the **Animal** type as a base type in a traditional object oriented way. The compiler is very clear that the Dog and Cat types can’t be used as type Animal.

The **Animal** type and the use of type hierarchies in this case is not providing us any real value. I would argue it is leading us down a path of code that is not readable, simple or adaptable.

# Part 2
When coding in Go try to avoid these type hierarchies that promote the idea of common state and think about common behavior. 
We can group a set of **Dog** and **Cat** values if we think about the common behavior they exhibit. In this case there is a common behavior of **Speak**.

Let’s look at another implementation of this code that focuses on behavior.

```go
package main

import "fmt"

// Speaker provide a common behavior for all concrete types
// to follow if they want to be a part of this group. This
// is a contract for these concrete types to follow.
type Speaker interface {
 Speak()
}
```

The new program we have added a new type called **Speaker**. This is not a concrete type like the struct types we declared before. This is an interface type that declares a contract of behavior that will let us group and work with a set of different concrete types that implement the Speak method.

```go
// Dog contains everything a Dog needs.
type Dog struct {
 Name       string
 IsMammal   bool
 PackFactor int
}

// Speak knows how to speak like a dog.
// This makes a Dog now part of a group of concrete
// types that know how to speak.
func (d Dog) Speak() {
 fmt.Println("Woof!",
 	"My name is", d.Name,
 	", it is", d.IsMammal,
 	"I am a mammal with a pack factor of", d.PackFactor)
}

// Cat contains everything a Cat needs.
type Cat struct {
 Name        string
 IsMammal    bool
 ClimbFactor int
}

// Speak knows how to speak like a cat.
// This makes a Cat now part of a group of concrete
// types that know how to speak.
func (c Cat) Speak() {
 fmt.Println("Meow!",
 	"My name is", c.Name,
 	", it is", c.IsMammal,
 	"I am a mammal with a climb factor of", c.ClimbFactor)
}
```

In above code we have the declaration of the concrete types **Dog** and **Cat** again. This code removes the **Animal** type and copies those common fields directly into Dog and Cat.

Why did we do that?
- The Animal type was providing an abstraction layer of reusable state.
- The program never needed to create or solely use a value of type Animal.
- The implementation of the **Speak** method for the **Animal** type was a generalization.
- The Speak method for the **Animal** type was never going to be called.

Here are some guidelines around declaring types:
- **Declare types that represent something new or unique**.
- Validate that a value of any type is created or used on its own.
- **Embed types to reuse existing behaviors you need to satisfy**.
- Question types that are an alias or abstraction for an existing type.
- Question types whose sole purpose is to share common state.

Let’s look at the main function now.
```go
func main() {

 // Create a list of Animals that know how to speak.
 speakers := []Speaker{

 	// Create a Dog by initializing its Animal parts
 	// and then its specific Dog attributes.
 	Dog{
 		Name:       "Fido",
 		IsMammal:   true,
 		PackFactor: 5,
 	},

 	// Create a Cat by initializing its Animal parts
 	// and then its specific Cat attributes.
 	Cat{
 		Name:        "Milo",
 		IsMammal:    true,
 		ClimbFactor: 4,
 	},
 }

 // Have the Animals speak.
 for _, spkr := range speakers {
 	spkr.Speak()
 }
}
```

Above, we create a slice of **Speaker** interface values to group both **Dog** and **Cat** values together under their common behavior. We create a value of type **Dog** and a value of type **Cat**. Finally we iterate over the slice of **Speaker** interface values and have the **Dog** and **Cat** speak.

Some final thoughts about the changes we made:
- We didn’t need a base type or type hierarchies to group concrete type values together.
- The Interface allowed us to create a slice of different concrete type values and work with these values through their common behavior.
- We removed any type pollution by not declaring a type that was never solely used by the program.

# Conclusion
There is a lot more to composition in Go but this is an initial understanding around the problems with using type hierarchies. There are always exceptions to every rule, but try to follow these guidelines until you know enough to understand the tradeoffs for making an exception.

To learn more about composition and other topics this post touches, check out these other blog posts:

https://www.ardanlabs.com/blog/2014/03/exportedunexported-identifiers-in-go.html
https://www.ardanlabs.com/blog/2014/05/methods-interfaces-and-embedded-types.html
https://www.ardanlabs.com/blog/2015/03/object-oriented-programming-mechanics.html
https://www.ardanlabs.com/blog/2015/09/composition-with-go.html


