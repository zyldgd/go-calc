package gocalc

type scanner struct {
	source []rune
	index  int
	char   rune // next one
}

func newScanner(e string) *scanner {
	if len(e) == 0 {
		return nil
	}

	l := &scanner{
		source: []rune(e),
		index:  0,
	}

	l.char = l.source[0]

	return l
}

func (s *scanner) next() {
	s.index++
	if s.index < len(s.source) {
		s.char = s.source[s.index]
		return
	}
	s.char = -1
}

func (s *scanner) nextChar() rune {
	idx := s.index + 1
	if idx < len(s.source) {
		return s.source[idx]
	}
	return -1
}

func (s *scanner) skip() {
	for isSpace(s.char) {
		s.next()
	}
}

func (s *scanner) scan() (token, string) {
	s.skip()
	if s.index >= len(s.source) {
		return EOF, ""
	}

	tok, lit := Illegal, ""

	switch {
	case isDecimal(s.char):
		tok, lit = s.scanNumber()
	case isLetter(s.char) || '_' == s.char:
		tok, lit = s.scanIdentifier()
	case '"' == s.char:
		tok, lit = s.scanString()
	case '\'' == s.char:
		tok, lit = s.scanChar()
	default:
		switch s.char {
		case '+':
			tok, lit = OpAdd, "+"
		case '-':
			tok, lit = OpMinus, "-"
		case '*':
			tok, lit = OpMultiply, "*"
		case '/':
			tok, lit = OpDivide, "/"
		case '%':
			tok, lit = OpModulus, "%"
		case '(':
			tok, lit = OpLParen, "("
		case ')':
			tok, lit = OpRParen, ")"
		case '[':
			tok, lit = OpLBracket, "["
		case ']':
			tok, lit = OpRBracket, "]"
		case '.':
			tok, lit = OpAccess, "."
		case '!':
			if '=' == s.nextChar() {
				s.next()
				tok, lit = OpNeq, "!="
			} else {
				tok, lit = OpNot, "!"
			}
		case '=':
			if '=' == s.nextChar() {
				s.next()
				tok, lit = OpEq, "=="
			} else {
				tok, lit = Illegal, "" // Illegal
			}
		case '&':
			if '&' == s.nextChar() {
				s.next()
				tok, lit = OpAnd, "&&"
			} else {
				tok, lit = OpBitwiseAnd, "&"
			}
		case '|':
			if '|' == s.nextChar() {
				s.next()
				tok, lit = OpOr, "||"
			} else {
				tok, lit = OpBitwiseOr, "|"
			}
		case '^':
			tok, lit = OpBitwiseXor, "^"
		case '~':
			tok, lit = OpBitwiseNot, "~"
		case '<':
			if '=' == s.nextChar() {
				s.next()
				tok, lit = OpLte, "<="
			} else if '<' == s.nextChar() {
				s.next()
				tok, lit = OpBitwiseLShift, "<<"
			} else {
				tok, lit = OpLt, "<"
			}
		case '>':
			if '=' == s.nextChar() {
				s.next()
				tok, lit = OpGte, ">="
			} else if '>' == s.nextChar() {
				s.next()
				tok, lit = OpBitwiseRShift, ">>"
			} else {
				tok, lit = OpGt, ">"
			}
		}

		s.next()
	}

	return tok, lit
}

func (s *scanner) scanNumber() (token, string) {
	start := s.index
	tok := Integer
	for isDecimal(s.char) {
		s.next()
	}

	if s.char == '.' {
		tok = Float
		s.next()
		if !isDecimal(s.char) {
			return Illegal, ""
		}
		for isDecimal(s.char) {
			s.next()
		}
	}
	return tok, string(s.source[start:s.index])
}

func (s *scanner) scanEscape() bool {
	s.next()
	switch s.char {
	case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', '"', '0':
		return true
	default:
		//msg := "unknown escape sequence"
		//if l.char < 0 {
		//	msg = "escape sequence not terminated"
		//}
		// s.error(offs, msg)
		return false
	}
}

func (s *scanner) scanString() (token, string) {
	start := s.index // start with "

	for {
		s.next()
		if s.char == '"' {
			break
		} else if s.char == '\\' {
			if !s.scanEscape() {
				return Illegal, ""
			}
		} else if s.char == -1 {
			return Illegal, ""
		}
	}

	if s.char != '"' || start == s.index {
		return Illegal, ""
	}
	s.next()

	return String, string(s.source[start:s.index])
}

func (s *scanner) scanChar() (token, string) {
	start := s.index // start with '
	tok := Char

	s.next()
	if s.char == '\\' {
		if !s.scanEscape() {
			return Illegal, ""
		}
	} else if s.char < 0 {
		return Illegal, ""
	}
	s.next()
	if s.char != '\'' {
		return Illegal, ""
	}

	s.next()
	return tok, string(s.source[start:s.index])
}

func (s *scanner) scanIdentifier() (token, string) {
	start := s.index

	for isLetter(s.char) || isDecimal(s.char) || '_' == s.char {
		s.next()
	}

	return Ident, string(s.source[start:s.index])
}
