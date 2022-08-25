package telegram

func ParseArgs(text string) []string {
	ret := make([]string, 0)
	str := ""
	concat := false
	var tag rune
	for _, c := range text {
		if (c == ' ') && !concat {
			// Skip empty elements
			if (str != "") && (str != " ") {
				ret = append(ret, str)
			}
			str = ""
		} else if !concat && str == "" && ((c == '\'') || (c == '"')) {
			tag = c
			concat = true
		} else if concat && (c == tag) {
			tag = '\n'
			concat = false
		} else {
			str += string(c)
		}
	}

	ret = append(ret, str)

	return ret
}