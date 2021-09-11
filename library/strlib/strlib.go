package strlib

import "fmt"

func Concat(a, b string) string {
	return fmt.Sprintf("%s%s", a, b)
}