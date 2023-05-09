package main

import (
  "fmt"
  "math"
)

func main() {

  fmt.Printf("(x | y) >= max(x, y)\n")

  a(100, 55)
  a(12, 3)
  a(119, 119)

  fmt.Printf("\n(x & y) <= min(x, y)\n")

  b(100, 55)
  b(12, 3)
  b(119, 119)

  fmt.Printf("\n(x | y) <= x + y\n")

  c(100, 55)
  c(12, 3)
  c(119, 119)

  fmt.Printf("\n|x - y| <= x ^ y\n")

  e(55, 100)
  e(3, 12)
  e(119, 119)
}

func a(x uint32, y uint32) {
  var result1 uint32 = x | y
  var result2 uint32 = uint32(math.Max(float64(x), float64(y)))
  fmt.Printf("%08b(%d) | %08b(%d) = %08b(%d)\n", x, x, y, y, result1, result1)
  fmt.Printf("max(%08b(%d), %08b(%d)) = %08b(%d)\n", x, x, y, y, result2, result2)
  fmt.Printf("%08b(%d) >= %08b(%d)\n", result1, result1, result2, result2)
}

func b(x uint32, y uint32) {
  var result1 uint32 = x & y
  var result2 uint32 = uint32(math.Min(float64(x), float64(y)))
  fmt.Printf("%08b(%d) & %08b(%d) = %08b(%d)\n", x, x, y, y, result1, result1)
  fmt.Printf("min(%08b(%d), %08b(%d)) = %08b(%d)\n", x, x, y, y, result2, result2)
  fmt.Printf("%08b(%d) <= %08b(%d)\n", result1, result1, result2, result2)
}

func c(x uint32, y uint32) {
  var result1 uint32 = x | y
  var result2 uint32 = x + y
  fmt.Printf("%08b(%d) | %08b(%d) = %08b(%d)\n", x, x, y, y, result1, result1)
  fmt.Printf("%08b(%d) + %08b(%d) = %08b(%d)\n", x, x, y, y, result2, result2)
  fmt.Printf("%08b(%d) <= %08b(%d)\n", result1, result1, result2, result2)
}

// |x - y| = max(x, y) - min(x, y) <= (x | y) - (x & y) = x ^ y
func e(x uint32, y uint32) {
  var result1 uint32 = uint32(math.Abs(float64(x) - float64(y)))
  var result2 uint32 = x ^ y
  fmt.Printf("|%08b(%d) - %08b(%d)| = %08b(%d)\n", x, x, y, y, result1, result1)
  fmt.Printf("%08b(%d) ^ %08b(%d) = %08b(%d)\n", x, x, y, y, result2, result2)
  fmt.Printf("%08b(%d) <= %08b(%d)\n", result1, result1, result2, result2)
}