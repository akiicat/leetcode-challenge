package main

// T: O(n)
// M: O(1)
// -- start --

func addBinary(a string, b string) string {
  b2s := map[bool]string{false: "0", true: "1"}

  rtn := ""
  s, c := false, false // sum, carry
  for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
    s, c = adder(i >= 0 && a[i] == '1', j >= 0 && b[j] == '1', c)
    rtn = b2s[s] + rtn
  }

  if c {
    return "1" + rtn
  }

  return rtn
}

func adder(a, b, c bool) (bool, bool) {
  // sum: a ^ b ^ c
  // carry: (a & b) | (b & c) | (c & a)
  return ((a != b) != c), a && b || b && c || c && a
}

// -- end --

