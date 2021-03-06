package main

// T: O(1)
// M: O(1)
// -- start --

// wiki adder 
// https://en.wikipedia.org/wiki/Adder_(electronics)
func getSum(a int, b int) int {
  if (a == 0) {
    return b
  }
  if (b == 0) {
    return a
  }

  return getSum((a & b) << 1, a ^ b)
}

// -- end --

