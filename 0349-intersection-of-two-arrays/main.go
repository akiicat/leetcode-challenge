package main
import "fmt"

func main() {
  i, n, o_n := []int{3,2,2,3}, 3, 2
  fmt.Printf("Input:  %v, %d\n", i, n)
  n = removeElement(i, n)
  fmt.Printf("Output: %v, %d\n", i[:n], n)
  fmt.Printf("Expect: %d\n", o_n)

  i, n, o_n = []int{0,1,2,2,3,0,4,2}, 2, 5
  fmt.Printf("Input:  %v, %d\n", i, n)
  n = removeElement(i, n)
  fmt.Printf("Output: %v, %d\n", i[:n], n)
  fmt.Printf("Expect: %d\n", o_n)
}

// T: O(N)
// M: O(1)
// -- start --

func removeElement(nums []int, val int) int {
  cur := 0

  for i, num := range nums {
    if num != val {
      nums[cur] = nums[i]
      cur++
    }
  }

  return cur
}

// -- end --

