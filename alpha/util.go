package alpha

import (
	"os"
)

// Leer el contenido de un archivo
func ReadTextFile(path string) (string, error) {
	byteData, err := os.ReadFile(path)

	return string(byteData), err
}
