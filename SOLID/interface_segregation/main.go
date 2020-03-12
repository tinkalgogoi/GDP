package main

import "fmt"

type VehicleCDPlayer interface {
	PlayCD()
}
type VehicleCdPlay struct {
}

func (v VehicleCdPlay) PlayCD() {
	fmt.Println("Playing Guns n Roses")
}

//t- not combining both methods in single interface
type VehicleAccelerator interface {
	Accelerate()
}
type VehicleAccelerate struct {
}

func (v VehicleAccelerate) Accelerate() {
	fmt.Println("Accelerating ....")
}

type Car struct {
	VehicleAccelerate
	VehicleCdPlay
}

type Bike struct {
	// Clients should not be forced to depend on methods they do not use. –Robert C. Martin
	// bike don't play CD and only accelerate
	VehicleAccelerate
}

func main() {
	var c Car
	var b Bike

	c.Accelerate()
	c.PlayCD()

	b.Accelerate()
}
