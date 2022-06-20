package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to waving-hands")
	
	def_bmp := default_bitmaps()

	fmt.Printf("DEBUG def_bmp.empty: %x\n", def_bmp.empty)
	fmt.Printf("DEBUG def_bmp.full: %x\n", def_bmp.full)

	fmt.Println("Thank you for playing waving-hands.")
}
