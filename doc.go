// Package translitbg is a Bulgarian-language transliteration package. The API
// can be used in a very simple way to encode cyrillic-characters text to its
// corresponding latin-characters version.
//
// Example:
//
//	tr := translitbg.New()
//	tr.Encode("абвгдежзийклмнопрстуфхцчшщъьюя")
//	// Output: abvgdezhziyklmnoprstufhtschshshtayyuya
//
// There are no state-related parameters stored in the translitbg object, so
// the same instance can be used to encode more text without the need to create
// new objects every time.
package translitbg
