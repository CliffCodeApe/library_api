package token

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateRefreshToken(token string) (uint64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtConfig.publicKey, nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		data, valid := claims["data"].(map[string]interface{})
		if !valid {
			return 0, errors.New("invalid token")
		}

		id, valid := data["id"]
		if !valid {
			return 0, errors.New("invalid token")
		}

		return uint64(id.(float64)), nil
	}

	return 0, errors.New("invalid token")
}

func ValidateAccessToken(token string) (*UserAuthToken, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtConfig.publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		data, valid := claims["data"]
		if !valid {
			return nil, fmt.Errorf("token missing 'data' field")
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal 'data' field: %w", err)
		}

		var user UserAuthToken
		err = json.Unmarshal(jsonData, &user)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal 'data' into UserAuthToken: %w", err)
		}

		return &user, nil
	}

	return nil, fmt.Errorf("invalid token")
}
