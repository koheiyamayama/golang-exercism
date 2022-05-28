package raindrops

import "fmt"

func Convert(number int) string {
	results := ""

	if number%3 == 0 {
		results = results + "Pling"
	}

	if number%5 == 0 {
		results = results + "Plang"
	}

	if number%7 == 0 {
		results = results + "Plong"
	}

	if results == "" {
		results = fmt.Sprintf("%d", number)
	}

	return results
}
