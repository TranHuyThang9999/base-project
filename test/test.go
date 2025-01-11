package main

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

const GooglePublicKeyURL = "https://www.googleapis.com/oauth2/v3/certs"

type GoogleClaims struct {
	Iss           string `json:"iss"`
	Aud           string `json:"aud"`
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	jwt.RegisteredClaims
}

// GetGooglePublicKeys tải các khóa công khai từ Google
func GetGooglePublicKeys() (map[string]*rsa.PublicKey, error) {
	resp, err := http.Get(GooglePublicKeyURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var keyData struct {
		Keys []struct {
			Kid string `json:"kid"`
			N   string `json:"n"`
			E   string `json:"e"`
		} `json:"keys"`
	}

	if err := json.Unmarshal(body, &keyData); err != nil {
		return nil, err
	}

	keys := make(map[string]*rsa.PublicKey)
	for _, key := range keyData.Keys {
		n, err := base64.RawURLEncoding.DecodeString(key.N)
		if err != nil {
			return nil, err
		}

		e, err := base64.RawURLEncoding.DecodeString(key.E)
		if err != nil {
			return nil, err
		}

		keys[key.Kid] = &rsa.PublicKey{
			N: new(big.Int).SetBytes(n),
			E: int(new(big.Int).SetBytes(e).Int64()),
		}
	}

	return keys, nil
}

// VerifyGoogleToken xác minh và giải mã token của Google
func VerifyGoogleToken(idToken string) (*GoogleClaims, error) {
	publicKeys, err := GetGooglePublicKeys()
	if err != nil {
		return nil, fmt.Errorf("error getting Google public keys: %v", err)
	}

	token, err := jwt.ParseWithClaims(idToken, &GoogleClaims{}, func(token *jwt.Token) (interface{}, error) {
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, errors.New("kid header not found")
		}

		key, ok := publicKeys[kid]
		if !ok {
			return nil, errors.New("key not found for the given kid")
		}

		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	claims, ok := token.Claims.(*GoogleClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if !strings.HasPrefix(claims.Iss, "https://accounts.google.com") {
		return nil, errors.New("invalid issuer")
	}

	return claims, nil
}

func main() {
	// Thay idToken bằng token thực tế nhận được từ Google
	idToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6Ijg5Y2UzNTk4YzQ3M2FmMWJkYTRiZmY5NWU2Yzg3MzY0NTAyMDZmYmEiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiIxMDU1MDgxODg0NTQ2LW81amFrMG1pcGZyaHV1bTUzNm1qaWQ5bzkwazQ3cDQ5LmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwiYXVkIjoiMTA1NTA4MTg4NDU0Ni1vNWphazBtaXBmcmh1dW01MzZtamlkOW85MGs0N3A0OS5hcHBzLmdvb2dsZXVzZXJjb250ZW50LmNvbSIsInN1YiI6IjEwMDI0MDg1Nzg2MTQ4OTE0OTQxMCIsImVtYWlsIjoidHJhbmh1eXRoYW5nOTk5OUBnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmJmIjoxNzM2NTgwOTcxLCJuYW1lIjoiVGjhuq9uZyBUcuG6p24gSHV5IiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hL0FDZzhvY0xiTk5xWkNfRXhGUFU2NE13U3l4b1hqTGNyaWFsejZCTEZwZjhad09LZjJZZHhHSEtUPXM5Ni1jIiwiZ2l2ZW5fbmFtZSI6IlRo4bqvbmciLCJmYW1pbHlfbmFtZSI6IlRy4bqnbiBIdXkiLCJpYXQiOjE3MzY1ODEyNzEsImV4cCI6MTczNjU4NDg3MSwianRpIjoiYjcxZDJmMmY3YmZjNGQ2NDk5NDcwZWFhYzRiYTAxMDViMGU2YmExMSJ9.HcxxbhJpBZBXqtmFRYHLcvakfGdEc_dkeqHFdb16GX-rquUwLpY8VyHy4PkbYDaHJgdVFJB6pOwiYwr4VmgF4vrVUBer8YlITI6ZW0QMElUEDfVgJWWay4POJNeARMZ14rhQj2o7KPWbXcO_9c2ZL6abvoULI8aui6_9_ZXygc2dwUkay_-xRsA0jzN-xbO9oGpkBVqDd-zYa1eS1a_Uc2KRU6r99RQ2U-_oxo8ckfwmQ5H4ZKjJGEiNRB9t1--_H9ZpKQnwyDdmC-5boTK2o-l9IpLZx0lHpl0AqktNnaqvj3KhaF9rtPwRoeoE8oDiHvM6yorhjuFk8zqn8Z8GSQ"

	claims, err := VerifyGoogleToken(idToken)
	if err != nil {
		fmt.Println("Error verifying token:", err)
		return
	}
	json.Unmarshal([]byte(&claims), GoogleClaims)
	fmt.Printf("Token claims: %+v\n", claims)
}
