package main

import "github.com/miekg/dns"

func handleQuery(w dns.ResponseWriter, r *dns.Msg) {
	if len(r.Question) == 0 {
		return
	}

	name := r.Question[0].Name
	command, args := parseQuery(name)

	var answer string
	var answers []string

	// switch command {
	// case "time":
	// 	answer = getTime(location)
	// case "weather":
	// 	_, lat, lon := findLocation(location)
	// 	answer = getWeather(lat, lon)
	// case "pi":
	// 	answer = getPi(location)
	// case "country":
	// 	answer = getCountry(location)
	// case "crypto":
	// 	answer = getCrypto(location)
	// default:
	// 	answer = getHelp()
	switch command {
	case "time":
		answer = getTime(args)

	case "weather":
		lat, lon := findLatLon(args)
		answer = getWeather(lat, lon)

	case "pi":
		answer = getPi(args)

	case "country":
		answer = getCountry(args)

	case "crypto":
		answer = getCrypto(args)

	case "uuid":
		answers = getUUIDs(args) // returns []string

	default:
		answer = getHelp()
	}

	m := new(dns.Msg)
	m.SetReply(r)

	if len(answers) > 0 {
		m.Answer = append(m.Answer, &dns.TXT{
			Hdr: dns.RR_Header{
				Name:   name,
				Rrtype: dns.TypeTXT,
				Class:  dns.ClassINET,
				Ttl:    0, // UUIDs must never be cached
			},
			Txt: answers,
		})
	} else {
		m.Answer = append(m.Answer, &dns.TXT{
			Hdr: dns.RR_Header{
				Name:   name,
				Rrtype: dns.TypeTXT,
				Class:  dns.ClassINET,
				Ttl:    30,
			},
			Txt: []string{answer},
		})
	}
	w.WriteMsg(m)

}
