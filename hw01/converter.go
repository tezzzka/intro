package hw01

import "fmt"

func Calc() float32 {

	const tail *float32 = 82.32
	var rub *int

	fmt.Print("Введите сумму в рублях: ")
	fmt.Scanln(&rub)

	return *rub * *tail

}
