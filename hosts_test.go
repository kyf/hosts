package hosts

import (
	"log"
	"testing"
)

func TestLoad(t *testing.T) {
	h, err := Load()
	if err != nil {
		t.Fatal(err)
	}

	for _, it := range h.Items {
		log.Print(*it)
	}

	host := h.Get("192.168.0.36", "www.baidu.com")
	if host != nil {
		log.Print(host)
		host.Disable()
	}

	err = h.Save()
	if err != nil {
		t.Fatal(err)
	}

}
