package security


import "strings"

const (
	LOWER      = "abcdefghijklmnopqrstuvwxyz"
	UPPER      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SYMBOL     = "!@#$%^&*()-+=_[]{}.,/?;:\\"
	NUMBERS    = "0123456789"
	WHITESPACE = " \t\n\r"
)

var (
	DefaultPasswordCheck = NewPasswordCheck(8, false, false, false, false)
)

type PasswordCheck struct {
	ContainsUpper  bool
	ContainsLower  bool
	ContainsSymbol bool
	ContainsNumber bool
	MinimumLength  int
}

func NewPasswordCheck(minLen int, upper, lower, symbol, number bool) *PasswordCheck {
	return &PasswordCheck{
		ContainsLower:  lower,
		ContainsNumber: number,
		ContainsSymbol: symbol,
		ContainsUpper:  upper,
		MinimumLength:  minLen,
	}
}

func (c *PasswordCheck) CheckPasswordValidity(password string) bool {
	if c.MinimumLength > 0 {
		if len(password) < c.MinimumLength {
			return false
		}
	}
	if c.contained(password, WHITESPACE) {
		return false
	}
	if c.ContainsUpper && !c.contained(password, UPPER) {
		return false
	}
	if c.ContainsLower && !c.contained(password, LOWER) {
		return false
	}
	if c.ContainsSymbol && !c.contained(password, SYMBOL) {
		return false
	}
	if c.ContainsNumber && !c.contained(password, NUMBERS) {
		return false
	}
	return true
}

func (c *PasswordCheck) contained(password, charset string) bool {
	for i := range charset {
		if strings.ContainsAny(password, charset[i:i+1]) {
			return true
		}
	}
	return false
}