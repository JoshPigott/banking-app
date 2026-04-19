package domain

import "unicode"

type Password struct {
	value string
}

func NewPassword(p string) Password {
	return Password{
		value: p,
	}
}

func (p Password) containsUpperCase() bool {
	for _, c := range p.value {
		if unicode.IsUpper(c) {
			return true
		}
	}
	return false
}

func (p Password) containsLowerCase() bool {
	for _, c := range p.value {
		if unicode.IsLower(c) {
			return true
		}
	}
	return false
}

func (p Password) containsDigit() bool {
	for _, c := range p.value {
		if unicode.IsDigit(c) {
			return true
		}
	}
	return false
}

func (p Password) isLongEnough() bool {
	return len(p.value) >= minPasswordLen
}

func (p Password) IsStrong() bool {
	return p.containsUpperCase() &&
		p.containsLowerCase() &&
		p.containsDigit() &&
		p.isLongEnough()
}
