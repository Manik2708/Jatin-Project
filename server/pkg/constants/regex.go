package constants

import "regexp"

var (
	EMAIL_REGEX = regexp.MustCompile(`[A-Za-z0-9\._%+\-]+@[A-Za-z0-9\.\-]+\.[A-Za-z]{2,}`)
)