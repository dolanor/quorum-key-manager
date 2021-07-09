package testutils

import (
	"crypto/rsa"
	"encoding/json"
	"strings"
	"time"

	"github.com/consensys/quorum-key-manager/pkg/tls/certificate"
	"github.com/consensys/quorum-key-manager/src/auth/authenticator/oicd"
	"github.com/golang-jwt/jwt"
)

type JWTGenerator struct {
	privateKey *rsa.PrivateKey
	claims     *oicd.ClaimsConfig
}

func NewJWTGenerator(keyPair *certificate.KeyPair, claims *oicd.ClaimsConfig) (*JWTGenerator, error) {
	cert, err := certificate.X509(keyPair)
	if err != nil {
		return nil, err
	}

	return &JWTGenerator{
		privateKey: cert.PrivateKey.(*rsa.PrivateKey),
		claims:     claims,
	}, nil
}

func (j *JWTGenerator) GenerateAccessToken(username string, groups []string, ttl time.Duration) (tokenValue string, err error) {
	sc := jwt.StandardClaims{
		Issuer:    "quorum-key-manager",
		IssuedAt:  time.Now().UTC().Unix(),
		NotBefore: time.Now().UTC().Unix(),
		Subject:   "test-token",
		ExpiresAt: time.Now().UTC().Add(ttl).Unix(),
	}
	
	bsc, _ := json.Marshal(sc)
	
	c := jwt.MapClaims{}
	_ = json.Unmarshal(bsc, &c)
	
	c[j.claims.Username] = username
	c[j.claims.Group] = strings.Join(groups, ",")

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	s, err := token.SignedString(j.privateKey)
	if err != nil {
		return "", err
	}

	return s, nil
}
