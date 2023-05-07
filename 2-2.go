package main

import "fmt"

func main() {

  var x uint32 = 874986927
  var y uint32 = 384732832
  var r uint32

  r = a(x)
  fmt.Printf("%032b(%d) -> %032b(%d)\n", x, x, r, r)

  r = b(x)
  fmt.Printf("%032b(%d) -> %032b(%d)\n", x, x, r, r)

  r = c(x)
  fmt.Printf("%032b(%d) -> %032b(%d)\n", x, x, r, r)

  r = d(x)
  fmt.Printf("%032b(%d) -> %032b(%d)\n", x, x, r, r)

  r = e(x)
  fmt.Printf("%032b(%d) -> %032b(%d)\n", x, x, r, r)

  r = f(x, y)
  fmt.Printf("%032b(%d) + %032b(%d) = %032b(%d)\n", x, x, y, y, r, r)

  r = g(x, y)
  fmt.Printf("%032b(%d) + %032b(%d) = %032b(%d)\n", x, x, y, y, r, r)

  r = h(x, y)
  fmt.Printf("%032b(%d) + %032b(%d) = %032b(%d)\n", x, x, y, y, r, r)

  r = i(x, y)
  fmt.Printf("%032b(%d) + %032b(%d) = %032b(%d)\n", x, x, y, y, r, r)
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
