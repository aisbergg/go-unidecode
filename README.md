<a name="readme-top"></a>

# go-unidecode

[![GoDoc](https://pkg.go.dev/badge/github.com/aisbergg/go-unidecode)](https://pkg.go.dev/github.com/aisbergg/go-unidecode/pkg/unidecode)
[![GoReport](https://goreportcard.com/badge/github.com/aisbergg/go-unidecode)](https://goreportcard.com/report/github.com/aisbergg/go-unidecode)
[![Coverage Status](https://codecov.io/gh/aisbergg/go-unidecode/branch/main/graph/badge.svg)](https://codecov.io/gh/aisbergg/go-unidecode)
[![CodeQL](https://github.com/aisbergg/go-unidecode/actions/workflows/codeql.yml/badge.svg
)](https://github.com/aisbergg/go-unidecode/actions/workflows/codeql.yml)
[![License](https://img.shields.io/github/license/aisbergg/go-unidecode)](https://pkg.go.dev/github.com/aisbergg/go-unidecode)
[![LinkedIn](https://img.shields.io/badge/-LinkedIn-green.svg?logo=linkedin&colorB=555)](https://www.linkedin.com/in/andre-lehmann-97408221a/)

ASCII transliterations of Unicode text for Go. Unicode characters are mapped to ASCII characters based on their phonetic representation. E.g.: `André` ➟ `Andre`, `北京` ➟ `Bei Jing`

Inspired by [python-unidecode](https://github.com/avian2/unidecode).

<details open="open">
  <summary>Table of Contents</summary>

- [Installation](#installation)
- [Usage](#usage)
- [Benchmark](#benchmark)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)
- [Acknowledgments](#acknowledgments)

</details>


## Installation

```
go get -u github.com/aisbergg/go-unidecode
```

Install CLI tool:

```
$ go install github.com/aisbergg/go-unidecode/cmd/unidecode

$ unidecode 北京kožušček
Bei Jing kozuscek

$ cat file.txt | unidecode -e replace -r "#" -
```

<p align="right"><a href="#readme-top" alt="abc"><b>back to top ⇧</b></a></p>



## Usage

```go
package main

import (
	"fmt"
	"strings"

	"github.com/aisbergg/go-unidecode/pkg/unidecode"
)

func main() {
	//
	// General Usage
	//

	s := "abc 北京kožušček"
	d, _ := unidecode.Unidecode(s, unidecode.Ignore)
	fmt.Println(d)
	// Output: abc Bei Jing kozuscek

	s = "北京"
	b, _ := unidecode.UnidecodeBytes([]byte(s), unidecode.Ignore)
	fmt.Println(string(b))
	// Output: Bei Jing

	//
	// Error Handling
	//

	// return an error if an untransliteratable character is found
	s = "⁐"
	_, err := unidecode.Unidecode(s, unidecode.Strict)
	fmt.Println(err)
	// Output: no replacement found for character ⁐ in position 0

	// preserve untransliteratable characters
	d, _ = unidecode.Unidecode(s, unidecode.Preserve)
	fmt.Println(d)
	// Output: ⁐

	// replace untransliteratable characters with specified replacement text.
	d, _ = unidecode.Unidecode(s, unidecode.Replace, "?")
	fmt.Println(d)
	// Output: ?

	//
	// Append existing buffer to prevent allocations while unidecoding
	//

	s = "kožušček"
	buf := make([]byte, 0, len(s)+len(s)/3)
	b, _ = unidecode.Append(buf, s, unidecode.Ignore)
	fmt.Println(string(b))
	// Output: kozuscek

	//
	// Writing to an io.Writer
	//

	bld := strings.Builder{}
	w := unidecode.NewWriter(&bld, unidecode.Ignore)
	w.WriteString(s)
	fmt.Println(bld.String())
	// Output: kozuscek
}
```

<p align="right"><a href="#readme-top" alt="abc"><b>back to top ⇧</b></a></p>



## Benchmark

The source code for the benchmarks is located in the [benchmarks](./benchmarks) directory.

```plaintext
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkAisberggUnidecode-4         	   34971	     32703 ns/op	    6144 B/op	       1 allocs/op
BenchmarkAisberggUnidecodeAppend-4   	   38949	     30046 ns/op	       0 B/op	       0 allocs/op
BenchmarkAisberggUnidecodeWriter-4   	   27589	     43437 ns/op	   23981 B/op	       0 allocs/op
BenchmarkFiamUnidecode-4             	     949	   1211890 ns/op	 4305247 B/op	    2335 allocs/op
BenchmarkMozillazgUnidecode-4        	   10000	    102804 ns/op	  107960 B/op	     608 allocs/op
```

<p align="right"><a href="#readme-top" alt="abc"><b>back to top ⇧</b></a></p>



## Contributing

If you have any suggestions, want to file a bug report or want to contribute to this project in some other way, please read the [contribution guideline](CONTRIBUTING.md).

And don't forget to give this project a star 🌟! Thanks again!

<p align="right"><a href="#readme-top" alt="abc"><b>back to top ⇧</b></a></p>



## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right"><a href="#readme-top" alt="abc"><b>back to top ⇧</b></a></p>



## Contact

André Lehmann

- Email: aisberg@posteo.de
- [GitHub](https://github.com/aisbergg)
- [LinkedIn](https://www.linkedin.com/in/andre-lehmann-97408221a/)

<p align="right"><a href="#readme-top" alt="abc"><b>back to top ⇧</b></a></p>



## Acknowledgments

I needed an up-to-date and efficient library for decoding of unicode characters. I looked at [mozillazg/go-unidecode](https://github.com/mozillazg/go-unidecode), but it didn't deliver what I was searching for. Therefore I took it on my own and build my own library using the transliteration tables from the Python library [avian2/unidecode](https://github.com/avian2/unidecode). A big thanks to all you contributors of avian2/unidecode!

<p align="right"><a href="#readme-top" alt="abc"><b>back to top ⇧</b></a></p>
