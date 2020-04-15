package utils

import "testing"

func TestStringEmpty(t *testing.T) {
	result := StringEmpty("")
	t.Logf("isStringEmpty=%t", result)

	result = StringEmpty(" ")
	t.Logf("isStringEmpty=%t", result)

	result = StringEmpty(" 123 ")
	t.Logf("isStringEmpty=%t", result)
}

func TestStringNotEmpty(t *testing.T) {
	result := StringNotEmpty("")
	t.Logf("isStringNotEmpty=%t", result)

	result = StringNotEmpty("  ")
	t.Logf("isStringNotEmpty=%t", result)

	result = StringNotEmpty(" 123 ")
	t.Logf("isStringNotEmpty=%t", result)
}

func TestStringAllEmpty(t *testing.T) {
	result := stringAllEmpty("", " ")
	t.Logf("stringAllEmpty=%t", result)

	result = stringAllEmpty("", " ", " 123 ", "123")
	t.Logf("stringAllEmpty=%t", result)
}

func TestStringAnyEmpty(t *testing.T) {
	result := StringAnyEmpty(" 123 ", "  ", "456")
	t.Logf("stringAnyEmpty=%t", result)

	result = StringAnyEmpty(" 123 ", " 789 ", "456")
	t.Logf("stringAnyEmpty=%t", result)
}
