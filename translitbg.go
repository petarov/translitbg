package translitbg

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
)

type TranslitBG struct {
	chars  map[string]string
	tokens map[string]string
	combos map[rune]string
	regex  *regexp.Regexp
}

// isBGChar returns true, if the rune r is a cyrillic rune
// See https://symbl.cc/en/alphabets/bulgarian
func isBGChar(r rune) bool {
	return (r >= 1040 && r <= 1103) || r == 1117 || r == 1037
}

// isUpperBGChar returns true, if the rune r is an uppercase cyrillic rune
func isUpperBGChar(r rune) bool {
	return (r >= 1040 && r <= 1071) || r == 1037
}

func New() *TranslitBG {
	pattern := "^\\w+$"
	regex, err := regexp.Compile(pattern)
	if err != nil {
		panic(fmt.Errorf("error compiling regex: %v", err))
	}

	return &TranslitBG{
		STREAMLINED,
		STREAMLINED_TOKENS,
		STREAMLINED_CYR2COMBO_UC,
		regex,
	}
}

// Encode transliterates Bulgarian string input to its latin equivalent.
// Non-cyrillic characters will be left as they are.
func (tr *TranslitBG) Encode(input string) (string, error) {
	source := bytes.NewBufferString(input)
	dest := bytes.NewBuffer(nil)
	ch_1 := ""

	for {
		ch, _, err := source.ReadRune()

		if err == io.EOF {
			break
		} else if err != nil {
			return "", fmt.Errorf("error reading source text: %v", err)
		}

		if isBGChar(ch) {
			ch2, _, err := source.ReadRune()

			if err != nil && err != io.EOF {
				return "", fmt.Errorf("error reading source text: %v", err)
			} else if err == nil {
				token := string([]rune{ch, ch2})

				found, ok := tr.tokens[token]
				if ok {
					ch3, _, err := source.ReadRune()
					if err != io.EOF || !tr.regex.MatchString(string(ch3)) {
						source.UnreadRune()
						dest.WriteString(found)
						ch_1 = string(ch3)
						continue
					} else {
						source.UnreadRune()
					}
				} else {
					source.UnreadRune()
				}
			}

			token, ok := tr.chars[string(ch)]
			if ok {
				ucc, ok := tr.combos[ch]
				if ok && (ch2 == 0 || isUpperBGChar(ch2) || !isBGChar(ch2) || len(tr.chars[ch_1]) > 0) {
					dest.WriteString(ucc)
				} else {
					dest.WriteString(token)
				}
			}
		} else {
			dest.WriteRune(ch)
		}

		ch_1 = string(ch)
	}

	return dest.String(), nil
}
