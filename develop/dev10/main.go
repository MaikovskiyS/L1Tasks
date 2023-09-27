package main

import "fmt"

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.

	Объединить данные значения в группы с шагом в 10 градусов.
	Последовательность в подмножноствах не важна.

	Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/
func main() {
	deg := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 65.4, -44.3, -3.0}
	groups := make(map[int][]float64)
	for _, v := range deg {
		name := groupValue(v)
		group, ok := groups[name]
		if !ok {
			groups[name] = make([]float64, 0)
			groups[name] = append(groups[name], v)
		} else {
			group = append(group, v)
			groups[name] = group
		}

	}
	fmt.Println(groups)

}
func groupValue(v float64) int {
	toInt := ((v * 10) / 100)
	groupname := int(toInt) * 10
	return groupname
}
