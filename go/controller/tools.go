package controller

import (
	"crypto/md5" //nolint:gosec
	"fmt"
	"strings"
)

const gravatarURLTpl = "https://www.gravatar.com/avatar/%x?d=%s&r=%s"

func GravatarURLFunc(enabled bool, def string, rating string) func(string) string {
	if !enabled {
		return func(email string) string { return "" }
	}

	return func(email string) string {
		emailHash := md5.Sum([]byte(strings.ToLower(email))) //nolint:gosec
		return fmt.Sprintf("https://www.gravatar.com/avatar/%x?d=%s&r=%s", emailHash, def, rating)
	}
}
