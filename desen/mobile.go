package desen

import (
	"github.com/fusugongzi/desen"
	"github.com/fusugongzi/desen/reco"
	"github.com/fusugongzi/gotools/slices"
	"strings"
)

func MobileRandom(mobile string) string {
	if !reco.IsMobile(mobile) {
		return mobile
	}
	builder := strings.Builder{}
	builder.WriteString(slices.Random(desen.AllMobilePrefix))
	builder.WriteString(string(slices.Randoms(desen.Digits, 8)))
	return builder.String()
}
