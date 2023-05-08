package main

import "fmt"

func main() {

  var x uint32 = 100
  var y uint32 = 55
  var result uint32

  result = a(x)
  fmt.Printf("%032b(%d) -> %032b(%d)\n", x, x, result, result)

  result = b(x)
  fmt.Printf("%032b(%d) -> %032b(%d)\n", x, x, result, result)

  result = c(x)
  fmt.Printf("%032b(%d) -> %032b(%d)\n", x, x, result, result)

  result = d(x)
  fmt.Printf("%032b(%d) -> %032b(%d)\n", x, x, result, result)

  result = e(x)
  fmt.Printf("%032b(%d) -> %032b(%d)\n", x, x, result, result)

  result = f(x, y)
  fmt.Printf("%032b(%d) + %032b(%d) = %032b(%d)\n", x, x, y, y, result, result)

  result = g(x, y)
  fmt.Printf("%032b(%d) + %032b(%d) = %032b(%d)\n", x, x, y, y, result, result)

  result = h(x, y)
  fmt.Printf("%032b(%d) + %032b(%d) = %032b(%d)\n", x, x, y, y, result, result)

  result = i(x, y)
  fmt.Printf("%032b(%d) + %032b(%d) = %032b(%d)\n", x, x, y, y, result, result)

  result = j(x, y)
  fmt.Printf("%032b(%d) - %032b(%d) = %032b(%d)\n", x, x, y, y, result, result)

  result = k(x, y)
  fmt.Printf("%032b(%d) - %032b(%d) = %032b(%d)\n", x, x, y, y, result, result)

  result = l(x, y)
  fmt.Printf("%032b(%d) - %032b(%d) = %032b(%d)\n", x, x, y, y, result, result)

  result = m(x, y)
  fmt.Printf("%032b(%d) - %032b(%d) = %032b(%d)\n", x, x, y, y, result, result)

  result = n(x, y)
  fmt.Printf("%032b(%d) ^ %032b(%d) = %032b(%d)\n", x, x, y, y, result, result)

  result = o(x, y)
  fmt.Printf("%032b(%d) & ^%032b(%d) = %032b(%d)\n", x, x, y, y, result, result)

  result = p(x, y)
  fmt.Printf("%032b(%d) & ^%032b(%d) = %032b(%d)\n", x, x, y, y, result, result)

  result = q(x, y)
  fmt.Printf("^(%032b(%d) - %032b(%d)) = %032b(%d)\n", x, x, y, y, result, result)

  result = r(x, y)
  fmt.Printf("^(%032b(%d) - %032b(%d)) = %032b(%d)\n", x, x, y, y, result, result)

  result = s(x, y)
  fmt.Printf("^(%032b(%d) ^ %032b(%d)) = %032b(%d)\n", x, x, y, y, result, result)

  result = t(x, y)
  fmt.Printf("^(%032b(%d) ^ %032b(%d)) = %032b(%d)\n", x, x, y, y, result, result)

  result = u(x, y)
  fmt.Printf("%032b(%d) | %032b(%d) = %032b(%d)\n", x, x, y, y, result, result)

  result = v(x, y)
  fmt.Printf("%032b(%d) & %032b(%d) = %032b(%d)\n", x, x, y, y, result, result)
}

// -x
func a(x uint32) uint32 {
  return ^x + 1
}

// -x
func b(x uint32) uint32 {
  return ^(x - 1)
}

// ^x
func c(x uint32) uint32 {
  return -x - 1
}

// -^x
func d(x uint32) uint32 {
  return x + 1
}

// ^-x
func e(x uint32) uint32 {
  return x - 1
}

// x + y
func f(x uint32, y uint32) uint32 {
  return x - ^y - 1
}

// x + y
func g(x uint32, y uint32) uint32 {
  return (x ^ y) + 2 * (x & y)
}

// x + y
func h(x uint32, y uint32) uint32 {
  return (x | y) + (x & y)
}

// x + y
func i(x uint32, y uint32) uint32 {
  return 2 * (x | y) - (x ^ y)
}

// x - y
func j(x uint32, y uint32) uint32 {
  return x + ^y + 1
}

// x - y
func k(x uint32, y uint32) uint32 {
  return (x ^ y) - 2 * (^x & y)
}

// x - y
func l(x uint32, y uint32) uint32 {
  return (x & ^y) - (^x & y)
}

// x - y
func m(x uint32, y uint32) uint32 {
  return 2 * (x & ^y) - (x ^ y)
}

// x ^ y
func n(x uint32, y uint32) uint32 {
  return (x | y) - (x & y)
}

// x & ^y
func o(x uint32, y uint32) uint32 {
  return (x | y) - y
}

// x & ^y
func p(x uint32, y uint32) uint32 {
  return (x | y) - y
}

// ^(x - y)
func q(x uint32, y uint32) uint32 {
  return y - x - 1
}

// ^(x - y)
func r(x uint32, y uint32) uint32 {
  return ^x + y
}

// ^(x ^ y)
func s(x uint32, y uint32) uint32 {
  return (x & y) - (x | y) - 1
}

// ^(x ^ y)
func t(x uint32, y uint32) uint32 {
  return (x & y) + ^(x | y)
}

// x | y
func u(x uint32, y uint32) uint32 {
  return (x & ^y) + y
}

// x & y
func v(x uint32, y uint32) uint32 {
  return (^x | y) - ^x
}
