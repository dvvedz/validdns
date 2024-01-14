package validdnsservers

import (
	"io"
	"log"
	"net/http"
	"strings"
)

type dns []byte

func GetPublicDnsServers() dns {
	resp, httpErr := http.Get("https://public-dns.info/nameservers.txt")
	if httpErr != nil {
		log.Fatalf("http get error: %s", httpErr)
	}
	defer resp.Body.Close()

	body, bodyErr := io.ReadAll(resp.Body)
	if bodyErr != nil {
		log.Fatalf("read error %s", body)
	}

	return body
}

func (ips dns) AsSlice() []string {
	lines := strings.Split(string(ips), "\n")
	return lines
}
