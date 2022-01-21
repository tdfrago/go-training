package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, ave int) int {
	if ave == 0 {
		return len(layers) * 2
	} else {
		return len(layers) * ave
	}
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	grams := 0
	liters := 0.0
	for _, value := range layers {
		if value == "noodles" {
			grams += 50
		} else if value == "sauce" {
			liters += 0.2
		}
	}
	return grams, liters
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portion int) []float64 {
	scaledQuantities := []float64{}
	for _, value := range quantities {
		scaledQuantities = append(scaledQuantities, value*(float64(portion)/2))
	}
	return scaledQuantities
}
