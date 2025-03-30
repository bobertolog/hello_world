package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Print("Введите размер шахматной доски: ")
	fmt.Scan(&size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				fmt.Print(" ") // символ пробела для белой клетки
			} else {
				fmt.Print("#") // символ # для черной клетки
			}
		}
		fmt.Println() // переход на новую строку после каждой линии
	}
}
