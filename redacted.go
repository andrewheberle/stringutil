package stringutil

// Redacted can be used to print a redacted string using fmt.Println(Redacted("this will be printed as all *"))
type Redacted string

var RedactedCharacter = '*'

func (r Redacted) String() string {
	b := make([]byte, 0)
	for range r {
		b = append(b, byte(RedactedCharacter))
	}

	return string(b)
}

func Redact(s string, c rune) string {
	b := make([]byte, 0)
	for range s {
		b = append(b, byte(c))
	}

	return string(b)
}
