// Solved by Github YoonBaek

package main

import "fmt"

type Truck string

func (t Truck) Accelerate() {
	fmt.Println("Speeding Up")
}

func (t Truck) Brake() {
	fmt.Println("Stopping")
}

func (t Truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}

func (t Truck) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(string)
}

func TryVehicle(vehicle Vehicle) {
	vehicle.Accelerate()
	vehicle.Steer("left")
	vehicle.Steer("right")
	vehicle.Brake()
	truck, ok := vehicle.(Truck)
	if ok {
		truck.LoadCargo("Test Cargo")
	}
}

func main() {
	TryVehicle(Truck("Ford F180"))
}
