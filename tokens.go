package translitbg

var (
	// Възстановяването на оригиналната дума не е водещ принцип
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

	// uppercase cyrillic character to its uppercase latin combo equivalent
	STREAMLINED_CYR2COMBO_UC = map[rune]string{
		1046: "ZH",  // Ж
		1062: "TS",  // Ц
		1063: "CH",  // Ч
		1064: "SH",  // Ш
		1065: "SHT", // Щ
		1070: "YU",  // Ю
		1071: "YA",  // Я
	}

	// БЪЛГАРИЯ
	BULGARIA_CYR_UP = []rune{1041, 1066, 1051, 1043, 1040, 1056, 1048, 1071}
	// българия
	BULGARIA_CYR_LOW = []rune{1073, 1098, 1083, 1075, 1072, 1088, 1080, 1103}
	// BULGARIA
	BULGARIA_LAT_UP = []rune{66, 85, 76, 71, 65, 82, 73, 65}
	// bulgaria
	BULGARIA_LAT_LOW = []rune{98, 117, 108, 103, 97, 114, 105, 97}
)
