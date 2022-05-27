package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.

type SillyNephewError struct {
	negativeAmount int
}

func (sn *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", sn.negativeAmount)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	amount, err := weightFodder.FodderAmount()

	if cows == 0 {
		return 0, errors.New("division by zero")
	} else if cows < 0 {
		return 0, &SillyNephewError{negativeAmount: cows}
	}

	if err == nil && amount >= 0 {
		return amount / float64(cows), nil
	} else if err == ErrScaleMalfunction && amount > 0 {
		return amount * 2 / float64(cows), nil
	} else if err != ErrScaleMalfunction && amount > 0 {
		return 0, err
	} else if (err == ErrScaleMalfunction || err == nil) && amount <= 0 {
		return 0, errors.New("negative fodder")
	}

	return 0, err
}
