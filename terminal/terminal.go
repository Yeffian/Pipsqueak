package terminal

import "fmt"

func Clear() {
	fmt.Println("\x1b[2J")
}
