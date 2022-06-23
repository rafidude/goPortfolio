package main

import (
	"errors"
	"strconv"
)

func main() {

	// Create a new stockMonitor object
	stockMonitor := &StockMonitor{
		ticker: "AAPL",
		price:  100.0,
	}

	observerA := &StockObserver{
		name: "A",
	}
	observerB := &StockObserver{
		name: "B",
	}

	// Subscribe our Observers to the stockMonitor
	stockMonitor.Subscribe(observerA)
	stockMonitor.Subscribe(observerB)

	// Start the stockMonitor
	stockMonitor.Notify()

	// Change the price of the stockMonitor
	stockMonitor.SetPrice(500)

	// Unsubscribe an Observer from the stockMonitor
	stockMonitor.Unsubscribe(observerA)

	// Change the price of the stockMonitor
	stockMonitor.SetPrice(528)
}

// Subject interface
type Subject interface {
	Subscribe(o Observer) (bool, error)
	Unsubscribe(o Observer) (bool, error)
	Notify() (bool, error)
}

// Observer Interface
type Observer interface {
	Update(string)
}

// Concrete Observer: StockObserver
type StockObserver struct {
	name string
}

func (s *StockObserver) Update(t string) {
	// do something
	println("StockObserver:", s.name, "updated,", "received:", t)
}

// Concrete Subject: stockMonitor
type StockMonitor struct {
	// internal state
	ticker string
	price  float64

	observers []Observer
}

func (s *StockMonitor) Subscribe(o Observer) (bool, error) {

	for _, observer := range s.observers {
		if observer == o {
			return false, errors.New("Observer already exists")
		}
	}
	s.observers = append(s.observers, o)
	return true, nil
}

func (s *StockMonitor) Unsubscribe(o Observer) (bool, error) {

	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("Observer not found")
}

func (s *StockMonitor) Notify() (bool, error) {
	for _, observer := range s.observers {
		observer.Update(s.String())
	}
	return true, nil
}

func (s *StockMonitor) SetPrice(price float64) {
	s.price = price
	s.Notify()
}

func (s *StockMonitor) String() string {
	convertFloatToString := strconv.FormatFloat(s.price, 'f', 2, 64)
	return "StockMonitor: " + s.ticker + " $" + convertFloatToString
}
