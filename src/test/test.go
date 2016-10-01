package main

import (
	"fetchsymbols"
	"fmt"
)

func main() {
	alic := fetchsymbols.GetSymbols()
	fmt.Println(alic)
}
