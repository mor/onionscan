package utils

import (
        "log"
	"regexp"
	"strings"
)

func IsOnion(identifier string) bool {
	// TODO: At some point we will want to support i2p
	log.Printf("parse .onion string: %s", identifier)
	if len(identifier) >= 22 && strings.HasSuffix(identifier, ".onion") {
		matched, _ := regexp.MatchString(`(^|\.)[a-z2-7]{56}\.onion$`, identifier)
		return matched
	}
	return false
}
