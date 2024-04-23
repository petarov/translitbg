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
}

func New() *TranslitBG {
	return &TranslitBG{
		STREAMLINED,
		STREAMLINED_TOKENS,
	}
}

func (tr *TranslitBG) Run(input string) (string, error) {
	if len(input) == 0 {
		return "", nil
	}

	pattern := "^\\w+$"
	regex, err := regexp.Compile(pattern)
	if err != nil {
		panic(fmt.Errorf("error compiling regex: %v", err))
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

		// TODO: is this cyr char

		ch2, _, err := source.ReadRune()

		if err != nil && err != io.EOF {
			return "", fmt.Errorf("error reading source text: %v", err)
		} else if err == nil {
			token := string([]rune{ch, ch2})

			found, ok := tr.tokens[token]
			if ok {
				ch3, _, err := source.ReadRune()
				if err != io.EOF || !regex.MatchString(string(ch3)) {
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
			dest.WriteString(token)
		} else {
			dest.WriteRune(ch)
		}
	}

	return dest.String(), nil
}
