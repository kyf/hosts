package hosts

import (
	"strings"
)

func StringSliceContains(list []string, need string) bool {
	for _, it := range list {
		if strings.EqualFold(need, it) {
			return true
		}
	}

	return false
}
