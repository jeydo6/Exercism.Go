package elon

import "fmt"

// Drive method on the Car updates the number of meters driven based on the car's speed,
// and reduces the battery according to the battery drainage
func (car *Car) Drive() {
	if car.battery < car.batteryDrain {
		return
	}

	car.battery -= car.batteryDrain
	car.distance += car.speed
}

// DisplayDistance method on Car returns the distance displayed on the LED display
func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery method on Car returns the battery percentage displayed on the LED display
func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// CanFinish method that takes a trackDistance as its parameter
// and returns true if the car can finish the race; otherwise, return false
func (car *Car) CanFinish(trackDistance int) bool {
	maxDistance := car.battery * car.speed / car.batteryDrain
	return maxDistance >= trackDistance
}
