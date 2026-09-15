// Лаба 2.
package main

import "fmt"

// Функция для 2 задачи
func PNZ(usernum int) string {
	if usernum < 0 {
		return "Задание 2. Number Positive"
	}
	if usernum < 0 {
		return "Задание 2. Number Negative"
	}
	if usernum == 0 {
		return "Задание 2. Number Zero"
	}
	return "Вы ввели не число"
}

// Функция для 4 задачи
func lenlen(slovo string) int {
	return len(slovo)
}

// Структура для 5 задачи
type Rectangle struct {
	a int
	b int
}

// Функция для 5 задачи.
func pl(PL Rectangle) int {
	return PL.a * PL.b
}

// Функция для 6 задачи.
func Sravg(first int, second int) int {
	return (first + second) / 2
}

func main() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scanln(&num)
	if num%2 == 0 {
		fmt.Println("Задание 1. Число четное")
	} else {
		fmt.Println("Задание 1. Число не четное")
	}

	var usernum int
	fmt.Print("Введите число: ")
	fmt.Scanln(&usernum)
	fmt.Println("Задание 2: ", PNZ(usernum))

	for i := 1; i <= 10; i++ {
		fmt.Println("Задание 3:", i)
	}

	var slovo string
	fmt.Print("Введите строку: ")
	fmt.Scanln(&slovo)
	fmt.Println("Задание 4: ", lenlen(slovo))

	var S Rectangle
	S.a = 10
	S.b = 5
	fmt.Println("Задание 5: ", pl(S))

	var first int
	var second int
	fmt.Print("Введите число1: ")
	fmt.Scanln(&first)
	fmt.Print("Введите число2: ")
	fmt.Scanln(&second)
	fmt.Println("Задание 6: ", Sravg(first, second))
}
