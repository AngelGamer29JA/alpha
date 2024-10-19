package alpha

import (
	"fmt"
)

func PrintError(cause error, content string, line int, script string) {
	fmt.Printf(
		"error: %s\n\tat %s:%d\ncause: %s\n",
		content, script, line, cause,
	)
	/*
		error:
			at
		cause:
	*/
}
