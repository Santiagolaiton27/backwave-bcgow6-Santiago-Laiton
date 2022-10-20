package calculador

import (
	"fmt"
	"sort"
)

func Add(a, b int) int {
	return a + b
}

func Substrab(a, b int) int {
	return a - b
}

func Sort(list []int) []int {
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	return list
}
func Split(a, b int) (float64, error) {
	if a == 0 || b == 0 {
		return 0, fmt.Errorf("ninguno de los numeros puede ser 0 : a %d, %d", a, b)
	}
	return float64(a / b), nil
}
