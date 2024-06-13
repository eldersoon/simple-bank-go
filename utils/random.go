package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvqyxz"

func init() {
	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))
}

// RandomInt generates a random integer between min and max
func RandomInt (min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

// RandomString generate a random word contain size n
func RandomString (n int) string {
	var sb strings.Builder
	j := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(j)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random name for owner account
func RandomOwner () string {
	return RandomString(6)
}

// RandomMoney return a integer to fill account amount
func RandomMoney () int64 {
	return RandomInt(0, 1000000)
}

// RandomCurrency choose a currency for transfer
func RandomCurrency () string {
	currencies := []string{"BRL", "USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}