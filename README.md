# translitbg

[![CI Build](https://github.com/petarov/translitbg/actions/workflows/build.yml/badge.svg)](https://github.com/petarov/translitbg/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/petarov/translitbg)](https://goreportcard.com/report/github.com/petarov/translitbg)

Transliteration of Bulgarian to Latin characters for Go

Транслитерация на българските букви с латински, наречена още латинска транслитерация или латинизация/романизация на българския език.

# Install

    go get github.com/petarov/translitbg

# Usage

```go
    tr := translitbg.New()
    tr.Encode("абвгдежзийклмнопрстуфхцчшщъьюя")
    // Output: abvgdezhziyklmnoprstufhtschshshtayyuya
```

# References

* [Закон за транслитерацията](http://bg.wikisource.org/wiki/%D0%97%D0%B0%D0%BA%D0%BE%D0%BD_%D0%B7%D0%B0_%D1%82%D1%80%D0%B0%D0%BD%D1%81%D0%BB%D0%B8%D1%82%D0%B5%D1%80%D0%B0%D1%86%D0%B8%D1%8F%D1%82%D0%B0)
* [Транслитерация на българските букви с латински](http://bg.wikipedia.org/wiki/%D0%A2%D1%80%D0%B0%D0%BD%D1%81%D0%BB%D0%B8%D1%82%D0%B5%D1%80%D0%B0%D1%86%D0%B8%D1%8F_%D0%BD%D0%B0_%D0%B1%D1%8A%D0%BB%D0%B3%D0%B0%D1%80%D1%81%D0%BA%D0%B8%D1%82%D0%B5_%D0%B1%D1%83%D0%BA%D0%B2%D0%B8_%D1%81_%D0%BB%D0%B0%D1%82%D0%B8%D0%BD%D1%81%D0%BA%D0%B8)
* [translitbg.js](https://github.com/petarov/translitbg.js) - JavaScript implementation

# License

[MIT](LICENSE)
