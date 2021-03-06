package main
import "strings"

// T: O(n)
// M: O(1)
// -- start --

func detectCapitalUse(word string) bool {
  n := len(word)
  if n <= 1 {
    return true
  }

  if (word[0] >= 'a') && (word[1] < 'a') {
    return false
  }

  for i := 1; i < n - 1; i++ {
    if (word[i] >= 'a') != (word[i+1] >= 'a') {
      return false
    }
  }

  return true
}

func detectCapitalUseString(word string) bool {
	return strings.ToUpper(word) == word || strings.ToLower(word[1:]) == word[1:]
}

// -- end --

