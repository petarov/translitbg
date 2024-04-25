package translitbg

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type TranslitBG struct {
	chars  map[string]string
	tokens map[string]string
	regex  *regexp.Regexp
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
		regex,
	}
}

func (tr *TranslitBG) Encode(input string) (string, error) {
	length := len(input)
	if length == 0 {
		return "", nil
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
			if isComboChar(ch) && isUpperBGChar(ch) && (ch2 == 0 || isUpperBGChar(ch2)) {
				dest.WriteString(strings.ToUpper(token))
			} else {
				dest.WriteString(token)
			}
		}
	}

	return dest.String(), nil
}

// isBGChar returns true, if the rune r is a cyrillic rune
// See https://symbl.cc/en/alphabets/bulgarian
func isBGChar(r rune) bool {
	return (r >= 1040 && r <= 1103) || r == 1117 || r == 1037
}

// isComboChar returns true, if the rune r is to be transformed into a
// combination of latin characters
func isComboChar(r rune) bool {
	switch r {
	case 1046, 1078, // Ж, ж
		1062, 1094, // Ц, ц
		1063, 1095, // Ч, ч
		1064, 1096, // Ш, ш
		1065, 1097, // Щ, щ
		1070, 1102, // Ю, ю
		1071, 1103: // Я, я
		return true
	}
	return false
}

// isUpperBGChar returns true, if the rune r is an uppercase cyrillic rune
func isUpperBGChar(r rune) bool {
	return (r >= 1040 && r <= 1071) || r == 1037
}
