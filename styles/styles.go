package styles

const (
	StyleReset     = "\x1b[0m"
	StyleBold      = "\x1b[1m"
	StyleDim       = "\x1b[2m"
	StyleItalic    = "\x1b[3m"
	StyleUnderline = "\x1b[4m"
)

func Stylize(style, text string) string {
	return style + text + StyleReset
}
