package main
import (
  xor "github.com/UnityTheCoder/lib"
)

func main() {
  key := "QXA5F"
  str := "SecretPassword"


  encrypted := xor.EncryptDecrypt(str)
  fmt.Println("Encrypted:", encrypted)

  decrypted := xor.EncryptDecrypt(encrypted)
  fmt.Println("Decrypted:", decrypted)
}
