package main

import "fmt"

func main() {

    // #(IF) - #(IF - ELSE) - #(IF - ELSE IF - ELSE)
    condition := true

    if condition {
        // code to execute if the condition is true
        fmt.Println("Only true")
    }

    condition = false

    if condition {
        // code to execute if the condition is true
    } else {
        // code to execute if the condition is false
        fmt.Println("Only else with false after 1 condition")
    }

    condition1 := false
    condition2 := false

    if condition1 {
        // code to execute if condition1 is true
    } else if condition2 {
        // code to execute if condition2 is true
    } else {
        // code to execute if none of the previous conditions are true
        fmt.Println("Only else with false after 2 conditions")
    }
    
    y := 11
    if a := 10; a < y {
        // code to execute if a < y
        fmt.Println("Local variable inside if clause")
    }

    // #SWITCH

    // switch with expression
    expression := 1
    
    switch expression {
    case 1:
        // code to execute if expression == 1
        fmt.Println("expression == 1")
    case 2:
        // code to execute if expression == 2
        fmt.Println("expression == 2")
    default:
        // code to execute if no matches are found
        fmt.Println("No match option")
    }

    // switch without expression
    condition1 = true
    condition2 = false

    switch {
    case condition1:
        // code to execute if condition1 is true
        fmt.Println("condition1 == true")
    case condition2:
        // code to execute if condition2 is true
        fmt.Println("condition2 == true")
    default:
        // code to execute if none of the previous conditions are true
        fmt.Println("No match option")
    }

    // switch with local expression
    y = 10
    switch a := 10; {
    case a < y:
        // code to execute if a < y
        fmt.Println("a < y")
    case a == y:
        // code to execute if a == y
        fmt.Println("a == y")
    default:
        // code to execute if none of the previous conditions are true
        fmt.Println("No match option")
    }
    
    // switch with fallthrough
    expression = 1
    value1 := 1
    value2 := 2

    switch expression {
    case value1:
        // code to execute if expression == value1
        fmt.Println("expression == value1")
        fallthrough
    case value2:
        // code to execute even if expression != value2 due to fallthrough
        fmt.Println("expression != value1 and fallthrough")
    default:
        // code to execute if no matches are found
        fmt.Println("No match option")
    }
    
    // #FOR
    // Print numbers from 0 to 4
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // Print numbers from 0 to 4 using a loop that resembles a 'while'
    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }

    // An infinite loop that prints "Hello" continuously (manually stop it)
    k := 0
    for {
        fmt.Println("Hello")
        k++
        if k == 3 {
            break
        }
    }

    numbers := []int{2, 4, 6, 8, 10}
    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

    colors := map[string]string{"red": "#ff0000", "green": "#00ff00", "blue": "#0000ff"}
    for key, value := range colors {
        fmt.Printf("Key: %s, Value: %s\n", key, value)
    }

    greeting := "Hello"
    for index, character := range greeting {
        fmt.Printf("Index: %d, Character: %c\n", index, character)
    }
}
