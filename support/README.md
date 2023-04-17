# Support

![Go](https://badgen.net/badge/Go/%3E=1.18/orange)
[![packagist Version](https://badgen.net/github/release/go-packagist/support/stable)](https://github.com/go-packagist/support/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-packagist/support)](https://pkg.go.dev/github.com/go-packagist/support)
[![codecov](https://codecov.io/gh/go-packagist/support/branch/master/graph/badge.svg?token=5TWGQ9DIRU)](https://codecov.io/gh/go-packagist/support)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-packagist/support)](https://goreportcard.com/report/github.com/go-packagist/support)
[![tests](https://github.com/go-packagist/support/actions/workflows/go.yml/badge.svg)](https://github.com/go-packagist/support/actions/workflows/go.yml)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## Installation

```bash
go get github.com/go-packagist/support
```

## Usage

```go
package main

import (
	"github.com/go-packagist/support/ints"
	"github.com/go-packagist/support/strs"
)

func main() {
	strsCase()
	intsCase()
}

func strsCase() {
	// strs
	println(strs.InArray("abc", []string{"abc", "def"}))
	println(strs.Is("ab*", "abc"))
	println(strs.Md5("abc"))
	println(strs.Strpos("aabbcc", "a"))
	println(strs.Strrpos("aabbcc", "a"))
	println(strs.Strrev("abc"))
	println(strs.Strtr("aabbcc", "a", "b"))
	println(strs.Shuffle("abc"))
	println(strs.StrPad("abc", 6, " ", strs.StrPadLeft))
	println(strs.Length("张三")) // 2
	println(strs.Strcut("abc", 0, 1))
	println(strs.Limit("abc", 1, "..."))
	println(strs.Sha1("abc"))
	println(strs.Strpos("aabbcc", "a"))
	println(strs.Strrpos("aabbcc", "a"))
	println(strs.Strrev("abc"))

	// strs Atoi
	println(strs.Atoi("1").Val())
	println(strs.Atoi("a").Err())
	println(strs.Atoi("a").IsOk())

	// Type String
	println(strs.String("abc").Is("ab*"))
	println(strs.String("abc").InArray([]string{"abc", "def"}))
	println(strs.String("abc").Md5())
	println(strs.String("abc").Sha1())
	println(strs.String("aabbcc").Strpos("a"))
	println(strs.String("aabbcc").Strrpos("a"))
	println(strs.String("abc").Strrev())
	println(strs.String("aabbcc").Strtr("a", "b"))
	println(strs.String("abc").Shuffle())
	println(strs.String("1").Atoi().Val())
	println(strs.String("a").Atoi().Err())
	println(strs.String("a").Atoi().IsOk())
	println(strs.String("abc").Bytes())
	println(strs.String("abc").StrPad(6, " ", strs.StrPadLeft))
	println(strs.String("张三").Length()) // 2
	println(strs.String("abc").Strcut(0, 1))
	println(strs.String("abc").Limit(1, "..."))

	// Type Runes
	println(strs.Runes("abc").Len())
}

func intsCase() {
	// ints
	ints.InArray(1, []int{1, 2, 3})
	ints.Itoa(1)
	ints.Max(1, 2)
	ints.Min(1, 2, 3, 4)
	ints.Range(1, 10)
	ints.Random(1, 10)
	ints.RandomString(10)
	ints.Split("1,2,3", ",")
	ints.Between(1, 2, 3)

	// Type Int
	ints.Int(1).InArray([]int{1, 2, 3})
	ints.Int(1).Itoa()
	ints.Int(1).String()
	ints.Int(1).Bytes()
	ints.Int(1).Val()
	ints.Int(1).Between(1, 2)
}
```

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.