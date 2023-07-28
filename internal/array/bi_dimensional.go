package array

type ArrayBiDimensional [100][100]int

var arrayBiDimensional ArrayBiDimensional

func Exec2() {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			arrayBiDimensional[i][j] = i + j
		}
	}
}
