package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var jwtPrivateKey = os.Getenv("JWT_PRIVATE_KEY")

func CreateToken(id int) (string, error) {
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"expireAt": time.Now().Add(1440 * time.Minute).Nanosecond(),
	})
	rs, err := tok.SignedString([]byte(jwtPrivateKey))
	if err != nil {
		return "", errors.New("can not sign key")
	}
	return rs, nil

}
func TokenIsExpire(str string) (bool, error) {
	if err := VerifyToken(str); err != nil {
		return false, err
	} else {
		if claims, err := ParseToken(str); err != nil {
			return false, err
		} else {
			return claims["expireAt"].(float64) < float64(time.Now().Nanosecond()), nil
		}
	}
}
func VerifyToken(str string) error {
	_, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(jwtPrivateKey), nil
	})
	return err
}
func ParseToken(str string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(jwtPrivateKey), nil
	})
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return token.Claims.(jwt.MapClaims), nil
}
