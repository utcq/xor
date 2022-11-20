package main
import (
  xor "github.com/UnityTheCoder/xor/xor"
  "fmt"
)

func main() {
  key := 431
  key2 := "FIHU33"
  str := "SecretPassword DF"


  encrypted := xor.DoubleEncrypt(str, key, key2)
  fmt.Println("Encrypted:", encrypted)

  decrypted := xor.DoubleDecrypt(encrypted, key, key2)
  fmt.Println("Decrypted:", decrypted)
}
