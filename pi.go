package main

import (
	"strings"
)

func getPi(command string) string {
	command = strings.ToLower(strings.TrimSpace(command))

	switch command {
	case "value":
		return "3.1415926535897932384626433"
	case "info":
		return "Pi is the ratio of a circle's circumference to its diameter. Archimedes of Syracuse is credited with the first theoretical calculation."
	case "digits":
		return "3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679"
	case "formula":
		return "Euler: e^(i*π) + 1 = 0 | Leibniz: π/4 = 1 - 1/3 + 1/5 - 1/7 + ..."
	case "fun":
		return "Pi is a 'transcendental number', meaning it is not the root of any non-zero polynomial with rational coefficients!"
	default:
		return "Try: value, info, digits, formula, or fun"
	}
}
