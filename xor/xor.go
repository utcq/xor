package xor

func Encrypt(input, key string) (output string) {
    kL := len(key)
	  for i := range input {
		  output += string(input[i] ^ key[i%kL])
	  }
	  return output
}

func DoubleEncrypt(input, key1, key2 string) (output string) {
  first_layer := Encrypt(input, key1)
  second_layer := Encrypt(first_layer, key2)
  return second_layer
}

func DoubleDecrypt(input, key1, key2 string) (output string) {
  second_layer := Decrypt(input, key2)
  first_layer := Decrypt(second_layer, key1)
  return first_layer
}


func Decrypt(input, key string) (output string) {
    kL := len(key)
	  for i := range input {
		  output += string(input[i] ^ key[i%kL])
	  }
	  return output
}
