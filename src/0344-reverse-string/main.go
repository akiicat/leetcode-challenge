package main

// T: O(n)
// M: O(1)
// -- start --

// one index more faster
func reverseString(s []byte) {
  n := len(s)
  for i := 0; i < n / 2; i++ {
    s[i], s[n-1-i] = s[n-1-i], s[i]
  }
}

func reverseStringTwoIndex(s []byte) {
  l, r := 0, len(s) - 1

  for l < r {
    s[l], s[r] = s[r], s[l]
    l++
    r--
  }
}


// -- end --
