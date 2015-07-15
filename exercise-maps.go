package main

import (
  "strings"
  "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
  m := make(map[string]int)
  words := strings.Fields(s)
  for _, value := range words {
    v, ok := m[value]
    if ok {
      m[value] = v + 1
    } else {
      m[value] = 1
    }
  }
  return m
}

func main() {
  wc.Test(WordCount)
}
