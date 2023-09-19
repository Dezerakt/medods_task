package utils

import (
	"crypto"
	"errors"
	"github.com/beevik/guid"
	"github.com/golang-jwt/jwt/v5"
	"medods_task/configs"
)

var (
	ErrInvalidAuthenticationToken = errors.New("invalid authentication token")
)

func GenerateJwtToken(payloadInfo map[string]interface{}) (string, error) {
	SigningMethod := &jwt.SigningMethodHMAC{Name: "SHA512", Hash: crypto.SHA512}
	jwt.RegisterSigningMethod(SigningMethod.Alg(), func() jwt.SigningMethod {
		return SigningMethod
	})

	token := jwt.New(SigningMethod)

	claims := token.Claims.(jwt.MapClaims)
	for index, value := range payloadInfo {
		claims[index] = value
	}

	tokenString, err := token.SignedString([]byte(configs.EnvConfigObject.JwtSecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateTokenCouple(guidToken *guid.Guid) (string, string, error) {
	accessToken, err := GenerateJwtToken(map[string]interface{}{
		"guid": guidToken,
		"type": "access",
	})
	if err != nil {
		return "", "", err
	}

	refreshToken, err := GenerateJwtToken(map[string]interface{}{
		"guid": guidToken,
		"type": "refresh",
	})
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, err
}
