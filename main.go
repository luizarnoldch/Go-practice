package main

import (
	"fmt"
	"strings"
)

func main() {

	ARN := "arn:aws:connect:us-west-2:356442333285:instance/7a3ba1cc-c5af-4c07-9f6e-7f1b0d651267"

	if strings.HasPrefix(ARN, "arn:aws") {
		fmt.Println("Hola mundo")
	}
}