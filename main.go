package main

import (
    "fmt"
    "math"
)

func main() {
    var history []string
    for {
        var a, b float64
        var op string

        fmt.Print("Enter first number: ")
        _, err := fmt.Scanln(&a)
        if err != nil {
            fmt.Println("Invalid input for number!")
            continue
        }

        fmt.Print("Enter operator (+, -, *, /, ^): ")
        fmt.Scanln(&op)

        fmt.Print("Enter second number: ")
        _, err = fmt.Scanln(&b)
        if err != nil {
            fmt.Println("Invalid input for number!")
            continue
        }

        var result float64
        var valid bool = true
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
                valid = false
            }
        case "^":
            result = math.Pow(a, b)
        default:
            fmt.Println("Unknown operator.")
            valid = false
        }

        if valid {
            fmt.Println("Result:", result)
            history = append(history, fmt.Sprintf("%v %v %v = %v", a, op, b, result))
        }

        fmt.Print("Do you want to calculate again? (y/n): ")
        var again string
        fmt.Scanln(&again)
        if again != "y" && again != "Y" {
            break
        }
    }

    fmt.Println("\nCalculation history:")
    for _, entry := range history {
        fmt.Println(entry)
    }
}
