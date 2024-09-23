package reco

import (
	"github.com/fusugongzi/desen"
	"strings"
)

// IsMobile determine the string is a 11 digit phone number
func IsMobile(s string) bool {
	if len(s) != 11 {
		return false
	}

	runes := []rune(s)
	for _, r := range runes {
		if r < '0' || r > '9' {
			return false
		}
	}

	hasPrefix := false
	for _, v := range desen.AllMobilePrefix {
		if strings.HasPrefix(s, v) {
			hasPrefix = true
			break
		}
	}
	if !hasPrefix {
		return false
	}

	return true
}
