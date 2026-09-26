package speed

// Car define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{speed: speed, batteryDrain: batteryDrain, battery: 100, distance: 0}
}

// Track define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{distance}
}

// Drive drives the car one time. If there is not enough battery to drive one more time, the car will not move.
func Drive(car Car) Car {
	isBatteryRunOut := car.battery < car.batteryDrain
	if isBatteryRunOut {
		return car
	}
	car.battery -= car.batteryDrain
	car.distance += car.speed
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	finishedRaceTime := track.distance / car.speed
	finishedRaceNeedBattery := finishedRaceTime * car.batteryDrain
	isFinshedRace := car.battery >= finishedRaceNeedBattery

	return isFinshedRace
}
