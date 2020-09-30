package hw01

import (
	"fmt"
)

func Calc() float32 {

	const tail float32 = 82.32
	var rub int

	fmt.Print("Введите сумму в рублях: ")
	fmt.Scanln(&rub)
	fmt.Printf("Результат: $%.2f USD ", float32(rub)/tail)

	return 0

}
