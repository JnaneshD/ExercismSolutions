package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	// Car and truck needs a license to be operated
    if (kind == "car" || kind == "truck") {
        return true
    }
    return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	// Byte by byte comparision is done by default and its called lexicographical order
    var suffix string = " is clearly the better choice."
    if option1 < option2 {
        return option1 + suffix
    }
    return option2 + suffix
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if (int(age) < 3) {
        return 0.8 * originalPrice
    } else if(int(age) < 10 && int(age) >=3) {
        return 0.7 * originalPrice
    } else {
        return 0.5 * originalPrice
    }
}
