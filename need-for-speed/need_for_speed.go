package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery,
	batteryDrain,
	speed,
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{battery: 100, batteryDrain: batteryDrain, speed: speed}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	return Track{distance: distance}
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	var remainingBattery = car.battery - car.batteryDrain
	if remainingBattery < 0 {
		return car
	}

	car.battery = remainingBattery
	car.distance = car.distance + car.speed
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	var drivableDistance = car.battery / car.batteryDrain * car.speed

	return track.distance <= drivableDistance
}
