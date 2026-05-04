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
	"crypto":  true,
	"uuid":    true,
}
var noArgsOK = map[string]bool{
	"uuid": true,
}

func parseQuery(domain string) (command, args string) {
	domain = strings.TrimSuffix(domain, ".")
	parts := strings.SplitN(domain, ".", 2)

	if len(parts) < 1 {
		return "", ""
	}

	command = strings.ToLower(parts[0])
	if !validCommands[command] {
		return "", ""
	}

	if len(parts) == 2 {
		args = parts[1]
	}

	if args == "" && !noArgsOK[command] {
		return "", ""
	}

	return command, args
}

func main() {
	dns.HandleFunc(".", handleQuery)

	server := &dns.Server{
		Addr: ":5053",
		Net:  "udp",
	}

	fmt.Println("DNSorcery running on :5053")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("error:", err)
	}
}
