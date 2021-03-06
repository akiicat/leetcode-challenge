package main
import "math"

// T: O(log(n))
// M: O(1)
// -- start --

func mySqrt(x int) int {
  l, r := 1, x

  for {
    mid := (l + r) / 2
    if mid == l {
      return mid
    }

    sqr := mid * mid
    if sqr == x {
      return mid
    }

    if sqr > x {
      r = mid
    }

    if sqr < x {
      l = mid
    }
  }
}

func mySqrtBinarySearch(x int) int {
  return binarySearch(1, x, x)
}

func binarySearch(start, end, num int) int {
  if start > end {
    return 0
  }

  mid := (start + end) / 2
  if mid * mid <= num && (mid + 1) * (mid + 1) > num {
    return mid
  }

  if mid * mid > num {
    return binarySearch(start, mid - 1, num)
  }

  return binarySearch(mid + 1, end, num)
}

func mySqrtMath(x int) int {
  return int(math.Sqrt(float64(x)))
}

// -- end --

