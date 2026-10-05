package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(lasagnaOfLayer []string, timeOfLayer int) int {

	if timeOfLayer == 0 {
		timeOfLayer = 2
	}

	numberOfLayers := len(lasagnaOfLayer)
	return numberOfLayers * timeOfLayer
}

// TODO: define the 'Quantities()' function
func Quantities(lasagnaOfLayer []string) (int, float64) {
	quantityOfNoodles := 0
	quantityOfSauce := 0.0

	for _, value := range lasagnaOfLayer {
		if value == "noodles" {
			quantityOfNoodles += 50
		} else if value == "sauce" {
			quantityOfSauce += 0.2
		} else {
		}
	}

	return quantityOfNoodles, quantityOfSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
func ScaleRecipe(quantities []float64, numberOfPortion int) []float64 {
	ratio := float64(numberOfPortion) / 2.0
	scaled := make([]float64, 0, len(quantities))

	for _, amount := range quantities {
		scaled = append(scaled, amount*ratio)
	}
	return scaled
}
