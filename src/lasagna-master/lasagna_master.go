package lasagna

func PreparationTime(layers []string, timePerLayer int) int {

	if timePerLayer > 0 {
		return timePerLayer * len(layers)
	}

	return 2 * len(layers)
}

func Quantities(layers []string) (int, float64) {

	var noodlesCount int
	var sauceCount int

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodlesCount++
		} else if layers[i] == "sauce" {
			sauceCount++
		}
	}

	return 50 * noodlesCount, 0.2 * float64(sauceCount)
}

func AddSecretIngredient(friendsList []string, myList []string) {

	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {

	var quantitiesLength = len(quantities)
	var multiplier = float64(portions) / 2.0

	var result = make([]float64, quantitiesLength)
	for i := 0; i < quantitiesLength; i++ {
		result[i] = multiplier * quantities[i]
	}

	return result
}
