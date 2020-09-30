package hw01

import (
	"fmt"
	"math"
)

func Ma() float32 {
	var kat1 float64 = 5.00
	var kat2 float64 = 3.00

	fmt.Println("Математика. Задача №2")
	fmt.Printf("Площадь прям.треугольника = %.2f\n", (kat1*kat2)/2)
	fmt.Printf("Периметр прям.треугольника = %.2f\n ", (kat1 + kat2))
	fmt.Printf("Гипотенуза прям.треугольника = %.2f\n ", math.Sqrt(math.Pow(kat1, 2)+math.Pow(kat2, 2)))

	return 0
}
