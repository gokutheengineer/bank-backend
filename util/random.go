package util

import (
	"math/rand"
	"time"
)

const lowerCaseLetters = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
	randSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSource)

	return random.Int63n(max - min + 1)
}

func RandomString(length int) string {
	randSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSource)
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = lowerCaseLetters[random.Intn(len(lowerCaseLetters))]
	}

	return string(result)
}

func RandomOwner() string {
	return RandomString(8)
}

func RandomMoneyAmount() int64 {
	return RandomInt(0, 1000000)
}

func RandomCurrency() string {
	currencies := []string{USD, TRY, EUR}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}

func RandomEmail() string {
	return RandomString(6) + "@" + RandomString(3) + ".com"
}

func RandomFullname() string {
	return RandomString(5) + " " + RandomString(5)
}

func RandomPassword() string {
	return HashPasswordBlake(RandomString(5))
}
