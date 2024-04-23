package translitbg

import "testing"

func TestAlphabet(t *testing.T) {
	expected := "abvgdezhziyklmnoprstufhtschshshtayyuyai"
	got, _ := New().Run("абвгдежзийклмнопрстуфхцчшщъьюяѝ")

	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}

func TestSentences(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"Всички хора се раждат свободни и равни по достойнство и права. Те са надарени с разум и съвест и следва да се отнасят помежду си в дух на братство.", "Vsichki hora se razhdat svobodni i ravni po dostoynstvo i prava. Te sa nadareni s razum i savest i sledva da se otnasyat pomezhdu si v duh na bratstvo."},
		{"Всички хора \nсе раждат свободни\n и равни по достойнство\n и права.", "Vsichki hora \nse razhdat svobodni\n i ravni po dostoynstvo\n i prava."},
		{"Ѝ може да бъде намерен и в други езици \nкато руския език и украинския език.", "I mozhe da bade nameren i v drugi ezitsi \nkato ruskia ezik i ukrainskia ezik."},
	}

	for _, tc := range testCases {
		output, _ := New().Run(tc.input)

		if output != tc.expected {
			t.Errorf("For sentence '%s', expected '%s', got '%s'", tc.input, tc.expected, output)
		}
	}
}

func TestPeopleNames(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"Самуил", "Samuil"},
		{"Синтия", "Sintia"},
		{"Марияна ИваноВа", "Mariana IvanoVa"},
		{"Явор", "Yavor"},
	}

	for _, tc := range testCases {
		output, _ := New().Run(tc.input)

		if output != tc.expected {
			t.Errorf("For name '%s', expected '%s', got '%s'", tc.input, tc.expected, output)
		}
	}
}

func TestOtherNames(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"Стара планина", "Stara planina"},
		{"Атанасовско езеро", "Atanasovsko ezero"},
		{"Централен Балкан", "Tsentralen Balkan"},
		{"София-юг", "Sofia-yug"},
		{"СофИя-юг", "SofIa-yug"},
		{"СофиЯ-юг", "SofiA-yug"},
		{"СофИЯ-ЮГ", "SofIA-YuG"},
		{"гр. София, ул. Тракия.", "gr. Sofia, ul. Trakia."},
		{"гр. СофИЯ, ул. ТракИя.", "gr. SofIA, ul. TrakIa."},
		{"Перник-север", "Pernik-sever"},
		{"Златни пясъци", "Zlatni pyasatsi"},
		{"Горна Оряховица", "Gorna Oryahovitsa"},
	}

	for _, tc := range testCases {
		output, _ := New().Run(tc.input)

		if output != tc.expected {
			t.Errorf("For other name '%s', expected '%s', got '%s'", tc.input, tc.expected, output)
		}
	}
}
