package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  z := float64(1)
  i := 0
  for ; i < 10000; i++ {
    z = z - ((math.Pow(z, 2.0) - x) / 2 * z)
  }
  fmt.Println("Number of iterations:", i)
  return z
}

func SqrtStop(x float64) float64 {
  z := float64(1)
  y := math.Sqrt(x)
  i := 0
  for ; math.Abs(y - z) > 0.001; i++ {
    z = z - ((math.Pow(z, 2.0) - x) / 2 * z)
  }
  fmt.Println("Number of iterations:", i)
  return z
}

func main() {
  fmt.Println(Sqrt(2))
  fmt.Println("-----")
  fmt.Println(SqrtStop(2))
}
