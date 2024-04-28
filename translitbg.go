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

// tryDoBulgaria returns true for the case where input s is the text "България".
// In this case the "ъ" needs to be trasformed into an "u" as the law dictates
func tryDoBulgaria(input string) (bool, string) {
	runes := []rune(input)
	dest := make([]rune, 8)

	for i, r := range runes {
		if BULGARIA_CYR_LOW[i] == r {
			dest[i] = BULGARIA_LAT_LOW[i]
		} else if BULGARIA_CYR_UP[i] == r {
			dest[i] = BULGARIA_LAT_UP[i]
		} else {
			return false, ""
		}
	}

	return true, string(dest)
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
	length := len(input)
	if length == 0 {
		return "", nil
	} else if length == 16 {
		ok, result := tryDoBulgaria(input)
		if ok {
			return result, nil
		}
	}

	source := bytes.NewBufferString(input)
	dest := bytes.NewBuffer(nil)

	for {
		ch, _, err := source.ReadRune()

		if err == io.EOF {
			break
		} else if err != nil {
			return "", fmt.Errorf("error reading source text: %v", err)
		}

		if !isBGChar(ch) {
			dest.WriteRune(ch)
			continue
		}

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
			if ok && (ch2 == 0 || isUpperBGChar(ch2)) {
				dest.WriteString(ucc)
			} else {
				dest.WriteString(token)
			}
		}
	}

	return dest.String(), nil
}
