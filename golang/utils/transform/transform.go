package transform

import "strings"

func ReplaceAll(old string, new string) func(string) string {
	return func(str string) string {
		return strings.ReplaceAll(str, old, new)
	}
}
