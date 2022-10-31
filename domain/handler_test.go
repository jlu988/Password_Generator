package domain

import (
	"testing"
)

var testHelper PasswordGenerator

func TestRandomPassString(t *testing.T) {
	expectedSize := 285
	acutal := testHelper.randomizedString()

	if expectedSize != len(acutal) {
		t.Fatalf("TestRandomPassString() - want: %+v but got: %+v ", expectedSize, acutal)
	}
}

func TestGeneratePassword(t *testing.T) {
	testHelper.Size = 16
	expectedSize := 16
	actual := testHelper.generatePassword()

	if expectedSize != len(actual) {
		t.Fatalf("generatePassword() - want: %+v but got: %+v\n", expectedSize, len(actual))
	}
}
