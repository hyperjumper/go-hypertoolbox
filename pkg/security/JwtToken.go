package security


import (
	"errors"
	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jwt"
	"github.com/hyperjumper/go-hypertoolbox/pkg/array"
	"gopkg.in/SermoDigital/jose.v1/jws"
	"log"
	"time"
)

func NewTokenFactory(secret []byte, accessExpire, refreshExpire time.Duration, issuer string, audiences []string) *TokenFactory {
	return &TokenFactory{
		secret:                secret,
		AccessExpireDuration:  accessExpire,
		RefreshExpireDuration: refreshExpire,
		Issuer:                issuer,
		Audience:              audiences,
	}
}

// TokenFactory holds the method for manipulating access and refresh token
type TokenFactory struct {
	secret                []byte
	AccessExpireDuration  time.Duration
	RefreshExpireDuration time.Duration
	Issuer                string
	Audience              []string
}

// MakeTokens produce access and refresh token.
func (factory *TokenFactory) MakeTokens(claims map[string]string) (string, string, error) {
	accessClaims := jws.Claims{}
	refreshClaims := jws.Claims{}

	for k, v := range claims {
		accessClaims.Set(k, v)
		refreshClaims.Set(k, v)
	}

	accessClaims.Set("TYPE", "ACCESS")
	accessClaims.SetExpiration(time.Now().Add(factory.AccessExpireDuration))
	accessClaims.SetIssuer(factory.Issuer)
	accessClaims.SetAudience(factory.Audience...)
	accessClaims.SetIssuedAt(time.Now())
	accessClaims.SetNotBefore(time.Now())

	refreshClaims.Set("TYPE", "REFRESH")
	refreshClaims.SetExpiration(time.Now().Add(factory.RefreshExpireDuration))
	refreshClaims.SetIssuer(factory.Issuer)
	refreshClaims.SetAudience(factory.Audience...)
	refreshClaims.SetIssuedAt(time.Now())
	refreshClaims.SetNotBefore(time.Now())

	jwtAccess := jws.NewJWT(accessClaims, crypto.SigningMethodHS256)
	jwtRefresh := jws.NewJWT(refreshClaims, crypto.SigningMethodHS256)

	accessByte, err := jwtAccess.Serialize(factory.secret)
	if err != nil {
		return "", "", err
	}
	refreshByte, err := jwtRefresh.Serialize(factory.secret)
	if err != nil {
		return "", "", err
	}

	return string(accessByte), string(refreshByte), nil
}

// ValidateToken will validate either access or refresh token.
// Validation mostly based on their signature and expiry.
func (factory *TokenFactory) ValidateToken(token string) (bool, jwt.Claims) {
	jwt, err := jws.ParseJWT([]byte(token))
	if err != nil {
		log.Fatal(err)
		return false, nil
	}
	if err := jwt.Validate(factory.secret, crypto.SigningMethodHS256); err != nil {
		return false, nil
	}
	claims := jwt.Claims()

	if auds, ok := claims.Audience(); !ok {
		return false, nil
	} else if !array.StringArrayEquals(factory.Audience, auds) {
		return false, nil
	}

	if iss, ok := claims.Issuer(); !ok {
		return false, nil
	} else if iss != factory.Issuer {
		return false, nil
	}

	return true, claims
}

// RefreshToken will create a new access token if the suplied refresh token is still valid
func (factory *TokenFactory) RefreshToken(refreshToken string) (string, error) {
	valid, claims := factory.ValidateToken(refreshToken)
	if !valid {
		return "", errors.New("token invalid")
	}
	if !claims.Has("TYPE") || claims.Get("TYPE") != "REFRESH" {
		return "", errors.New("not refresh token")
	}

	accessClaims := jws.Claims{}
	for k, v := range claims {
		accessClaims.Set(k, v)
	}
	accessClaims.Set("TYPE", "ACCESS")
	accessClaims.SetExpiration(time.Now().Add(factory.AccessExpireDuration))
	accessClaims.SetIssuer(factory.Issuer)
	accessClaims.SetAudience(factory.Audience...)
	accessClaims.SetIssuedAt(time.Now())
	accessClaims.SetNotBefore(time.Now())

	jwtAccess := jws.NewJWT(accessClaims, crypto.SigningMethodHS256)
	accessByte, err := jwtAccess.Serialize(factory.secret)
	if err != nil {
		return "", err
	}
	return string(accessByte), nil
}

