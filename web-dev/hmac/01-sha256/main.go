package main

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func main() {
	m := make(map[string][]byte)

	fmt.Println(`Your key is "key"`)
	fmt.Println("Enter text to generate equivalent hash: ")
	reader := bufio.NewReader(os.Stdin)
	code, _ := reader.ReadString('\n')

	c := []byte(getCode(code))
	m[code] = c
	for k, v := range m {
		fmt.Println("Generated hash value for ", k, " is ", string(v))
	}
	fmt.Println("Want to check hash value? Let's go. Enter hash value: ")

	hash, _ := reader.ReadString('\n')

	for k, v := range m {
		for i := range v {
			if v[i] != hash[i] {
				return
			}
		}
		fmt.Println("That hash value corresponds to ", k)
		return
	}

}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("key"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
