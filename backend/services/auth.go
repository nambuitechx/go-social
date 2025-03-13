package services

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nambuitechx/go-social/configs"
	"github.com/nambuitechx/go-social/models"
	"github.com/nambuitechx/go-social/repositories"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = configs.GetJwtSecret()

type AuthService struct {
	UserRepository *repositories.UserRepository
}

func NewAuthService(userRepository *repositories.UserRepository) *AuthService {
	return &AuthService{ UserRepository: userRepository }
}

func (s *AuthService) Health() string {
	return "Auth service is available"
}

func (s *AuthService) Register(payload *models.CreateUserPayload) (*models.UserModel, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)

	if err != nil {
		return nil, err
	}

	payload.Password = string(hashedPassword)
	user, err := s.UserRepository.InsertUser(payload)
	return user, err
}

func (s *AuthService) Login(payload *models.CreateUserPayload) (*models.TokenInfo, error) {
	user, err := s.UserRepository.SelectUserByEmail(&payload.Email, true)

	if err != nil {
		return nil, errors.New("email or password is invalid")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))

	if err != nil {
		return nil, errors.New("email or password is invalid")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"email": user.Email,
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return nil, err
	}

	myToken := &models.TokenInfo{ Token: tokenString }

	return myToken, err
}

func ValidateToken(payload *models.TokenInfo) (*models.UserInfo, error) {
	tokenString := payload.Token
	userInfo := &models.UserInfo{}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
	
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if exp := claims["exp"].(float64); exp < float64(time.Now().Unix()) {
			return nil, errors.New("token is expired")
		}

		userInfo.ID = claims["id"].(string)
		userInfo.Email = claims["email"].(string)
	} else {
		return nil, errors.New("failed to parse token")
	}

	return userInfo, nil
}

func AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if len(authHeader) < 8 {
			c.JSON(http.StatusUnauthorized, gin.H{ "message": "Invalid authorization header"})
			c.Abort()
			return
		}

		// Token: Bearer xxxx
		tokenString := authHeader[7:]

		userInfo, err := ValidateToken(&models.TokenInfo{ Token: tokenString })

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{ "message": err.Error()})
			c.Abort()
			return
		}

		c.Set("userId", userInfo.ID)
		c.Set("userEmail", userInfo.Email)
		c.Next()
	}
}
