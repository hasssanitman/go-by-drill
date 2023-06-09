// Go provides built-in support for base64 encoding/decoding.

package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// Here’s the string we’ll encode/decode.
	data := "abc123!?$*&()'-=@~"

	// Go supports both standard and URL-compatible base64. Here’s
	// how to encode using the standard encoder. The encoder requires
	// a []byte so we convert our string to that type.
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// Decoding may return an error, which you can check if you don’t
	// already know the input to be well-formed.
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// This encodes/decodes using a URL-compatible base64 format.
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}

// The string encodes to slightly different values with the standard
// and URL base64 encoders (trailing + vs -) but they both decode
// to the original string as desired.
// $ go run base64_encoding/base64_encoding.go
// YWJjMTIzIT8kKiYoKSctPUB+
// abc123!?$*&()'-=@~

// YWJjMTIzIT8kKiYoKSctPUB-
// abc123!?$*&()'-=@~
