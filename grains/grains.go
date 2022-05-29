package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number <= 0 || 64 < number {
		return 0, errors.New("number must be greater than 0")
	}

	return uint64(1 * math.Pow(float64(2), float64(number-1))), nil
}

func Total() uint64 {
	var sum uint64 = 0

	for i := 0; i <= 64; i++ {
		n, err := Square(i)
		if err == nil {
			sum += n
		}
	}
	return sum
	// return uint64(-1 * (1 - math.Pow(float64(2), float64(64))))
}
