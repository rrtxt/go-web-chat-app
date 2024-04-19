package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"log"
	"os"
)

type JWTClaims struct {
	Id string
	Name string
	Iat int64
}

func JWTEncode(jwt JWTClaims) string {
	header := `{ "alg": "HS256", "typ": "JWT" }`

	payload, err := json.Marshal(jwt)
	if err != nil {
		log.Fatalln(err)
	}

	signature_key := os.Getenv("SIGNATURE_KEY")

	headerEncoded := base64.StdEncoding.EncodeToString([]byte(header))
	payloadEncoded := base64.StdEncoding.EncodeToString(payload)

	//Merge encoded header and payload
	data := headerEncoded + "." + payloadEncoded


	//Hash the merge data with SHA256
	hmac := hmac.New(sha256.New, []byte(signature_key))
	hmac.Write([]byte(data))
	dataHmac := hmac.Sum(nil)

	//Convert hash to string
	signature := hex.EncodeToString(dataHmac) 

	token := headerEncoded + "." + payloadEncoded + "." + signature

	return token
}