package pikeregex

import (
    "testing"
)

var isMatch bool
var regex, text string

func TestSimpleCompile(t *testing.T) {	
    regex = "ab*c"
	rs := simpleCompile(regex)
	if string(rs) != regex {
        t.Errorf("Should have matched: >>%s<< and >>%s<<", regex, rs)
	}
}

func TestRegexCharStar(t *testing.T) {
    regex = "ab*c"
    text = "abc"
    isMatch = Match(regex, text)
    if !isMatch {
        t.Errorf("Should have matched: >>%s<< and >>%s<<", regex, text)
    }
    regex = "ab*c"
    text = "abbbc"
    isMatch = Match(regex, text)
    if !isMatch {
        t.Errorf("Should have matched: >>%s<< and >>%s<<", regex, text)
    }
    regex = "ab*c"
    text = "aacc"
    isMatch = Match(regex, text)
    if !isMatch {
        t.Errorf("Should have matched: >>%s<< and >>%s<<", regex, text)
    }

    regex = "ab.*c"
    text = "abIamCornhulio-howdyhozcc"
    isMatch = Match(regex, text)
    if !isMatch {
        t.Errorf("Should have matched: >>%s<< and >>%s<<", regex, text)
    }

    regex = "aa*b.*c"
    text = "abIamCornhulio-howdyhozcc"
    isMatch = Match(regex, text)
    if !isMatch {
        t.Errorf("Should have matched: >>%s<< and >>%s<<", regex, text)
    }
    regex = "aa*bb*cc*"
    text = "abc"
    isMatch = Match(regex, text)
    if !isMatch {
        t.Errorf("Should have matched: >>%s<< and >>%s<<", regex, text)
    }

    regex = "aa*bb*cc*"
    text = "acbc"
    isMatch = Match(regex, text)
    if isMatch {
        t.Errorf("Should NOT have matched: >>%s<< and >>%s<<", regex, text)
    }
}


func TestRegexEndingAnchor(t *testing.T) {
    regex = "blue$"
    text = "armchair blue"
    isMatch = Match(regex, text)
    if !isMatch {
        t.Errorf("Should have matched: >>%s<< and >>%s<<", regex, text)
    }

    regex = "blue$"
    text = "lue"
    isMatch = Match(regex, text)
    if isMatch {
        t.Errorf("Should NOT have matched: >>%s<< and >>%s<<", regex, text)
    }

    regex = "blue$"
    text = "blue suede"
    isMatch = Match(regex, text)
    if isMatch {
        t.Errorf("Should NOT have matched: >>%s<< and >>%s<<", regex, text)
    }

    regex = "^blue$"
    text = "armchair blue"
    isMatch = Match(regex, text)
    if isMatch {
        t.Errorf("Should NOT have matched: >>%s<< and >>%s<<", regex, text)
    }

    regex = "^Blue$"
    text = "Blue"
    isMatch = Match(regex, text)
    if !isMatch {
        t.Errorf("Should have matched: >>%s<< and >>%s<<", regex, text)
    }
}


func TestRegexBeginningAnchor(t *testing.T) {
    regex = "^fo."
    text = "fox"
    isMatch = Match(regex, text)
    if !isMatch {
        t.Errorf("Should have matched: >>%s<< and >>%s<<", regex, text)
    }

    regex = "^fo."
    text = "the fox"
    isMatch = Match(regex, text)
    if isMatch {
        t.Errorf("Should NOT have matched: >>%s<< and >>%s<<", regex, text)
    }

    regex = "^fo."
    text = " fox"
    isMatch = Match(regex, text)
    if isMatch {
        t.Errorf("Should NOT have matched: >>%s<< and >>%s<<", regex, text)
    }
}

func TestRegexDotModifier(t *testing.T) {
    regex = "fo."
    text = "fox"
    isMatch = Match(regex, text)
    if !isMatch {
        t.Errorf("Should have matched: >>%s<< and >>%s<<", regex, text)
    }

    regex = "fo."
    text = "the fox is smart"
    isMatch = Match(regex, text)
    if !isMatch {
        t.Errorf("Should have matched: >>%s<< and >>%s<<", regex, text)
    }

    regex = "f..l"
    text = "my feelings"
    isMatch = Match(regex, text)
    if !isMatch {
        t.Errorf("Should have matched: >>%s<< and >>%s<<", regex, text)
    }

    regex = "f..l"
    text = "your fabulous face"
    isMatch = Match(regex, text)
    if isMatch {
        t.Errorf("Should NOT have matched: >>%s<< and >>%s<<", regex, text)
    }
}

func TestRegexNoModifiers(t *testing.T) {
    regex = "foo"
    text = "foo"
    isMatch = Match(regex, text)
    if !isMatch {
        t.Errorf("Should have matched: >>%s<< and >>%s<<", regex, text)
    }

	regex = "r5tgt6 \"foo\""
	text = "r5tgt6 \"foo\""
    isMatch = Match(regex, text)
    if !isMatch {
        t.Errorf("Should have matched: >>%s<< and >>%s<<", regex, text)
    }
}

func TestEmptyRegex(t *testing.T) {
    regex = ""
    text = ""
    isMatch = Match(regex, text)
    if !isMatch {
        t.Errorf("Should have matched: >>%s<< and >>%s<<", regex, text)
    }

    text = "z"
    isMatch = Match(regex, text)
    if isMatch {
        t.Errorf("Should not have matched: >>%s<< and >>%s<<", regex, text)
    }
}
