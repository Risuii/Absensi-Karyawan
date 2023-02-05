package user

import (
	"context"
	"time"

	newJWT "github.com/dgrijalva/jwt-go"

	"github.com/Risuii/config/bcrypt"
	"github.com/Risuii/config/jwt"
	"github.com/Risuii/helpers/exception"
	"github.com/Risuii/helpers/response"
	"github.com/Risuii/models/token"
	"github.com/Risuii/models/users"
)

type (
	UserUseCase interface {
		Register(ctx context.Context, params users.Employee) response.Response
		Login(ctx context.Context, params users.EmployeeLogin) (response.Response, token.Token)
	}

	userUseCaseImpl struct {
		repository UserRepository
		bcrypt     bcrypt.Bcrypt
	}
)

func NewUserUseCase(repo UserRepository, bcrypt bcrypt.Bcrypt) UserUseCase {
	return &userUseCaseImpl{
		repository: repo,
		bcrypt:     bcrypt,
	}
}

func (uu *userUseCaseImpl) Register(ctx context.Context, params users.Employee) response.Response {

	_, err := uu.repository.FindByEmail(ctx, params.Email)
	if err == nil {
		return response.Error(response.StatusConflicted, exception.ErrConflicted)
	}

	hashedPassword, err := uu.bcrypt.HashPassword(params.Password)
	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	users := users.Employee{
		ID:        params.ID,
		Name:      params.Name,
		Password:  hashedPassword,
		Email:     params.Email,
		CreatedAt: time.Now(),
	}

	userID, err := uu.repository.Create(ctx, users)
	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	users.ID = userID
	users.Password = ""

	return response.Success(response.StatusCreated, users)
}

func (uu *userUseCaseImpl) Login(ctx context.Context, params users.EmployeeLogin) (response.Response, token.Token) {
	users, err := uu.repository.FindByEmail(ctx, params.Email)

	if err == exception.ErrNotFound {
		return response.Error(response.StatusNotFound, exception.ErrNotFound), token.Token{}
	}

	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer), token.Token{}
	}

	isPasswordValid := uu.bcrypt.ComparePasswordHash(params.Password, users.Password)

	if !isPasswordValid {
		return response.Error(response.StatusUnauthorized, err), token.Token{}
	}

	users.Password = ""

	claims := &jwt.JWTclaim{
		ID:    users.ID,
		Email: users.Email,
		Name:  users.Name,
		StandardClaims: newJWT.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 24 * 1).Unix(),
		},
	}

	tokenAlgo := newJWT.NewWithClaims(newJWT.SigningMethodHS256, claims)

	tokens, err := tokenAlgo.SignedString(jwt.JWT_KEY)
	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer), token.Token{}
	}

	newToken := token.Token{
		Token: tokens,
	}

	return response.Success(response.StatusOK, users), newToken
}
