// we will use this file to generate random names and numbers for test cases

package util

import (
	"math/rand"
	"strings"
	"time"
)

const alpha = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder

	k := 26
	for i := 0; i < n; i++ {
		alphabet := alpha[rand.Intn(k)]
		sb.WriteByte(alphabet)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(10)
}

func RandomBalance() int64 {
	return RandomInt(10, 10000000000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "INR", "CAD", "EUR", "YUAN"}
	return currencies[rand.Intn(5)]
}
