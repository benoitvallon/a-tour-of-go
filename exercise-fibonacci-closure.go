package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
  prevNum := 0
  prevPrevNum := 0
  return func() int {
    num := prevPrevNum + prevNum
    if prevPrevNum == 0 {
      prevNum = 1
    }
    prevPrevNum = prevNum
    prevNum = num
    return num
  }
}

func main() {
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(f())
  }
}
