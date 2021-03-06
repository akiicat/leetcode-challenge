package main
import "testing"
import . "main/pkg/testing_helper"

func TestLetterCasePermutation(_t *testing.T) {
  i, o := "a1b2", []string{"a1b2", "a1B2", "A1b2", "A1B2"}
  T(_t, S(i), S(Sort(letterCasePermutation(i))), S(Sort(o)))

  i, o = "3z4", []string{"3z4", "3Z4"}
  T(_t, S(i), S(Sort(letterCasePermutation(i))), S(Sort(o)))

  i, o = "12345", []string{"12345"}
  T(_t, S(i), S(Sort(letterCasePermutation(i))), S(Sort(o)))
}

func TestLetterCasePermutationBit(_t *testing.T) {
  i, o := "a1b2", []string{"a1b2", "a1B2", "A1b2", "A1B2"}
  T(_t, S(i), S(Sort(letterCasePermutationBit(i))), S(Sort(o)))

  i, o = "3z4", []string{"3z4", "3Z4"}
  T(_t, S(i), S(Sort(letterCasePermutationBit(i))), S(Sort(o)))

  i, o = "12345", []string{"12345"}
  T(_t, S(i), S(Sort(letterCasePermutationBit(i))), S(Sort(o)))
}

