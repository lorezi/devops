// Package scan provides types and functions to perform TCP port scans on a list of hosts
package scan

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

var (
	ErrExists    = errors.New("host already in the list")
	ErrNotExists = errors.New("host not in the list")
)

type HostList struct {
	Hosts []string
}

func (h *HostList) search(host string) (bool, int) {
	sort.Strings(h.Hosts)
	i := sort.SearchStrings(h.Hosts, host)
	if i < len(h.Hosts) && h.Hosts[i] == host {
		return true, i
	}
	return false, -1
}

func (h *HostList) Add(host string) error {
	found, _ := h.search(host)

	if found {
		return fmt.Errorf("%w: %s", ErrExists, host)
	}

	h.Hosts = append(h.Hosts, host)
	return nil
}

// Remove deletes a host from the host
func (h *HostList) Remove(host string) error {
	found, i := h.search(host)
	if found {
		h.Hosts = append(h.Hosts[:i], h.Hosts[i+1:]...)
		return nil
	}
	return fmt.Errorf("%w: %s", ErrNotExists, host)
}

// Load obtains hosts from a hosts file
func (h *HostList) Load(hostFile string) error {

	f, err := os.Open(hostFile)
	if err != nil {
		switch {
		case errors.Is(err, os.ErrNotExist):
			return nil

		default:
			return err
		}
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		h.Hosts = append(h.Hosts, scanner.Text())
	}

	return nil

}

// Save saves hosts to a hosts file
func (h *HostList) Save(hostFile string) error {
	output := ""
	for _, v := range h.Hosts {
		output += fmt.Sprintln(v)
	}
	return ioutil.WriteFile(hostFile, []byte(output), 0644)
}
