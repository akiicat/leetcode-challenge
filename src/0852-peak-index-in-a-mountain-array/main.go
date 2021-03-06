package main

// T: O(log(n))
// M: O(1)
// -- start --

// T: O(n)
// M: O(1)
func peakIndexInMountainArray(A []int) int {
  for i, v := range A {
    if v > A[i+1] {
      return i
    }
  }
  return -1
}

// T: O(log(n))
// M: O(1)
func peakIndexInMountainArrayBinarySearch(A []int) int {
    l, r := 0, len(A)-1
    for l < r {
        m := l + (r-l)/2
        if A[m] < A[m+1] {
            l = m + 1
        } else {
            r = m
        }
    }
    return l
}

// T: O(log(n))
// M: O(1)
func peakIndexInMountainArrayBinarySearchRecursive(A []int) int {
  return R(A, 0, len(A)-1)
}

func R(A []int, lo, hi int) int {
  if lo >= hi {
    return -1
  }

  mi := (lo + hi) / 2

  max := Max(R(A, lo, mi), R(A, mi+1, hi))

  if A[mi] < A[mi + 1] {
    return Max(max, mi + 1)
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

