package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	expr := regexp.MustCompile(`\b\w`)
	texto := "antonio carlos jobim"

	transfomadora := func(s string) string {
		return strings.ToUpper(s)
	}

	fmt.Println(transfomadora(texto))
	fmt.Println(expr.ReplaceAllStringFunc(texto, transfomadora))
}
