package colors

const (
	// Foreground Colors
	ForeColorDefault      = "\x1b[39m"
	ForeColorRed          = "\x1b[31m"
	ForeColorBlack        = "\x1b[30m"
	ForeColorGreen        = "\x1b[32m"
	ForeColorYellow       = "\x1b[33m"
	ForeColorBlue         = "\x1b[34m"
	ForeColorMagenta      = "\x1b[35m"
	ForeColorCyan         = "\x1b[36m"
	ForeColorWhite        = "\x1b[37m"
	ForeColorLightGrey    = "\x1b[90m"
	ForeColorLightRed     = "\x1b[91m"
	ForeColorLightGreen   = "\x1b[92m"
	ForeColorLightYellow  = "\x1b[93m"
	ForeColorLightBlue    = "\x1b[94m"
	ForeColorLightMagenta = "\x1b[95m"
	ForeColorLightCyan    = "\x1b[96m"
	ForeColorLightWhite   = "\x1b[97m"

	// Background Colors
	BackColorDefault      = "\x1b[49m"
	BackColorBlack        = "\x1b[40m"
	BackColorRed          = "\x1b[41m"
	BackColorGreen        = "\x1b[42m"
	BackColorYellow       = "\x1b[43m"
	BackColorBlue         = "\x1b[44m"
	BackColorMagenta      = "\x1b[45m"
	BackColorCyan         = "\x1b[46m"
	BackColorWhite        = "\x1b[47m"
	BackColorLightGrey    = "\x1b[100m"
	BackColorLightRed     = "\x1b[101m"
	BackColorLightGreen   = "\x1b[102m"
	BackColorLightYellow  = "\x1b[103m"
	BackColorLightBlue    = "\x1b[104m"
	BackColorLightMagenta = "\x1b[105m"
	BackColorLightCyan    = "\x1b[106m"
	BackColorLightWhite   = "\x1b[107m"
)

func ColorizeFore(style, text string) string {
	return style + text + ForeColorDefault
}

func ColorizeBack(style, text string) string {
	return style + text + BackColorDefault
}
