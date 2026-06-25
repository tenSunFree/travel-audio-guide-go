package auth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type JWTVerifier struct {
	ecPublicKey *ecdsa.PublicKey
	hmacSecret  []byte
}

type jwks struct {
	Keys []jwk `json:"keys"`
}

type jwk struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Alg string `json:"alg"`
	X   string `json:"x"`
	Y   string `json:"y"`
	N   string `json:"n"`
	E   string `json:"e"`
}

func NewJWTVerifier(secret string) *JWTVerifier {
	return &JWTVerifier{hmacSecret: []byte(secret)}
}

func NewJWTVerifierFromJWKS(jwksURL string) (*JWTVerifier, error) {
	resp, err := http.Get(jwksURL)
	if err != nil {
		return nil, fmt.Errorf("fetch jwks: %w", err)
	}
	defer resp.Body.Close()

	var keys jwks
	if err := json.NewDecoder(resp.Body).Decode(&keys); err != nil {
		return nil, fmt.Errorf("decode jwks: %w", err)
	}

	for _, key := range keys.Keys {
		if key.Alg == "ES256" && key.Kty == "EC" {
			pub, err := ecPublicKeyFromJWK(key)
			if err != nil {
				return nil, fmt.Errorf("parse ec key: %w", err)
			}
			return &JWTVerifier{ecPublicKey: pub}, nil
		}
	}

	return nil, fmt.Errorf("no ES256 key found in JWKS")
}

func (v *JWTVerifier) Verify(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		switch token.Method.Alg() {
		case "ES256":
			if v.ecPublicKey == nil {
				return nil, fmt.Errorf("no EC public key configured")
			}
			if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %s", token.Header["alg"])
			}
			return v.ecPublicKey, nil
		case "HS256":
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %s", token.Header["alg"])
			}
			return v.hmacSecret, nil
		default:
			return nil, fmt.Errorf("unsupported signing method: %s", token.Header["alg"])
		}
	})

	if err != nil {
		return nil, fmt.Errorf("parse token: %w", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}
	if claims.Subject == "" {
		return nil, fmt.Errorf("token missing sub claim")
	}
	return claims, nil
}

func ecPublicKeyFromJWK(key jwk) (*ecdsa.PublicKey, error) {
	xBytes, err := base64.RawURLEncoding.DecodeString(key.X)
	if err != nil {
		return nil, fmt.Errorf("decode x: %w", err)
	}
	yBytes, err := base64.RawURLEncoding.DecodeString(key.Y)
	if err != nil {
		return nil, fmt.Errorf("decode y: %w", err)
	}

	pub := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     new(big.Int).SetBytes(xBytes),
		Y:     new(big.Int).SetBytes(yBytes),
	}

	if _, err := x509.MarshalPKIXPublicKey(pub); err != nil {
		return nil, fmt.Errorf("invalid ec public key: %w", err)
	}

	return pub, nil
}
