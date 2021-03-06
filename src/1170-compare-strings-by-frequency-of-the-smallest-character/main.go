package main

// T: O(n + m)
// M: O(n)
// -- start --

func numSmallerByFrequency(queries []string, words []string) []int {
  m := [12]int{}
  for _, v := range words {
    m[f(v)]++
  }

  for i := len(m)-2; i >= 0; i-- {
    m[i] += m[i+1]
  }

  rtn := []int{}
  for _, v := range queries {
    rtn = append(rtn, m[f(v)+1])
  }

  return rtn
}

func f(s string) int {
  w, c := 'z', 0
  for _, v := range s {
    if v > w {
      continue
    }

    if v < w {
      w, c = v, 0
    }

    c++
  }
  return c
}

// -- end --

