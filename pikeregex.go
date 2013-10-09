package pikeregex

// search for c*regex at beginning of text
func matchstar(c rune, regex []rune, text []rune) bool {
	if matchhere(regex, text) {
		return true
	}
	for len(text) > 0 && (text[0] == c || c == '.') {
		text = text[1:]
		if matchhere(regex, text) {
			return true
		}
	}
    return false
}

// search for regex at beginning of text
func matchhere(regex []rune, text []rune) bool {
    if len(regex) == 0 {
        return true
    }
    if len(regex) > 1 && regex[1] == '*' {
        return matchstar(regex[0], regex[2:], text)
    }
	if regex[0] == '$' && len(regex) == 1 {
		return len(text) == 0
	}
	if len(text) > 0  && (regex[0] == '.' || regex[0] == text[0]) {
		return matchhere(regex[1:], text[1:])
	}
    return false
}

// search for regex anywhere in the text
func Match(regex string, text string) bool {
    if len(text) == 0 || len(regex) == 0 {
        return len(regex) == len(text)
    }

    runerx := compile(regex)
    runetxt := []rune(text)

    if runerx[0] == '^' {
        return matchhere(runerx[1:], runetxt)
    }
    for ; len(runetxt) > 0; runetxt = runetxt[1:] {
        if matchhere(runerx, runetxt) {
            return true
        }
    }
    return false
}

// one enhancement: allow + (1 or more) notation
func compile(regex string) (regslc []rune) {
	regslc = make([]rune, 0, len(regex) + 10)

    for _, r := range regex {
        if r == '+' {
            regslc = append(regslc, regslc[len(regslc) - 1], '*')
        } else {
            regslc = append(regslc, r)
        }
    }	
	return regslc
}
