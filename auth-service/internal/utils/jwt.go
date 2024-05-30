package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func readPrivateKey() *rsa.PrivateKey {
	keyData, err := os.ReadFile("private.key")
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		panic("Error decode private key")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	return key
}

func readPublicKey() *rsa.PublicKey {
	keyData, err := os.ReadFile("public.key")
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		panic("Error decode public key")
	}

	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	return pubKey

}

func CreateToken(ttl time.Duration, payload any) (string, error) {

	privateKey := readPrivateKey()

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["payload"] = payload
	claims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(privateKey)

	if err != nil {
		return "", fmt.Errorf("error create sign token: %w", err)
	}

	return token, nil
}

func VerifyToken(token string) (any, error) {
	publicKey := readPublicKey()

	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
		}

		return publicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parse token: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("error claims token")
	}

	return claims["payload"], nil
}
