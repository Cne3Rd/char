package char

// char is a simple program that switch between
// between forward slash and backslash
// it takes a string in backtick and changes every ocurrence of
// forwrd slash to backwrd slash and vice versa
// depending on the function you called or
// it is useful when dealing with file path
// while switching between windows and linux

// chages backslash to forward slash
// (e.g char.FwSlash(`github.com\go-sql-driver\mysql` use backtick)
func FwSlash(path string) string {
	output := make([]rune, 0)
	for k, ch := range path {
		output = append(output, ch)
		//fmt.Printf("%v, %v, %v\n", k, ch, string(ch))
		if ch == rune(92) {
			output[k] = rune(47)
		}
	}

	return string(output)
}

// changes forward slash to backslash
// (e.g char.BwSlash(`github.com/go-sql-driver/mysql`) use backtick)
func BwSlash(path string) string {
	output := make([]rune, 0)
	for k, ch := range path {
		output = append(output, ch)
		if ch == rune(47) {
			output[k] = rune(92)
		}
	}

	return string(output)
}
