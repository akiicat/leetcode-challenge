package main

// T: O(n)
// M: O(1)
// -- start --

func removePalindromeSub(s string) int {
  if len(s) == 0 {
    return 0
  }

  for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1{
    if s[l] != s[r] {
      return 2
    }
  }

  return 1
}

// -- end --

