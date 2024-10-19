package color

import (
	"strings"
)

var (
	Reset   = Color{Name: "Reset", Value: "\033[0m", Code: "r"}
	Red     = Color{Name: "Red", Value: "\033[31m", Code: "4"}
	Green   = Color{Name: "Green", Value: "\033[32m", Code: "2"}
	Yellow  = Color{Name: "Yellow", Value: "\033[33m", Code: "e"}
	Blue    = Color{Name: "Blue", Value: "\033[34m", Code: "b"}
	Magenta = Color{Name: "Magenta", Value: "\033[35m", Code: "5"}
	Cyan    = Color{Name: "Cyan", Value: "\033[36m", Code: "3"}
	Gray    = Color{Name: "Gray", Value: "\033[37m", Code: "8"}
	White   = Color{Name: "White", Value: "\033[97m", Code: "f"}
)

type Color struct {
	Name  string
	Value string
	Code  string
}

func (c *Color) String() string {
	return c.Value
}

// Obtener el color mediante el nombre o codigo
//
// Ejemplo puede usarse "Red" o "4" en este caso
// el codigo del color rojo es "4"
func GetColor(input string) Color {
	colors := map[string]Color{
		Reset.Name:   Reset,
		Red.Name:     Red,
		Green.Name:   Green,
		Yellow.Name:  Yellow,
		Blue.Name:    Blue,
		Magenta.Name: Magenta,
		Cyan.Name:    Cyan,
		Gray.Name:    Gray,
		White.Name:   White,
	}

	for _, color := range colors {
		if color.Name == input || color.Code == input {
			return color
		}
	}

	return Reset
}

func Print(content string) {

}

func Formatter(content string) string {
	var formattedText string
	for i := 0; i < len(content); i++ {
		if content[i] == '&' && strings.Contains("0123456789abcdeo", string(content[i+1])) {
			/*
				"&8" <- remove &
				"8" <- replace with color
				"<color>..."
			*/

			runes := []rune(content)
			new_content := append(runes[0:i], runes[i+2:]...)

			formattedText = string(new_content)
		}
	}
	return formattedText
}
