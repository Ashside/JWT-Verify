package verify

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	// 密钥，实际生产环境中应该从配置文件中读取
	secretKey = []byte("aBcDeFgHiJkLmNoPqRsTuVwXyZ0123456789+Ab=")
	// 密钥文件路径
	filePath = "config/JWT_KEY"
	// Token过期时间
	tokenExpiration = 24 * time.Hour
)

func init() {
	log.Printf("[%s] [INFO] Initializing JWT configuration",
		time.Now().Format(time.RFC3339))

	key, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("[%s] [ERROR] Failed to read JWT key file: %v",
			time.Now().Format(time.RFC3339), err)
		panic(err)
	}
	secretKey = key
	log.Printf("[%s] [INFO] Successfully loaded JWT key",
		time.Now().Format(time.RFC3339))
}

// Claims 自定义JWT声明
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT token
func GenerateToken(username string) (string, error) {
	log.Printf("[%s] [INFO] Generating JWT token for user: %s",
		time.Now().Format(time.RFC3339), username)

	// 创建声明
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	log.Printf("[%s] [INFO] Creating token with claims: %+v",
		time.Now().Format(time.RFC3339), claims)

	// 创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名token
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Printf("[%s] [ERROR] Failed to sign token: %v",
			time.Now().Format(time.RFC3339), err)
		return "", err
	}

	log.Printf("[%s] [INFO] Successfully generated token",
		time.Now().Format(time.RFC3339))
	return tokenString, nil
}

// ValidateToken 验证JWT token
func ValidateToken(tokenString string) (*Claims, error) {
	log.Printf("[%s] [INFO] Validating JWT token",
		time.Now().Format(time.RFC3339))

	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		log.Printf("[%s] [ERROR] Failed to parse token: %v",
			time.Now().Format(time.RFC3339), err)
		return nil, err
	}

	// 验证token
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		log.Printf("[%s] [INFO] Token validation successful for user: %s",
			time.Now().Format(time.RFC3339), claims.Username)
		return claims, nil
	}

	err = errors.New("invalid token")
	log.Printf("[%s] [ERROR] %v", time.Now().Format(time.RFC3339), err)
	return nil, err
}
