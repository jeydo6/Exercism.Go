package thefarm

import (
	"errors"
	"fmt"
)

type SillyNephewError struct {
	Cows int
}

func (sn *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", sn.Cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodderAmount, err := weightFodder.FodderAmount()

	if err != nil && err != ErrScaleMalfunction {
		return 0, err
	}
	if err == ErrScaleMalfunction {
		fodderAmount *= 2
	}

	if fodderAmount < 0 {
		return 0, errors.New("negative fodder")
	}
	if cows == 0 {
		return 0, errors.New("division by zero")
	}
	if cows < 0 {
		return 0, &SillyNephewError{Cows: cows}
	}

	return fodderAmount / float64(cows), nil
}
