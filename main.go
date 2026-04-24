package main

import (
	"fmt"
	"strings"

	"github.com/miekg/dns"
)

var validCommands = map[string]bool{
	"time":    true,
	"weather": true,
	"pi":      true,
	"country": true,
	"convert": true,
	"crypto": true,
}

func parseQuery(domain string) (command, location string) {
	domain = strings.TrimSuffix(domain, ".")
	parts := strings.Split(domain, ".")
	if len(parts) < 2 {
		return "", ""
	}
	command = parts[0]
	strings.Join(parts[1:], " ")
	location = strings.Join(parts[1:], " ")

	if !validCommands[command] {
		return "", ""
	}

	if location == "" && command != "ip" {
		return "", ""
	}

	if location == "" {
		return "", ""
	}

	return command, location
}

func main() {
	dns.HandleFunc(".", handleQuery)

	server := &dns.Server{
		Addr: ":5053",
		Net:  "udp",
	}

	fmt.Println("DNS server running on :5053")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("error:", err)
	}
}
