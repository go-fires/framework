# encryption

[![Go Version](https://badgen.net/github/release/go-packagist/encryption/stable)](https://github.com/go-packagist/encryption/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-packagist/encryption)](https://pkg.go.dev/github.com/go-packagist/encryption)
[![codecov](https://codecov.io/gh/go-packagist/encryption/branch/master/graph/badge.svg?token=5TWGQ9DIRU)](https://codecov.io/gh/go-packagist/encryption)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-packagist/encryption)](https://goreportcard.com/report/github.com/go-packagist/encryption)
[![tests](https://github.com/go-packagist/encryption/actions/workflows/go.yml/badge.svg)](https://github.com/go-packagist/encryption/actions/workflows/go.yml)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## Installation

```bash
go get github.com/go-packagist/encryption
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/go-packagist/encryption"
)

func main() {
	e := encryption.NewEncrypter("EAFBSPAXDCIOGRUVNERQGXPYGPNKYATM")

	ciphertext, _ := e.Encrypt("test")
	plaintext, _ := e.Decrypt(ciphertext)

	fmt.Println(plaintext, ciphertext) // test 94Wpr_RCTTnDKw8u_zaqTsL8rr4=
}
```

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.