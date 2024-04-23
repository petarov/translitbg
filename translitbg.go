package translitbg

import (
	"fmt"
	"regexp"
	"strings"
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

func (tr *TranslitBG) Run(input string) string {
	result := []string{}
	length := len(input)

	pattern := "^\\w+$"
	regex, err := regexp.Compile(pattern)
	if err != nil {
		panic(fmt.Errorf("error compiling regex: %v", err))
	}

	for i := 0; i < length; i++ {
		cur := string(input[i])

		if i+1 < length {
			next := string(input[i+1])
			curToken := cur + next

			found, ok := STREAMLINED_TOKENS[curToken]
			if ok {
				if i+2 < length {
					nextNext := string(input[i+2])
					if regex.MatchString(nextNext) {
						result = append(result, found)
						i += 1
						continue
					}
				}
			}
		}

		token, ok := STREAMLINED[cur]
		if ok {
			result = append(result, token)
		} else {
			result = append(result, cur)
		}
	}

	return strings.Join(result, "")
}
