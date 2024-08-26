package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"strconv"
	"strings"
	"time"

	ginJwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserType string

const (
	ExternalUser                UserType = "external"
	ColaboradorEfetivoUser      UserType = "efetivo"
	ColaboradorTerceirizadoUser UserType = "terceirizado"
)

const (
	ExternalUserTokenPrefix            = "external-user-"
	ColaboradorEfetivoTokenPrefix      = "efetivo-user-"
	ColaboradorTerceirizadoTokenPrefix = "terceirizado-user-"
)

var EncrypterBytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

const cryptoSecret string = "abc&1*~#^2^#s0^=)^^7%b34"

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

func ComparePasswords(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

type AuthService struct {
	jwtSecret []byte
}

func NewAuthService(jwtSecret string) *AuthService {
	return &AuthService{jwtSecret: []byte(jwtSecret)}
}

// Creates a JWT with a sub (subject)
func (a *AuthService) CreateUserJwtWithSub(sub string) (string, error) {
	now := time.Now()

	claims := jwt.MapClaims{
		// private claims
		"iss": "Portal de Conhecimento SGI MS",
		"aud": "Portal de Conhecimento SGI MS",
		"sub": sub,
		"iat": now.Unix(),
		"jti": uuid.NewString(),
		// TASK: [546396] reduce me from 2 years to 1 days or a few hours
		"exp": now.Add(time.Duration(time.Hour * 24 * 365 * 2)).Unix(),
		// public claims
		"type": "bearer",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.jwtSecret)
}

type JwtUserInfo struct {
	Id   uint
	Type UserType
}

// gets the user id from the request context, this assumes the route calling
// this function uses the gin JWT middleware to validate and set the token
// on the request context
func (a *AuthService) GetUserInfoFromJwt(c *gin.Context) (JwtUserInfo, error) {
	jwtSubject, exists := ginJwt.ExtractClaims(c)["sub"]
	if !exists || jwtSubject == "" {
		return JwtUserInfo{}, errors.New("cannot get user id from JWT subject")
	}

	jwtSubjectStr, ok := jwtSubject.(string)
	if !ok {
		return JwtUserInfo{}, errors.New("jwt subject is not a valid string")
	}

	if strings.HasPrefix(jwtSubjectStr, ExternalUserTokenPrefix) {
		return getUserInfoByPrefix(jwtSubjectStr, ExternalUserTokenPrefix, ExternalUser)
	}

	if strings.HasPrefix(jwtSubjectStr, ColaboradorEfetivoTokenPrefix) {
		return getUserInfoByPrefix(jwtSubjectStr, ColaboradorEfetivoTokenPrefix, ColaboradorEfetivoUser)
	}

	if strings.HasPrefix(jwtSubjectStr, ColaboradorTerceirizadoTokenPrefix) {
		return getUserInfoByPrefix(jwtSubjectStr, ColaboradorTerceirizadoTokenPrefix, ColaboradorTerceirizadoUser)
	}

	return JwtUserInfo{}, errors.New("jwt subject is not a valid user identifier")
}

func getUserInfoByPrefix(jwtSub string, prefix string, userType UserType) (JwtUserInfo, error) {
	userId, err := strconv.ParseUint(strings.TrimPrefix(jwtSub, prefix), 10, 64)
	if err != nil {
		return JwtUserInfo{}, errors.New("failed to parse user id from JWT sub")
	}
	return JwtUserInfo{Id: uint(userId), Type: userType}, nil
}

func Encrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(cryptoSecret))
	if err != nil {
		return "", err
	}

	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, EncrypterBytes)

	cipherText := make([]byte, len(plainText))

	cfb.XORKeyStream(cipherText, plainText)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func Decrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(cryptoSecret))
	if err != nil {
		return "", err
	}

	cipherText, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}

	cfb := cipher.NewCFBDecrypter(block, EncrypterBytes)

	plainText := make([]byte, len(cipherText))

	cfb.XORKeyStream(plainText, cipherText)

	return string(plainText), nil
}
