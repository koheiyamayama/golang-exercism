package lasagna

// TODO: define the 'PreparationTime()' function

func PreparationTime(layers []string, avgMinutes int) int {
	layersNum := len(layers)
	if avgMinutes == 0 {
		return layersNum * 2
	} else {
		return layersNum * avgMinutes
	}
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodlesCount := 0
	sauceCount := 0

	for _, v := range layers {
		if v == "sauce" {
			sauceCount = sauceCount + 1
		} else if v == "noodles" {
			noodlesCount = noodlesCount + 1
		}
	}

	return noodlesCount * 50, float64(sauceCount) * 0.2
}

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(friendLayers []string, myLayers []string) {
	secretIngredient := friendLayers[len(friendLayers)-1]
	myLayers[len(myLayers)-1] = secretIngredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, scale int) []float64 {
	newQuantities := make([]float64, len(quantities))
	copy(newQuantities, quantities)
	for i, v := range newQuantities {
		newQuantities[i] = v * float64(scale) / 2
	}

	return newQuantities
}
