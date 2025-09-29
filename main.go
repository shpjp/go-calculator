package main

import (
    "fmt"
)

func main() {
    var a, b float64
var op string

fmt.Print("Enter first number: ")
fmt.Scanln(&a)
fmt.Print("Enter operator (+, -, *, /): ")
fmt.Scanln(&op)
fmt.Print("Enter second number: ")
fmt.Scanln(&b)

  var result float64
switch op {
case "+":
    result = a + b
case "-":
    result = a - b
case "*":
    result = a * b
case "/":
    if b != 0 {
        result = a / b
    } else {
        fmt.Println("Error: Division by zero!")
        return
    }
default:
    fmt.Println("Unknown operator.")
    return
}
fmt.Println("Result:", result)

  
}
