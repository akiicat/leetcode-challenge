package main

// https://leetcode.com/articles/longest-common-prefix/
// T: O(n) n is the sum of all characters in all strings.
// M: O(1)
// -- start --

// BinarySearch
// T: O(n)
// M: O(1)
func longestCommonPrefix(strs []string) string {
  if len(strs) == 0 || len(strs[0]) == 0 {
    return ""
  }

  low, high := 1, len(strs[0])
  for _, str := range strs {
    if len(str) < high {
      high = len(str)
    }
  }

  for low <= high {
    mid := (low + high) / 2
    if isCommonPrefix(strs, mid) {
      low = mid + 1
    } else {
      high = mid - 1
    }
  }

  return strs[0][0:(low + high) / 2]
}

func isCommonPrefix(strs []string, l int) bool {
  str := strs[0][0:l]
  for _, s := range strs[1:] {
    for i := 0; i < l; i++ {
      if s[i] != str[i] {
        return false
      }
    }
  }
  return true
}

// T: O(n)
// M: O(1)
func longestCommonPrefixDivideAndConquer(strs []string) string {
  if len(strs) == 0 || len(strs[0]) == 0 {
    return ""
  }

  return DivideAndConquer(strs, 0, len(strs) - 1)
}

func DivideAndConquer(strs []string, l, r int) string {
  if l == r {
    return strs[l]
  }

  mid := (l + r) / 2

  lcpLeft := DivideAndConquer(strs, l, mid)
  lcpRight := DivideAndConquer(strs, mid + 1, r)

  count := len(lcpLeft)
  if count > len(lcpRight) {
    count = len(lcpRight)
  }

  for i := 0; i < count; i++ {
    if lcpLeft[i] != lcpRight[i] {
      return lcpLeft[0:i]
    }
  }

  return lcpLeft[0:count]
}

// T: O(n)
// M: O(1)
func longestCommonPrefixVerticalScanning(strs []string) string {
  if len(strs) == 0 {
    return ""
  }

  count, str := len(strs[0]), strs[0]

  for i := 0; i < count; i++ {
    for _, s := range strs[1:] {
      if i >= len(s) || s[i] != str[i] {
        return str[0:i]
      }
    }
  }

  return str[0:count]
}

// T: O(n)
// M: O(1)
func longestCommonPrefixHorizontalScanning(strs []string) string {
  if len(strs) == 0 {
    return ""
  }

  count, str := len(strs[0]), strs[0]

  for _, s := range strs[1:] {

    if count > len(s) {
      count = len(s)
    }

    for i, c := range []byte(s) {
      if i >= count {
        break
      }

      if c != str[i] {
        count = i
      }
    }
  }

  return str[0:count]
}

// -- end --

