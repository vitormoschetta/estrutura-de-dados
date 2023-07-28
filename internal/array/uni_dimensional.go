package array

import "fmt"

type ArrayUniDimensional [100]int

var arrayUniDimensional ArrayUniDimensional

func Exec1() {
	for i := 0; i < 100; i++ {
		arrayUniDimensional[i] = i
	}
	fmt.Println("Array criado com sucesso!")

	achei := false
	for i := 0; i < 10000000; i++ {
		if arrayUniDimensional[i] == 9999999 {
			achei = true
			break
		}
	}
	if achei {
		fmt.Println("Achei o número 9999999!")
	} else {
		fmt.Println("Não achei o número 9999999!")
	}
}
