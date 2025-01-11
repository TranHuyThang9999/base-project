package main

import (
	"fmt"
	"rices/common/utils"
)

func main() {
	u := utils.NewUUID()

	fmt.Println(u.GenPassWord())
}

// // Định nghĩa cấu trúc để nhận dữ liệu của các cert từ Google
// type Certs struct {
// 	Keys []struct {
// 		Kid string `json:"kid"`
// 		Kty string `json:"kty"`
// 		Use string `json:"use"`
// 		N   string `json:"n"`
// 		E   string `json:"e"`
// 	} `json:"keys"`
// }

// func getGooglePublicKeys() (*Certs, error) {
// 	resp, err := http.Get("https://www.googleapis.com/oauth2/v3/certs")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	var certs Certs
// 	err = json.NewDecoder(resp.Body).Decode(&certs)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &certs, nil
// }

// func verifyJWT(tokenStr string) (*jwt.Token, error) {
// 	// Lấy danh sách các khóa công khai của Google
// 	certs, err := getGooglePublicKeys()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Cắt JWT thành ba phần: header, payload, và signature
// 	parts := strings.Split(tokenStr, ".")
// 	if len(parts) != 3 {
// 		return nil, fmt.Errorf("invalid token")
// 	}

// 	// Giải mã phần header của JWT
// 	var header map[string]interface{}
// 	_, err = jwt.DecodeSegment(parts[0], &header)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Lấy kid (Key ID) từ header của JWT
// 	kid, ok := header["kid"].(string)
// 	if !ok {
// 		return nil, fmt.Errorf("missing kid in token header")
// 	}

// 	// Tìm khóa công khai phù hợp với kid
// 	var key *jwt.PublicKey
// 	for _, cert := range certs.Keys {
// 		if cert.Kid == kid {
// 			key = &jwt.PublicKey{
// 				N: cert.N,
// 				E: cert.E,
// 			}
// 			break
// 		}
// 	}
// 	if key == nil {
// 		return nil, fmt.Errorf("public key not found")
// 	}

// 	// Giải mã JWT token bằng khóa công khai
// 	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
// 		return key, nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return token, nil
// }

// func main() {
// 	// Giả sử đây là JWT token của Google trả về sau khi người dùng đăng nhập
// 	tokenStr := "YOUR_GOOGLE_JWT_TOKEN"

// 	// Kiểm tra và xác minh JWT token
// 	token, err := verifyJWT(tokenStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Nếu token hợp lệ, hiển thị thông tin về token
// 	fmt.Println("Token is valid!")
// 	fmt.Println("Claims:", token.Claims)
// }
