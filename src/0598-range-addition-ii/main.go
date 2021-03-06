package main

// T: O(n) n for the number of operations
// M: O(1)
// -- start --

func maxCount(m int, n int, ops [][]int) int {
  for _, v := range ops {
    if v[0] < m {
      m = v[0]
    }
    if v[1] < n {
      n = v[1]
    }
  }
  return m * n
}

func maxCountMin(m int, n int, ops [][]int) int {
  for _, v := range ops {
    m = Min(m, v[0])
    n = Min(n, v[1])
  }

  return m * n
}

func Min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

// -- end --

