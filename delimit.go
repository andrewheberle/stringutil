package stringutil

import (
	"fmt"
	"strings"
)

// Delimit will join the provided string slice using the provided seperator,
// optionally using an alternate seperator for the final element of the slice
func Delimit(sl []string, sep string, lsep ...string) string {
	switch n := len(sl); n {
	case 0:
		// zero length or nil slice returns an empty string
		return ""
	case 1:
		// slice with one string returns that string unchanged
		return sl[0]
	default:
		if len(lsep) == 0 {
			// no value for lsep means its a normal join
			return strings.Join(sl, sep)
		}

		// otherwise join all but final element then add lsep and last element
		last := sl[n-1]
		s := strings.Join(sl[:n-1], sep)
		return fmt.Sprintf("%s%s%s", s, lsep[0], last)
	}
}
