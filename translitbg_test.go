package translitbg

import "testing"

func TestAlphabet(t *testing.T) {
	expected := "abvgdezhziyklmnoprstufhtschshshtayyuyai"
	got := New().Run("абвгдежзийклмнопрстуфхцчшщъьюяѝ")

	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
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
		output := New().Run(tc.input)

		if output != tc.expected {
			t.Errorf("For input '%s', expected '%s', got '%s'", tc.input, tc.expected, output)
		}
	}
}
