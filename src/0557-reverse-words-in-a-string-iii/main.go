package main
import "fmt"

func main() {
  i, o := "Let's a ab take LeetCode contest", "s'teL a ba ekat edoCteeL tsetnoc"
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %s\n", reverseWords(i))
  fmt.Printf("Expect: %s\n", o)
}

// T: O(n) n is the length of the string
// M: O(1)
// -- start --

func reverseWords(s string) string {
  n := len(s)
  b := []byte(s)

  l, r := 0, 0
  for r < n {
    if b[r] == ' ' {
      Reverse(b, l, r-1)
      l, r = r+1, r+1
    }
    r++
  }
  Reverse(b, l, r-1)

  return string(b)
}

// leetcode 0344-reverse-string
func Reverse(b []byte, l, r int) {
  for r > l {
    b[l], b[r] = b[r], b[l]
    l, r = l+1, r-1
  }
  return
}

// -- end --

