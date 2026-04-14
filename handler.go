package main

import "github.com/miekg/dns"

func handleQuery(w dns.ResponseWriter, r *dns.Msg) {

	name := r.Question[0].Name
	command, location := parseQuery(name)
	locTimezone := getTime(location)
	var answer string
	if command == "time" {
		answer = (locTimezone)
	} else {
		answer = "error: unknown command"
	}

	m := new(dns.Msg)
	m.SetReply(r)
	m.Answer = append(m.Answer, &dns.TXT{
		Hdr: dns.RR_Header{
			Name:   r.Question[0].Name,
			Rrtype: dns.TypeTXT,
			Class:  dns.ClassINET,
			Ttl:    30,
		},
		Txt: []string{answer},
	})
	w.WriteMsg(m)

}
