package prompts

import (
	"bufio"
	"fmt"
	"io"

	. "goterm/colors"
	. "goterm/styles"
)

func Prompt(prompt string, in io.Reader) (string, error) {
	reader := bufio.NewReader(in)
	fmt.Print(prompt)
	text, err := reader.ReadString('\n')

	return text, err
}

func PromptColored(prompt, color string, in io.Reader) (string, error) {
	reader := bufio.NewReader(in)
	fmt.Println(ColorizeFore(color, prompt))
	text, err := reader.ReadString('\n')

	return text, err
}

func PromptColoredAndStyled(prompt, color, style string, in io.Reader) (string, error) {
	reader := bufio.NewReader(in)
	fmt.Println(Stylize(style, ColorizeFore(color, prompt)))
	text, err := reader.ReadString('\n')

	return text, err
}
