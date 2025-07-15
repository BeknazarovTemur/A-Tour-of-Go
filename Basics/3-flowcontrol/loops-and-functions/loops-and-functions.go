package main

import "fmt"

func Sqrt(x float64) float64 {
	z := 1.0
	prevZ := 0.0
	for diff := 1.0; diff > 1e-10; {
		z -= (z*z - x) / (2 * z)
		diff = z - prevZ
		if diff < 0 {
			diff = -diff
		}
		prevZ = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
