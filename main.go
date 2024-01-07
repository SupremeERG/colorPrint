package colorPrint

import "fmt"

func Reset() {
	colorReset := "\033[0m"

	fmt.Print(colorReset)
}

func ClearConsole(operating_system string) {

}

func Color(color string, line string) string {
	bold := "\033[1m"
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	colorPurple := "\033[35m"
	colorCyan := "\033[36m"
	colorWhite := "\033[37m"
	switch color {
	case "red":
		return bold + colorRed + line
	case "green":
		return bold + colorGreen + line

	case "yellow":
		return bold + colorYellow + line
	case "blue":
		return bold + colorBlue + line
	case "purple":
		return bold + colorPurple + line
	case "cyan":
		return bold + colorCyan + line
	case "white":
		return bold + colorWhite + line
	}
	return ""

}

func Default() string {
	return "\033[0m"
}
