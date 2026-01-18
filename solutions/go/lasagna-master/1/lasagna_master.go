package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTime int) int {
    if avgTime == 0 {
        return 2 * len(layers)
    } else {
        return avgTime * len(layers)
    }
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    // This function will return the quantities
    total_water := 0
    var total_sauce = 0.0
    for i := 0 ; i < len(layers); i++ {
        if layers[i] == "noodles" {
            total_water = total_water + 50
        } else if layers[i] == "sauce" {
            total_sauce = total_sauce + 0.2
        }
    }
    return total_water, total_sauce
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
    last_item_in_friends := friendsList[len(friendsList) - 1]
    // We are gonna replace the '?' in the myList
    myList[len(myList) - 1] = last_item_in_friends
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numOfPortions int) []float64 {
    var actual_portions float64 = float64(numOfPortions) / 2
    var our_new_slice = make([]float64, len(quantities))
    for i := 0 ; i < len(quantities); i++ {
        our_new_slice[i] = quantities[i] * actual_portions
    }
    return our_new_slice
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
