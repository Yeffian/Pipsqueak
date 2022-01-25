package terminal

import "fmt"

func Clear() {
	// this doesnt actually properly clear the terminal, more just scrolls down so you cant see anything anymore, you can break the illusion by
	// just scrolling up, but its fine for now
	fmt.Println("\033[H\033[2J")
}
