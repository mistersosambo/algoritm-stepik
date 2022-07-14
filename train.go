package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var a, b int
	fmt.Print("Введите число a: ")
	fmt.Scan(&a)
	fmt.Print("Введите число b: ")
	fmt.Scan(&b)
	aNew := a
	bNew := b
	count1 := 0
	count2 := 0
	for a != 0 {
		count1 += 1
		a = a / 10
	}
	for b != 0 {
		count2 += 1
		b = b / 10
	}
	j1 := 0
	r1 := int(math.Pow10(count1 - 1))
	for j1 != count1 {
		s1 := aNew / r1
		m1 := s1 % 10
		j1 += 1
		r1 /= 10
		j2 := 0
		r2 := int(math.Pow10(count2 - 1))
		for j2 != count2 {
			s2 := bNew / r2
			m2 := s2 % 10
			j2 += 1
			r2 /= 10
			if m1 == m2 {
				fmt.Print(strconv.Itoa(m1), " ")
			} else {
				continue
			}
		}
	}
}
