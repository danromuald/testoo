package main

type Vehiche interface {
	Drive() string
	Start() string
	Stop() string
}

type Car struct {
	n    int
	name string
}

type Bike struct {
	n    int
	name string
}

func (r *Car) Drive() string {
	return "Driving a car!"
}

func (c *Car) Start() string {
	c.n = 1
	return "Started, set n to 1"
}

func (c *Car) Stop() string {
	return "Car Stopped!"
}
