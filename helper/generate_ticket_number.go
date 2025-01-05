package helper

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateTicketNumber() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const digits = "0123456789"
	rand.Seed(time.Now().UnixNano())

	var ticketNumber strings.Builder
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			ticketNumber.WriteByte(letters[rand.Intn(len(letters))])
		} else {
			ticketNumber.WriteByte(digits[rand.Intn(len(digits))])
		}
	}
	return ticketNumber.String()
}
