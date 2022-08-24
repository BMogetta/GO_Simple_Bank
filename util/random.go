package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Random interger between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Random string of lenght n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// Random owner name
func RandomOwner() string {
	return RandomString(6)
}

// Random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Random currency
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "YEN"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
