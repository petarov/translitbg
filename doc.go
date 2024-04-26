// Package translitbg is a Bulgarian-language transliteration package.
// It complies with the Bulgarian transliteration law which describes the
// "Streamlined System for the Romanization of Bulgarian" as the official
// transliteration method in the country.
//
// The API can be used in a very simple way to encode cyrillic-characters text
// to its corresponding latin-characters version.
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
//
//	tr := translitbg.New()
//	tr.Encode("Стара планина")
//	// Output: Stara planina
//
//	tr.Encode("Горна Оряховица")
//	// Output: Gorna Oryahovitsa
//
// There is an edge case where according to Bulgarian law the name "БЪЛГАРИЯ"
// is an exception to the transliteration rule and must have its "Ъ"
// transformed into a "U" instead of an "A". For example:
//
//	tr.Encode("БЪЛГАРИЯ")
//	// Output: BULGARIA
package translitbg
