package strreq

import (
	"strings"
	"unicode"
)

type Requirements struct {
	MinLength int `toml:"minLength"`
	MaxLength int `toml:"maxLength"`

	MinLetter int `toml:"minLetter"`
	MinLower  int `toml:"minLower"`
	MinUpper  int `toml:"minUpper"`
	MinNumber int `toml:"minNumber"`

	ValidChars   string `toml:"validChars"`
	InvalidChars string `toml:"invalidChars"`
}

func (r Requirements) IsValid(s string) bool {
	if r.MinLength > 0 && len(s) < r.MinLength {
		return false
	}
	if r.MaxLength > 0 && len(s) > r.MaxLength {
		return false
	}

	if r.MinLetter > 0 {
		n := 0
		for _, r := range s {
			if unicode.IsLetter(r) {
				n += 1
			}
		}
		if n < r.MinLetter {
			return false
		}
	}
	if r.MinLower > 0 {
		n := 0
		for _, r := range s {
			if unicode.IsLower(r) {
				n += 1
			}
		}
		if n < r.MinLower {
			return false
		}
	}
	if r.MinUpper > 0 {
		n := 0
		for _, r := range s {
			if unicode.IsUpper(r) {
				n += 1
			}
		}
		if n < r.MinUpper {
			return false
		}
	}
	if r.MinNumber > 0 {
		n := 0
		for _, r := range s {
			if unicode.IsNumber(r) {
				n += 1
			}
		}
		if n < r.MinNumber {
			return false
		}
	}

	if len(r.ValidChars) > 0 {
		for _, l := range s {
			if !strings.ContainsRune(r.ValidChars, l) {
				return false
			}
		}
	}
	if len(r.InvalidChars) > 0 {
		for _, l := range s {
			if strings.ContainsRune(r.InvalidChars, l) {
				return false
			}
		}
	}

	return true
}
