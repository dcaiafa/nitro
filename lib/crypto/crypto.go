package crypto

import (
	"crypto/rand"
	"fmt"

	"github.com/dcaiafa/nitro"
)

var defaultCharset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generate_secret0(vm *nitro.VM, size int64, charset string) (string, error) {
	buf := make([]byte, size)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}

	if len(charset) > 0 {
		if len(charset) > 256 {
			return "", fmt.Errorf("charset cannot be larger than 256")
		}

		for i := range buf {
			buf[i] = charset[int(buf[i])%len(defaultCharset)]
		}
	}
	return string(buf), nil
}
