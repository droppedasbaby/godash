package strutil

import "strings"

// PadLeft returns a string of length n that is filled with the string p on the left.
// If the length of s is already greater than or equal to n, the original string is returned.
// If n is negative, behavior is undefined.
func PadLeft(s string, n int, p string) string {
	if len(s) >= n || len(p) == 0 {
		return s
	}
	r := strings.Builder{}
	r.Grow(n)
	for i := 0; i < (n-len(s))/len(p); i++ {
		r.WriteString(p)
	}
	r.WriteString(strings.Repeat(p, (n-len(s))%len(p)))
	r.WriteString(s)
	return r.String()
}

// PadRight returns a string of length n that is filled with the string p on the right.
// If the length of s is already greater than or equal to n, the original string is returned.
// If n is negative, behavior is undefined.
func PadRight(s string, n int, p string) string {
	if len(s) >= n || len(p) == 0 {
		return s
	}
	r := strings.Builder{}
	r.Grow(n)
	r.WriteString(s)
	for i := 0; i < (n-len(s))/len(p); i++ {
		r.WriteString(p)
	}
	r.WriteString(strings.Repeat(p, (n-len(s))%len(p)))
	return r.String()
}

// Pad returns a string of length n that is centered and filled with the string p.
// If the length of s is already greater than or equal to n, the original string is returned.
// If n is negative, behavior is undefined.
// If an odd number of padding characters is required, the extra characters are added to the left.
func Pad(s string, n int, p string) string {
	if len(s) >= n || len(p) == 0 {
		return s
	}
	r := strings.Builder{}
	r.Grow(n)
	paddingLen := (n - len(s)) / 2
	extraPad := (n - len(s)) % 2 // 1 if extra padding is needed, otherwise 0

	for i := 0; i < (paddingLen+extraPad)/len(p); i++ {
		r.WriteString(p)
	}
	r.WriteString(strings.Repeat(p, (paddingLen+extraPad)%len(p)))

	r.WriteString(s)

	for i := 0; i < paddingLen/len(p); i++ {
		r.WriteString(p)
	}
	r.WriteString(strings.Repeat(p, paddingLen%len(p)))

	return r.String()
}
