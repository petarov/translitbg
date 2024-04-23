package translitbg

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
)

var (
	// Възстановяването на оригиналната дума не е водещ принцип!
	STREAMLINED = map[string]string{
		// lower case
		"а": "a",
		"б": "b",
		"в": "v",
		"г": "g",
		"д": "d",
		"е": "e",
		"ж": "zh",
		"з": "z",
		"и": "i",
		"ѝ": "i",
		"й": "y",
		"к": "k",
		"л": "l",
		"м": "m",
		"н": "n",
		"о": "o",
		"п": "p",
		"р": "r",
		"с": "s",
		"т": "t",
		"у": "u",
		"ф": "f",
		"х": "h",
		"ц": "ts",
		"ч": "ch",
		"ш": "sh",
		"щ": "sht",
		"ъ": "a",
		"ь": "y",
		"ю": "yu",
		"я": "ya",
		// upper case
		"А": "A",
		"Б": "B",
		"В": "V",
		"Г": "G",
		"Д": "D",
		"Е": "E",
		"Ж": "Zh",
		"З": "Z",
		"И": "I",
		"Ѝ": "I",
		"Й": "Y",
		"К": "K",
		"Л": "L",
		"М": "M",
		"Н": "N",
		"О": "O",
		"П": "P",
		"Р": "R",
		"С": "S",
		"Т": "T",
		"У": "U",
		"Ф": "F",
		"Х": "H",
		"Ц": "Ts",
		"Ч": "Ch",
		"Ш": "Sh",
		"Щ": "Sht",
		"Ъ": "A",
		"Ь": "Y",
		"Ю": "Yu",
		"Я": "Ya",
	}

	STREAMLINED_TOKENS = map[string]string{
		// Буквеното съчетание „ия“, когато е в края на думата, се изписва и предава чрез „ia“
		"ия": "ia",
		"Ия": "Ia",
		"иЯ": "iA",
		"ИЯ": "IA",
	}
)

type TranslitBG struct {
	params []int
}

func New() *TranslitBG {
	// TODO: add type
	tr := &TranslitBG{}
	return tr
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

			found, ok := STREAMLINED_TOKENS[token]
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

		token, ok := STREAMLINED[string(ch)]
		if ok {
			dest.WriteString(token)
		} else {
			dest.WriteRune(ch)
		}
	}

	return dest.String(), nil
}
