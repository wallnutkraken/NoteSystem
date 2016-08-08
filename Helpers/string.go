package Helpers

import (
	"strings"
)

func PathBuilder(path ...string) string {
	return strings.Join(path, "/")
}
