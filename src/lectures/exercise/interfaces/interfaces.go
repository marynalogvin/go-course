//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:

//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

// * The shop has lifts for multiple vehicle sizes/types:
//   - Motorcycles: small lifts
//   - Cars: standard lifts
//   - Trucks: large lifts

type LiftPicker interface {
	PickLift()
}

type Motorcycle string
type Car string
type Truck string

// * Vehicles have a model name in addition to the vehicle type:
//   - Example: "Truck" is the vehicle type, "Road Devourer" is a model name

func (m Motorcycle) String() string {
	return fmt.Sprintf("motorcycle: %v", string(m))
}
func (c Car) String() string {
	return fmt.Sprintf("car: %v", string(c))
}
func (t Truck) String() string {
	return fmt.Sprintf("truck: %v", string(t))
}

func (m Motorcycle) PickLift() {
	fmt.Printf("send %v to small lift\n", m)
}
func (c Car) PickLift() {
	fmt.Printf("send %v to standard lift\n", c)
}
func (t Truck) PickLift() {
	fmt.Printf("send %v to large lift\n", t)
}

//   - Write a single function to handle all of the vehicles
//     that the shop works on.

func sendToLift(p LiftPicker) {
	p.PickLift()

}
func main() {
	//* Direct at least 1 of each vehicle type to the correct
	//  lift, and print out the vehicle information.
	motorcycle := Motorcycle("Retro")
	car := Car("Sporty")
	truck := Truck("Fire truck")
	sendToLift(motorcycle)
	sendToLift(car)
	sendToLift(truck)

}
