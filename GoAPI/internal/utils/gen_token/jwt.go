package gen_token

import (
	"errors"
	"mind-set/internal/model"
	"mind-set/internal/utils"
	e "mind-set/internal/utils/errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserInfo struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Fcmtoken    string `json:"fcmtoken"`
	Token       string `json:"token"`
	AccessToken string `json:"access_token"`
	Level       string `json:"address"`
	Lesson      string `json:"lesson"`
	Score       int    `json:"score"`
	Star        int    `json:"star"`
}

func GetUserInfo(info any) (userInfo UserInfo) {
	userInfo, _ = info.(UserInfo)
	return
}

// Generate JWT Token
func Generate(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(utils.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

// Refresh JWT Token
func Refresh(claims jwt.Claims) (string, error) {
	return Generate(claims)
}

// Parse token
func Parse(accessToken string, claims jwt.Claims, options ...jwt.ParserOption) error {
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(utils.SecretKey), err
	}, options...)
	if err != nil {
		return err
	}

	if token.Valid {
		return nil
	}

	return e.NewBusinessError(1, "invalid token")
}

// GetAccessToken
func GetAccessToken(authorization string) (accessToken string, err error) {
	if authorization == "" {
		return "", errors.New("authorization header is missing")
	}
	if !strings.HasPrefix(authorization, "Bearer ") {
		return "", errors.New("invalid Authorization header format")
	}

	accessToken = strings.TrimPrefix(authorization, "Bearer ")
	return
}

type StudClaims struct {
	UserInfo
	jwt.RegisteredClaims
}

// new function for web students
func NewWebStudClaims(user *model.Student, expiresAt time.Time) StudClaims {
	return StudClaims{
		UserInfo: UserInfo{
			Id:          user.ID,
			Name:        user.Name,
			Gender:      user.Gender,
			Phone:       user.Phone,
			Email:       user.Email,
			Fcmtoken:    user.Fcmtoken,
			Token:       user.Token,
			AccessToken: "access_token",
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			Issuer:    utils.IsStudent,
			Subject:   utils.SubjectStudent,
		},
	}
}
