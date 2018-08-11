package main

import (
	"fmt"
)

func main() {
	const (
		sun = iota
		mon
		tue
		wed
		the
		fry
		sat
	)

	fmt.Printf("sun: %v, mon: %v, tue: %v, ...", sun, mon, tue)
}
