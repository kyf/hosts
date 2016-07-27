package hosts

import (
	"fmt"
	"strings"
)

func encode(item Host) (line string) {
	enable := "#"
	if item.Enabled {
		enable = ""
	}
	ip := item.Ip
	domain := strings.Join(item.Domains, " ")
	comment := item.Comment
	if len(comment) > 0 {
		comment = "#" + comment
	}

	return fmt.Sprintf("%s %s  %s  %s", enable, ip, domain, comment)
}
