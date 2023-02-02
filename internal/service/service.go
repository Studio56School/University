package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/Studio56School/university/internal/config"
	"github.com/Studio56School/university/internal/model"
	"github.com/Studio56School/university/internal/storage"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"time"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
}

const (
	salt       = "dsfmwemf234iomnr3u49u5mn2klfmklwecmw"
	signingKey = "lwqeclmclksmc#noldfnveoovrmv"
	tokenTTL   = 12 * time.Hour
)

type IService interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	conf   *config.Config
	logger *zap.Logger
	urepo  *storage.Repo
}

func NewService(conf *config.Config, logger *zap.Logger, urepo *storage.Repo) *Service {
	return &Service{conf: conf, logger: logger, urepo: urepo}
}

func (s *Service) CreateUser(user model.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.urepo.CreateUser(user)
}

func (s *Service) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))

}

func (s *Service) GenerateToken(username, password string) (string, error) {

	user, err := s.urepo.GetUser(username, s.generatePasswordHash(password))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		user.Name,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *Service) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)

	if !ok {
		return 0, errors.New("token claims are not the type of token.Claims.(*tokenClaims) ")
	}

	return claims.UserId, nil

}
