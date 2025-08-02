package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	authVal := headers.Get("Authorization")
	if authVal == "" {
		return "", fmt.Errorf("authorization not found")
	}
	authSlc := strings.Fields(authVal)
	if len(authSlc) != 2 {
		return "", fmt.Errorf("couldn't parse authorization")
	}
	if authSlc[0] != "ApiKey" {
		return "", fmt.Errorf("improper authorization header")
	}
	return authSlc[1], nil
}
