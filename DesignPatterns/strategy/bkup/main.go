package main

import "fmt"

func main() {
	newToy := NewToy(SpiderMan{})
	newToy.PerformSay()
	// Change the behaviour at runtime.
	newToy.SetSuperHero(SuperMan{})
	newToy.PerformSay()
}

// Speak known how to Say a dialogue
type Speak interface {
	// Concrete types should implement this method.
	Say()
}

// SpiderMan is a concrete type that implements Say
type SpiderMan struct{}

// Say -- SpiderMan says the dialogue
func (spm SpiderMan) Say() {
	fmt.Println("SpiderMan Says: No Man Can Win Every Battle, " +
		"But No Man Should Fall Without A Struggle")
}

// SuperMan is a concrete type that implements Say
type SuperMan struct{}

// Say -- SuperMan says the dialogue
func (sum SuperMan) Say() {
	fmt.Println("SuperMan Says: There is a superhero in all of us, " +
		"we just need the courage to put on the cape")
}

type toy struct {
	// Speak is the behaviour that is encapsulated
	// Now this Speak is of interface type and hence
	// can hold any concrete type.
	// Now the concrete type implements the methods of the
	// Speak interface.
	Speak Speak
}

// NewToy returns a toy object
func NewToy(dr Speak) *toy {
	return &toy{
		Speak: dr,
	}
}

// PerformSay performs the dialogue
func (t *toy) PerformSay() {
	t.Speak.Say()
}

// SetSuperHero sets the superhero for the toy
func (t *toy) SetSuperHero(dr Speak) {
	t.Speak = dr
}
