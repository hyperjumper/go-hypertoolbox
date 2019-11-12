package security

import (
	"encoding/base64"
	"fmt"
)

// GenerateBasicAuthToken creates a token that compatible for HTTP basic authentication.
func GenerateBasicAuthToken(user, password string) string {
	strToB64 := fmt.Sprintf("%s:%s", user, password)
	str := base64.StdEncoding.EncodeToString([]byte(strToB64))
	return str
}