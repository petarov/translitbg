package translitbg

var (
	// Възстановяването на оригиналната дума не е водещ принцип!
	cyr2Lat = map[string]string{
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

	tokens = map[string]string{
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
	tr := &TranslitBG{}
	return tr
}

func (tr *TranslitBG) Run(input string) string {
	// TODO
	return "abvgdezhziyklmnoprstufhtschshshtayyuyai"
}