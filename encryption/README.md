# encryption

## Usage

```go
package main

import (
	"fmt"
	"github.com/go-fires/fires/encryption"
)

func main() {
	e := encryption.NewEncrypter("EAFBSPAXDCIOGRUVNERQGXPYGPNKYATM")

	ciphertext, _ := e.Encrypt("test")
	plaintext, _ := e.Decrypt(ciphertext)

	fmt.Println(plaintext, ciphertext) // test 94Wpr_RCTTnDKw8u_zaqTsL8rr4=
}
```