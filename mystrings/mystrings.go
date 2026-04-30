package mystrings

// the name should always be capitalized in order to make this fn public
// there's no main func cause this file is just to export some fns
func Reverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result

	}
	return result
}
