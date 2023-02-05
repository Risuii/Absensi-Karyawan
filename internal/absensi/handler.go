package absensi

import (
	"encoding/json"
	"net/http"

	newJWT "github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/Risuii/config/jwt"
	"github.com/Risuii/helpers/exception"
	"github.com/Risuii/helpers/response"
	"github.com/Risuii/models/absensis"
)

type AbsensiHandler struct {
	Validate *validator.Validate
	UseCase  AbsensiUseCase
}

func NewAbsensiHandler(router *mux.Router, validate *validator.Validate, usecase AbsensiUseCase) {
	handler := &AbsensiHandler{
		Validate: validate,
		UseCase:  usecase,
	}

	api := router.PathPrefix("/account").Subrouter()

	api.HandleFunc("/checkin", handler.Checkin).Methods(http.MethodPost)
	api.HandleFunc("/checkout", handler.Checkout).Methods(http.MethodGet)
	api.HandleFunc("/riwayat", handler.Riwayat).Methods(http.MethodGet)
}

func (handler *AbsensiHandler) Checkin(w http.ResponseWriter, r *http.Request) {
	var res response.Response

	ctx := r.Context()

	c, err := r.Cookie("token")
	if err != nil {
		res = response.Error(response.StatusUnauthorized, exception.ErrUnauthorized)
		res.JSON(w)
		return
	}

	tokenString := c.Value
	claims := &jwt.JWTclaim{}

	newJWT.ParseWithClaims(tokenString, claims, func(t *newJWT.Token) (interface{}, error) {
		return jwt.JWT_KEY, nil
	})

	res, token := handler.UseCase.Checkin(ctx, claims.ID, claims.Name)

	http.SetCookie(w, &http.Cookie{
		Name:     "checkin-token",
		Path:     "/",
		Value:    token.Token,
		HttpOnly: true,
	})

	res.JSON(w)
}

func (handler *AbsensiHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var res response.Response

	ctx := r.Context()

	_, err := r.Cookie("token")
	if err != nil {
		res = response.Error(response.StatusUnauthorized, exception.ErrUnauthorized)
		res.JSON(w)
		return
	}

	c, err := r.Cookie("checkin-token")
	if err != nil {
		res = response.Error(response.StatusUnauthorized, exception.ErrUnauthorized)
		res.JSON(w)
		return
	}

	tokenString := c.Value
	claims := &jwt.JWTclaim{}

	newJWT.ParseWithClaims(tokenString, claims, func(t *newJWT.Token) (interface{}, error) {
		return jwt.JWT_KEY, nil
	})

	res = handler.UseCase.Checkout(ctx, claims.CheckinID)

	http.SetCookie(w, &http.Cookie{
		Name:     "checkin-token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	res.JSON(w)
}

func (handler *AbsensiHandler) Riwayat(w http.ResponseWriter, r *http.Request) {
	var res response.Response
	var userInput absensis.Riwayat

	ctx := r.Context()

	_, err := r.Cookie("token")
	if err != nil {
		res = response.Error(response.StatusUnauthorized, err)
		res.JSON(w)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		res = response.Error(response.StatusUnprocessableEntity, err)
		res.JSON(w)
		return
	}

	if err := handler.Validate.StructCtx(ctx, userInput); err != nil {
		res = response.Error(response.StatusBadRequest, err)
		res.JSON(w)
		return
	}

	res = handler.UseCase.Riwayat(ctx, userInput)

	res.JSON(w)
}
