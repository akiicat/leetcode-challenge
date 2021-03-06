package main

import (
	. "main/pkg/testing_helper"
	"testing"
)

func TestSpiralOrder(_t *testing.T) {
	i, o := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	T(_t, S(i), S(spiralOrder(i)), S(o))

	i, o = [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
	}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}
	T(_t, S(i), S(spiralOrder(i)), S(o))

}
