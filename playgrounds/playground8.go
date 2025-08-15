package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error()string {
	return fmt.Sprintf("The given num is negative: %v",float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x == 0 {
		return 0, nil
	}

	if x < 0 {
		return 0,ErrNegativeSqrt(x)
	}

	z := x / 2
	epsilon := 1e-10
	for {
		newZ := (z + x/z) / 2
		if abs(newZ-z) < epsilon {
			break
		}
		z = newZ
	}
	return z, nil
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}
