package shared

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims はJWTのClaims構造体です。
type JWTClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateJWT はUserIDからJWTトークンを生成します。
// 技術的なJWT処理をまとめたインフラ層の関数です。
func GenerateJWT(userID string) (string, error) {
	secret := getJWTSecret()

	claims := JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ParseJWT はJWTトークンを検証してClaimsを返します。
func ParseJWT(tokenString string) (*JWTClaims, error) {
	secret := getJWTSecret()

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || claims.UserID == "" {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}

func getJWTSecret() string {
	// TODO: configパッケージから取得するようにリファクタリング
	// 環境変数JWT_SECRETから取得するのが理想
	secret := "your-secret-key-change-in-production"
	return secret
}
