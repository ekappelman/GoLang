package main

import (
	"gofin/fetchsymbols"
	"fmt"
)

func main() {
	alic := fetchsymbols.GetSymbols()
	fmt.Println(alic)
}
