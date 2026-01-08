package chance

import "math/rand"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return rand.Intn(20) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return rand.Float64() * 12.0
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := make([]string, 8)
    animals[0] = "ant"
    animals[1] = "beaver"
    animals[2] = "cat"
    animals[3] = "dog"
    animals[4] = "elephant"
    animals[5] = "fox"
    animals[6] = "giraffe"
    animals[7] = "hedgehog"
    rand.Shuffle(len(animals), func(i, j int) {
        animals[i], animals[j] = animals[j], animals[i]
    })
    return animals
}
