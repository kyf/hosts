package hosts

import (
	"bufio"
	"os"
	"strings"
	"sync"
)

type Host struct {
	Ip      string
	Domains []string
	Comment string
	Enabled bool
}

type Hosts struct {
	Items    []*Host
	hostPath string
	sync.Mutex
}

func Load() (*Hosts, error) {
	h := &Hosts{}
	h.hostPath = findHost()

	fp, err := os.Open(h.hostPath)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		ip, domains, comment, enable, err := decode(line[:len(line)-1])
		if err != nil {
			continue
		}
		for _, domain := range domains {
			h.Set(ip, domain, comment, enable)
		}
	}

	return h, nil
}

func (h *Hosts) Set(ip, domain, comment string, enable bool) (host *Host) {
	_hosts := h.GetByIp(ip)
	for _, it := range _hosts {
		if it.Enabled == enable {
			if !StringSliceContains(it.Domains, domain) {
				it.Domains = append(it.Domains, domain)
			}
			host = it
			return
		}
	}

	host = &Host{Ip: ip, Domains: []string{domain}, Enabled: enable, Comment: comment}
	h.Lock()
	defer h.Unlock()
	h.Items = append(h.Items, host)
	return host
}

func (host *Host) Enable() {
	host.Enabled = true
}

func (host *Host) Disable() {
	host.Enabled = false
}

func (h *Hosts) GetByIp(ip string) []*Host {
	result := make([]*Host, 0)
	for _, it := range h.Items {
		if strings.EqualFold(ip, it.Ip) {
			result = append(result, it)
		}
	}

	return result
}

func (h *Hosts) Get(ip, domain string) (host *Host) {
	for _, it := range h.Items {
		if strings.EqualFold(ip, it.Ip) &&
			StringSliceContains(it.Domains, domain) {
			host = it
		}
	}
	return
}

func (h *Hosts) Save() error {
	fp, err := os.OpenFile(h.hostPath, os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer fp.Close()
	for _, item := range h.Items {
		line := encode(*item)
		_, err = fp.Write([]byte(line + CMD_WRAP))
		if err != nil {
			return err
		}
	}

	return nil
}
