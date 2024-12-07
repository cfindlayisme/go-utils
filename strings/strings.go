package strings

import "strings"

func StripNewlines(message string) string {
	message = strings.ReplaceAll(message, "\n", "")
	message = strings.ReplaceAll(message, "\r", "")

	return message
}
