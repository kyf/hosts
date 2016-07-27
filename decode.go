package hosts

import (
	"errors"
	"strings"
)

func decode(line string) (ip string, domains []string, comment string, enable bool, err error) {
	if len(line) == 0 {
		err = errors.New("line is empty")
		return
	}

	if line[0] == '#' {
		line = line[1:]
		enable = false
	} else {
		enable = true
	}

	if len(line) == 0 {
		err = errors.New("line is invalid")
		return
	}

	parts := strings.Split(line, "#")
	if len(parts) > 2 {
		comment = strings.Join(parts[1:], "#")
	} else if len(parts) > 1 {
		comment = parts[1]
	}
	body := parts[0]
	if len(body) == 0 {
		err = errors.New("line is invalid")
		return
	}

	bodies := strings.Split(body, " ")
	if len(bodies) < 2 {
		err = errors.New("line is invalid")
		return
	}
	for _, b := range bodies {
		b = strings.Trim(b, " ")
		if len(b) == 0 {
			continue
		}
		if len(ip) == 0 {
			ip = b
		} else {
			domains = append(domains, b)
		}
	}

	if len(ip) == 0 || len(domains) == 0 {
		err = errors.New("line is invalid")
		return
	}
	return
}
