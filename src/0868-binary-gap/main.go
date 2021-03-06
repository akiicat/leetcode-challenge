package main

// T: O(1) for 32
// M: O(1)
// -- start --

func binaryGap(N int) int {
  max, c := -1, 0

  for N != 0 {
    v := N & 0x1
    N >>= 1

    if v == 1 {
      max = Max(max, c)
      c = 0
    }

    if max == -1 {
      continue
    }

    c++
  }

  return max
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

// -- end --

