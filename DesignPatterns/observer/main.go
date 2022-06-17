// Observer: it is the object that is interested in the subject
// 	update
// Subject: it is the object that is being observed.
// 	state
// 	publish
// 	subscribe
// 	notify

package main

import "fmt"

func main() {
	// create a new subject
	subject := &Subject{state: "initial state"}

	// create two new observers
	o1 := Observer1{}
	o2 := Observer2{}

	// subscribe the observer to the subject
	subject.Subscribe(&o1)
	subject.Subscribe(&o2)

	// publish the subject
	subject.publish("Hello World")
	subject.publish("Have a good day!")
}

type Observer interface {
	// update is the method that will be called when the subject publishes a new state
	update(string)
}

// Observers implement the Observer interface which is update func.
type Observer1 struct{}

func (o *Observer1) update(state string) {
	fmt.Println("Observer1:", state)
}

type Observer2 struct{}

func (o *Observer2) update(state string) {
	fmt.Println("Observer2:", state)
}

type Subject struct {
	observers []Observer
	state     string
}

func (s *Subject) notify(msg string) {
	s.state = msg
	for _, observer := range s.observers {
		observer.update(msg)
	}
}

func (s *Subject) publish(msg string) {
	s.state = msg
	s.notify(msg)
}

func (s *Subject) Subscribe(observer Observer) {
	s.observers = append(s.observers, observer)
}
