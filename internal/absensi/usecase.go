package absensi

import (
	"context"
	"time"

	newJWT "github.com/dgrijalva/jwt-go"

	"github.com/Risuii/config/jwt"
	"github.com/Risuii/helpers/exception"
	"github.com/Risuii/helpers/response"
	"github.com/Risuii/models/absensis"
	"github.com/Risuii/models/token"
)

type (
	AbsensiUseCase interface {
		Checkin(ctx context.Context, userID int64, name string) (response.Response, token.Token)
		Checkout(ctx context.Context, checkinID int64) response.Response
		Riwayat(ctx context.Context, params absensis.Riwayat) response.Response
	}

	absensiUseCaseImpl struct {
		repository AbsensiRepository
	}
)

func NewAbsensiUseCase(repo AbsensiRepository) AbsensiUseCase {
	return &absensiUseCaseImpl{
		repository: repo,
	}
}

func (au *absensiUseCaseImpl) Checkin(ctx context.Context, userID int64, name string) (response.Response, token.Token) {
	checkin := absensis.Absensi{
		UserID:  userID,
		Name:    name,
		Checkin: time.Now(),
	}

	err := au.repository.SendMsg(ctx, checkin)

	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer), token.Token{}
	}

	data, err := au.repository.ReceiveMsg()
	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer), token.Token{}
	}

	ID, err := au.repository.Checkin(ctx, data)
	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer), token.Token{}
	}

	claims := &jwt.JWTclaim{
		ID:        userID,
		CheckinID: ID,
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

	msg := "Success Publish Message"

	return response.Success(response.StatusOK, msg), newToken
}

func (au *absensiUseCaseImpl) Checkout(ctx context.Context, checkinID int64) response.Response {
	checkin := absensis.Absensi{
		Checkout: time.Now(),
	}

	err := au.repository.Checkout(ctx, checkinID, checkin)
	if err == exception.ErrNotFound {
		return response.Error(response.StatusNotFound, exception.ErrNotFound)
	}
	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	msg := "Berhasil Checkout"

	return response.Success(response.StatusOK, msg)
}

func (au *absensiUseCaseImpl) Riwayat(ctx context.Context, params absensis.Riwayat) response.Response {

	absensi, err := au.repository.Riwayat(ctx, params.Name)

	if err == exception.ErrNotFound {
		return response.Error(response.StatusNotFound, exception.ErrNotFound)
	}

	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	return response.Success(response.StatusOK, absensi)
}
