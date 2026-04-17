package main

import "github.com/miekg/dns"

func handleQuery(w dns.ResponseWriter, r *dns.Msg) {

	name := r.Question[0].Name
	command, location := parseQuery(name)
	locTimezone := getTime(location)
	var answer string

	switch command {
	case "time":
		answer = getTime(location)
	case "weather":
		_, lat, lon := findLocation(location)
		answer = getWeather(lat, lon)
	case "pi":
		answer = getPi(location)
	default:
		answer = getHelp()

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
