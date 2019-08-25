package main
import "fmt"
import "math"

func main() {
  i, o := 2, true
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", judgeSquareSum(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = 3, false
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", judgeSquareSum(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = 5, true
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", judgeSquareSum(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = 13, true
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", judgeSquareSum(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = 2147483645, false
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", judgeSquareSum(i))
  fmt.Printf("Expect: %t\n", o)
}

// https://leetcode.com/articles/sum-of-square-numbers/
// T: O(sqrt(N) * log(N))
// M: O(1)
// -- start --

func judgeSquareSum(c int) bool {
  return judgeSquareSumFermatTheorem(c)
}

// T: O(sqrt(N) * log(N))
// M: O(1)
func judgeSquareSumSqrt(c int) bool {
  for a := 0; a * a <= c; a++ { // sqrt(N)
    b := math.Sqrt(float64(c - a * a)) // log(N)
    if b == float64(int(b)) {
      return true
    }
  }

  return false
}

// T: O(N)
// M: O(1)
func judgeSquareSumBetterBruteForce(c int) bool {
  for a := 0; a * a <= c; a++ {
    b := c - a * a
    if isPerfectSquare(b) {
      return true
    }
  }

  return false
}

// T: O(N)
// M: O(1)
func judgeSquareSumBruteForce(c int) bool {
  for a := 0; a * a <= c; a++ {
    for b := 0; b * b <= c; b++ {
      if a * a + b * b == c {
        return true
      }
    }
  }

  return false
}

// leetcode 367.
// T: O(sqrt(N))
// M: O(1)
func isPerfectSquare(num int) bool {
	i := 1
	for num > 0 {
		num -= i
		i += 2
	}
	return num == 0
}

// T: O(sqrt(N) * log(N))
// M: O(1)
func judgeSquareSumBinarySearch(c int) bool {
  for a := 0; a * a <= c; a++ { // sqrt(N)
    b := c - a * a
    if binarySearch(0, b, b) { // log(N)
      return true
    }
  }

  return false
}

// leetcode 69.
// T: O(log(N))
// M: O(1)
func binarySearch(start, end, num int) bool {
  if start > end {
    return false
  }

  mid := (start + end) / 2
  if mid * mid == num {
    return true
  }

  if mid * mid > num {
    return binarySearch(start, mid-1, num)
  }

  return binarySearch(mid + 1, end, num)
}

// T: O(sqrt(N) * log(N))
// M: O(1)
func judgeSquareSumFermatTheorem(c int) bool {
  for a := 2; a * a <= c; a++ {
    count := 0
    if c % a == 0 {
      for c % a == 0 {
        count++
        c /= a
      }
      if a % 4 == 3 && count % 2 != 0 {
        return false
      }
    }
  }

  return c % 4 != 3
}

// -- end --
