package main
import (
  xor "github.com/UnityTheCoder/xor/enc"
  "fmt"
)

func main() {
  key := "QXA5F"
  str := "SecretPassword"


  encrypted := xor.encrypt(str)
  fmt.Println("Encrypted:", encrypted)

  decrypted := xor.decrypt(encrypted)
  fmt.Println("Decrypted:", decrypted)
}
