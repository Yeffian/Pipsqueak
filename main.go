package main

import (
	. "goterm/colors"
	. "goterm/prompts"
	. "goterm/styles"
	"os"
)

func main() {
	println(Stylize(StyleItalic, ColorizeFore(ForeColorCyan, "================== Styles ===================")))
	println(Stylize(StyleBold, "Hello, World!"))
	println(Stylize(StyleDim, "Hello, World!"))
	println(Stylize(StyleItalic, "Hello, World!"))
	println(Stylize(StyleUnderline, "Hello, World!"))
	println(Stylize(StyleItalic, ColorizeFore(ForeColorCyan, "==================================================")))

	println("\n")

	println(Stylize(StyleItalic, ColorizeFore(ForeColorCyan, "================== Foreground Colours ===================")))
	println(ColorizeFore(ForeColorRed, "Hello, World!"))
	println(ColorizeFore(ForeColorGreen, "Hello, World!"))
	println(ColorizeFore(ForeColorYellow, "Hello, World!"))
	println(ColorizeFore(ForeColorBlue, "Hello, World!"))
	println(ColorizeFore(ForeColorMagenta, "Hello, World!"))
	println(ColorizeFore(ForeColorCyan, "Hello, World!"))
	println(ColorizeFore(ForeColorLightGrey, "Hello, World!"))
	println(ColorizeFore(ForeColorLightRed, "Hello, World!"))
	println(ColorizeFore(ForeColorLightGreen, "Hello, World!"))
	println(ColorizeFore(ForeColorLightYellow, "Hello, World!"))
	println(ColorizeFore(ForeColorLightBlue, "Hello, World!"))
	println(ColorizeFore(ForeColorLightMagenta, "Hello, World!"))
	println(ColorizeFore(ForeColorLightCyan, "Hello, World!"))
	println(ColorizeFore(ForeColorLightWhite, "Hello, World!"))
	println(ColorizeFore(ForeColorWhite, "Hello, World!"))
	println(Stylize(StyleItalic, ColorizeFore(ForeColorCyan, "==========================================================")))

	println(Stylize(StyleItalic, ColorizeFore(ForeColorCyan, "================== Background Colours ===================")))
	println(ColorizeBack(BackColorRed, "Hello, World!"))
	println(ColorizeBack(BackColorGreen, "Hello, World!"))
	println(ColorizeBack(BackColorYellow, "Hello, World!"))
	println(ColorizeBack(BackColorBlue, "Hello, World!"))
	println(ColorizeBack(BackColorMagenta, "Hello, World!"))
	println(ColorizeBack(BackColorCyan, "Hello, World!"))
	println(ColorizeBack(BackColorLightGrey, "Hello, World!"))
	println(ColorizeBack(BackColorLightRed, "Hello, World!"))
	println(ColorizeBack(BackColorLightGreen, "Hello, World!"))
	println(ColorizeBack(BackColorLightYellow, "Hello, World!"))
	println(ColorizeBack(BackColorLightBlue, "Hello, World!"))
	println(ColorizeBack(BackColorLightMagenta, "Hello, World!"))
	println(ColorizeBack(BackColorLightCyan, "Hello, World!"))
	println(ColorizeBack(BackColorLightWhite, "Hello, World!"))
	println(ColorizeBack(BackColorWhite, "Hello, World!"))
	println(Stylize(StyleItalic, ColorizeFore(ForeColorCyan, "==========================================================")))

	txt, err := PromptColored("> ", ForeColorCyan, os.Stdin)

	if err != nil {
		panic(err)
	}

	println(txt)
}
