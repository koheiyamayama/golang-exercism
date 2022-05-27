package purchase

import (
	"fmt"
	"strings"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var name string
	switch strings.Compare(option1, option2) {
	case 0:
		name = option1
	case -1:
		name = option1
	case 1:
		name = option2
	default:
		name = option1
	}

	return fmt.Sprintf("%s is clearly the better choice.", name)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return float64(originalPrice) * 80.0 / 100.0
	} else if 3 <= age && age < 10 {
		return float64(originalPrice) * 70.0 / 100.0
	} else {
		return float64(originalPrice) * 50.0 / 100.0
	}
}
