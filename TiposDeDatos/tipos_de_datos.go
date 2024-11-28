package main

import "fmt"

// Booleans:
//          bool
// Numeric Integers:
//          (int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64)
// Numeric Floats:
//          (float32, float64)
// Complex Numbers:
//          (complex64, complex128)
// Text Strings:
//          string

func main() {
    var myBool bool = true
    fmt.Printf("Type: %T, Value: %v\n", myBool, myBool)

    var myInt int = 42
    var myFloat float64 = 42.42
    fmt.Printf("Type: %T, Value: %v\n", myInt, myInt)
    fmt.Printf("Type: %T, Value: %v\n", myFloat, myFloat)

    var myComplex complex128 = complex(2, 3)
    fmt.Printf("Type: %T, Value: %v\n", myComplex, myComplex)

    var myString string = "Hello World"
    fmt.Printf("Type: %T, Value: %v\n", myString, myString)
}
