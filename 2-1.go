package main

import "fmt"

func main() {

  var x1 uint32 = 88         // 01011000
  var x2 uint32 = 167        // 10100111
  var x3 uint32 = 3113827056 // 10111001100110010011101011110000

  fmt.Printf("%08b -> %08b\n", x1, offTheRightMostBit(x1))
  fmt.Printf("%08b -> %08b\n", x1, detachTheRightMost1Bit(x1))
  fmt.Printf("%08b -> %08b\n", x2, detachTheRightMostZeroBit(x2))
  fmt.Printf("%08b -> %08b\n", x1, maskTrailingZero1(x1))
  fmt.Printf("%08b -> %08b\n", x1, maskTrailingZero2(x1))
  fmt.Printf("%08b -> %08b\n", x1, maskTrailingZero3(x1))
  fmt.Printf("%08b -> %08b\n", x1, maskTrailingZeroInclude1(x1))
  fmt.Printf("%08b -> %08b\n", x1, offTheRightMostBits(x1))
  fmt.Printf("%08b -> %08b\n", x2, onTheRightMostBit(x2))
  fmt.Printf("%08b\n", testDiffPowerOf2())
  fmt.Printf("%32b -> %32b\n", x3, snoob(x3))
}

// x & (x - 1)
func offTheRightMostBit(x uint32) uint32 {
                       // x -> 10100111
  var y uint32 = x - 1 //      01010111
  var z uint32 = x & y //      01010000
  return z
}

// x & (-x)
func detachTheRightMost1Bit(x uint32) uint32 {
                       // x -> 10100111
  var y uint32 = -x    //      10101000
  var z uint32 = x & y //      00001000
  return z
}

// ^x & (x + 1)
func detachTheRightMostZeroBit(x uint32) uint32 {
                       // x -> 10100111
  var y uint32 = x + 1 //      10101000
  var z uint32 = ^x    //      01011000
  var a uint32 = y & z //      00001000
  return a
}

// ^x & (x - 1)
func maskTrailingZero1(x uint32) uint32 {
                       // x -> 10100111
  var y uint32 = ^x    //      10100111
  var z uint32 = x - 1 //      01010111
  var a uint32 = y & z //      00000111
  return a
}

// ^(x | -x)
func maskTrailingZero2(x uint32) uint32 {
                       // x -> 10100111
  var y uint32 = -x    //      10101000
  var z uint32 = x | y //      11111000
  var a uint32 = ^z    //      00000111
  return a
}

// (x & -x) - 1
func maskTrailingZero3(x uint32) uint32 {
                       // x -> 10100111
  var y uint32 = -x    //      10101000
  var z uint32 = x & y //      00001000
  var a uint32 = z - 1 //      00000111
  return a
}

// x ^ (x - 1)
func maskTrailingZeroInclude1(x uint32) uint32 {
                       // x -> 10100111
  var y uint32 = x - 1 //      01010111
  var z uint32 = x ^ y //      00001111
  return z
}

// ((x | (x - 1)) + 1) & x
func offTheRightMostBits(x uint32) uint32 {
                       // x -> 10100111
  var y uint32 = x - 1 //      01010111
  var z uint32 = x | y //      01011111
  var a uint32 = z + 1 //      01100000
  var b uint32 = a & x //      01000000
  return b
}

// x | (x + 1)
func onTheRightMostBit(x uint32) uint32 {
                       // x -> 10100111
  var y uint32 = x + 1 //      10101000
  var z uint32 = x | y //      10101111
  return z
}

func testDiffPowerOf2() uint32 {
  var x uint32 = 128   // 10000000
  var y uint32 = 32    // 00100000
  var z uint32 = x - y // 01100000
  return offTheRightMostBits(z)
}

func snoob(x uint32) uint32 {
  var smallest, ripple, ones uint32 // x -> xxx0 1111 0000
  smallest = x & -x                 //      0000 0001 0000
  ripple = x + smallest             //      xxx1 0000 0000
  ones = x ^ ripple                 //      0001 1111 0000
  ones = (ones >> 2) / smallest     //      0000 0000 0111
  return ripple | ones              //      xxx1 0000 0111
}
