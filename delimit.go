package stringutil

// Delimit will join the provided string slice using the provided separator,
// optionally using an alternate separator for the final element of the slice
func Delimit(s []string, sep string, last ...string) string {
	n := len(s)

	var dLast *string
	if len(last) > 0 {
		dLast = &last[0]
	}

	var str string
	for i := 0; i < n; i++ {
		switch {
		case i == n-2 && dLast != nil:
			str += s[i] + *dLast
		case i == n-1:
			str += s[i]
		default:
			str += s[i] + sep
		}
	}

	return str
}
