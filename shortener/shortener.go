// generate a UUID
package shortener

import (
	"crypto/rand"
	"encoding/base64"
	"net/url"
	"strings"
)

//wanna get DB incremental order then obfuscate it then encode it with base64

type ShortCodeURL struct {
	ShortCode string
	OriginalURL string
}
func GenerateShortCode (len int8, url string) ShortCodeURL {
	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)

	}
	if len > 10 || len < 4 {
		panic("Length of the short code should be less than 10 and greter than 4")
	}


	encoding := base64.StdEncoding.EncodeToString(b)
	encoding = strings.ReplaceAll(encoding, "/", "")
	encoding = strings.ReplaceAll(encoding, "+", "")
	return ShortCodeURL{ShortCode: encoding[:len], OriginalURL: url}


}

func ValidateURL(ur string) bool {
	_, err := url.ParseRequestURI(ur)
	return err == nil
}