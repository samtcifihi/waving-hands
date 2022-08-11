package main

import (
	"fmt"
	"github.com/samtcifihi/waving-hands/src/gamestate"
)

func main() {
	fmt.Println("Welcome to waving-hands")

	test_gamestate := gamestate.New_Game_State() // DEBUG
	fmt.Println("DEBUG test_gamestate: ", test_gamestate)

	fmt.Println("Thank you for playing waving-hands.")
}
