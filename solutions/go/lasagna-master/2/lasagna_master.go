package lasagnamaster

// PreparationTime returns the time needed to cook the lasagna.
func PreparationTime(lasagnaOfLayer []string, timeOfLayer int) int {
	if timeOfLayer == 0 {
		timeOfLayer = 2
	}
	return len(lasagnaOfLayer) * timeOfLayer
}

// Quantities returns the amounts of noodles and sauce needed.
func Quantities(lasagnaOfLayer []string) (int, float64) {
	quantityOfNoodles := 0
	quantityOfSauce := 0.0

	for _, value := range lasagnaOfLayer {
		switch value {
		case "noodles":
			quantityOfNoodles += 50
		case "sauce":
			quantityOfSauce += 0.2
		}
	}

	return quantityOfNoodles, quantityOfSauce
}

// AddSecretIngredient replaces the last ingredient in myList with
// the last ingredient in friendsList.
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// ScaleRecipe returns the ingredients scaled to the given number of portions.
func ScaleRecipe(quantities []float64, numberOfPortion int) []float64 {
	ratio := float64(numberOfPortion) / 2.0
	scaled := make([]float64, 0, len(quantities))

	for _, amount := range quantities {
		scaled = append(scaled, amount*ratio)
	}
	return scaled
}
