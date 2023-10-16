package strutil

import "strings"

func padHelper(s string, n int, p string, l bool, r bool) string {
	if len(s) >= n || len(p) == 0 {
		return s
	}

	b := strings.Builder{}
	b.Grow(n)
	pl := n - len(s)

	if l && r {
		pl /= 2
	}

	if l {
		for i := 0; i < pl/len(p); i++ {
			b.WriteString(p)
		}
		b.WriteString(p[:pl%len(p)])
		b.WriteString(s)
	}

	if l && r {
		b.WriteString(p[:pl%len(p)])
	}

	if !l && r {
		b.WriteString(s)
	}

	if r {
		for i := 0; i < pl/len(p); i++ {
			b.WriteString(p)
		}
		b.WriteString(p[:pl%len(p)])
	}

	return b.String()
}

// PadLeft returns a string of length n that is filled with the string p on the left.
// If the length of s is already greater than or equal to n, the original string is returned.
// If n is negative, behavior is undefined.
func PadLeft(s string, n int, p string) string {
	return padHelper(s, n, p, true, false)
}

// PadRight returns a string of length n that is filled with the string p on the right.
// If the length of s is already greater than or equal to n, the original string is returned.
// If n is negative, behavior is undefined.
func PadRight(s string, n int, p string) string {
	return padHelper(s, n, p, false, true)
}

// Pad returns a string of length n that is centered and filled with the string p.
// If the length of s is already greater than or equal to n, the original string is returned.
// If n is negative, behavior is undefined.
// If an odd number of padding characters is required, the extra characters are added to the left.
func Pad(s string, n int, p string) string {
	if len(s) >= n || len(p) == 0 {
		return s
	}

	b := strings.Builder{}
	b.Grow(n)

	pl := (n - len(s)) / 2
	ep := (n - len(s)) % 2

	pll, plr := pl, pl
	if ep == 1 {
		pll++
	}

	for i := 0; i < pll/len(p); i++ {
		b.WriteString(p)
	}
	b.WriteString(p[:pll%len(p)])
	b.WriteString(s)
	for i := 0; i < plr/len(p); i++ {
		b.WriteString(p)
	}
	b.WriteString(p[:plr%len(p)])

	return b.String()
}
