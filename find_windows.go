// +build windows
package hosts

import (
	"os"
	"path"
)

func findHost() string {
	return path.Join(os.Getenv("WINDIR"), "System32", "drivers", "etc", "hosts")
}
