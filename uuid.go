package main

import (
	"crypto/rand"
	"fmt"
	"strconv"
)

func getUUIDs(args string) []string {
	count := 1
	if n, err := strconv.Atoi(args); err == nil && n >= 1 && n <= 10 {
		count = n
	}

	uuids := make([]string, count)
	for i := range uuids {
		uuids[i] = newUUID()
	}
	return uuids
}

func newUUID() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "error: could not generate uuid"
	}
	// Set version 4 and variant bits.
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
