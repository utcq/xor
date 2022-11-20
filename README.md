# GOXOR
Xor Implementation for Golang

**Example**
```go
package main
import (
  xor "github.com/UnityTheCoder/xor/xor"
  "fmt"
)

func main() {
  key := "dfi&*FHu"
  key2 := "fihu33"
  str := "secretpassword df"


  encrypted := xor.DoubleEncrypt(str, key, key2)
  fmt.Println("encrypted:", encrypted)

  decrypted := xor.DoubleDecrypt(encrypted, key, key2)
  fmt.Println("decrypted:", decrypted)
}
```

<br>

```go
encrypted := xor.Encrypt(str, key)
```

<br>

```go
decrypted := xor.Decrypt(encrypted, key)
```
