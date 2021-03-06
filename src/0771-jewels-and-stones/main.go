package main

// T: O(n + m)
// M: O(n)
// -- start --

func numJewelsInStones(J string, S string) int {
  j := Count(J)
  s := Count(S)

  c := 0
  for i, v := range j {
    if v != 0 {
      c += s[i]
    }
  }
  return c
}

func Count(s string) [52]int {
  m := [52]int{}
  for _, v := range s {
    if 'a' <= v && v <= 'z' {
      m[int(v-'a')]++
    } else if 'A' <= v && v <= 'Z' {
      m[int(v-'A')+26]++
    }
  }
  return m
}

// -- end --

