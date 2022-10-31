package domain

import (
	"math/rand"
	"strings"
	"time"
)

type PasswordGenerator struct {
	Size int `json:"size"`
}

func (ph *PasswordGenerator) GetPassword(size int) string {
	ph.Size = size
	return string(ph.generatePassword())
}

func (ph *PasswordGenerator) randomizedString() []byte {
	stringBuffer := map[string]string{
		"upperCase": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"lowerCase": "abcdefghijklmnopqrstuvwxyz",
		"symbols":   " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~",
		"number":    "0123456789",
	}
	passwordString := ""
	for _, val := range stringBuffer {
		passwordString += val
	}

	passwordString = strings.Repeat(passwordString, 3)

	return ph.shuffleString([]byte(passwordString))
}

func (ph *PasswordGenerator) shuffleString(strBuffer []byte) []byte {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(strBuffer), func(i, j int) {
		strBuffer[i], strBuffer[j] = strBuffer[j], strBuffer[i]
	})
	return strBuffer
}

func (ph *PasswordGenerator) generatePassword() []byte {
	password := ph.randomizedString()
	return password[:ph.Size]
}
