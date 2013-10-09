 
 package pikeregex
 
 // search for c*regex at beginning of text
 func matchstar(c rune, regex []rune, text []rune) bool {
     for {
         if matchhere(regex, text) {
             return true
         }
         if ! (len(text) > 0 && (text[0] == c || c == '.')) {
             return false
         }
         text = text[1:]
     }
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
     runerx := compile(regex)
     runetxt := []rune(text)
 
     if len(runerx) > 0 && runerx[0] == '^' {
         return matchhere(runerx[1:], runetxt)
     }
 
     for {
         if matchhere(runerx, runetxt) {
             return true
         }
         if len(runetxt) == 0 {
             return false
         }
         runetxt = runetxt[1:]
     }
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
 
