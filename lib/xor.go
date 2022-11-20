package xor
import "fmt"


func encrypt(input, key string) (output string) {
    for i := 0; i < len(input); i++ {
      output += string(input[i] ^ key[i % len(key)])
    }
    return output
}

func decrypt(input, key string) (output string) {
    for i := 0; i < len(input); i++ {
      output += string(input[i] ^ key[i % len(key)])
    }
    return output
}


