package activity

import (
	"encoding/json"
	"strconv"

	"net/http"

	newJWT "github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/Risuii/config/jwt"
	"github.com/Risuii/helpers/response"
	"github.com/Risuii/models/activitys"
)

type ActivityHandler struct {
	Validate *validator.Validate
	UseCase  ActivityUseCase
}

func NewActivityHandler(router *mux.Router, validate *validator.Validate, usecase ActivityUseCase) {
	handler := &ActivityHandler{
		Validate: validate,
		UseCase:  usecase,
	}

	api := router.PathPrefix("/account").Subrouter()

	api.HandleFunc("/activity", handler.AddActivity).Methods(http.MethodPost)
	api.HandleFunc("/activity/{id}", handler.UpdateActivity).Methods(http.MethodPatch)
	api.HandleFunc("/activity/{id}", handler.DeleteActivity).Methods(http.MethodDelete)
	api.HandleFunc("/activity/riwayat", handler.ReadActivity).Methods(http.MethodGet)
}

func (handler *ActivityHandler) AddActivity(w http.ResponseWriter, r *http.Request) {
	var res response.Response
	var userInput activitys.Activity

	c, err := r.Cookie("checkin-token")
	if err != nil {
		res = response.Error(response.StatusUnauthorized, err)
		res.JSON(w)
		return
	}

	tokenString := c.Value
	claims := &jwt.JWTclaim{}

	newJWT.ParseWithClaims(tokenString, claims, func(t *newJWT.Token) (interface{}, error) {
		return jwt.JWT_KEY, nil
	})

	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		res = response.Error(response.StatusUnprocessableEntity, err)
		res.JSON(w)
		return
	}

	err = handler.Validate.StructCtx(ctx, userInput)
	if err != nil {
		res = response.Error(response.StatusBadRequest, err)
		res.JSON(w)
		return
	}

	res = handler.UseCase.AddActivity(ctx, claims.ID, userInput)

	res.JSON(w)
}

func (handler *ActivityHandler) UpdateActivity(w http.ResponseWriter, r *http.Request) {
	var res response.Response
	var userInput activitys.Activity

	ctx := r.Context()

	c, err := r.Cookie("checkin-token")
	if err != nil {
		res = response.Error(response.StatusUnauthorized, err)
		res.JSON(w)
		return
	}

	tokenString := c.Value
	claims := &jwt.JWTclaim{}

	newJWT.ParseWithClaims(tokenString, claims, func(t *newJWT.Token) (interface{}, error) {
		return jwt.JWT_KEY, nil
	})

	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		res = response.Error(response.StatusUnprocessableEntity, err)
		res.JSON(w)
		return
	}

	err = handler.Validate.StructCtx(ctx, userInput)
	if err != nil {
		res = response.Error(response.StatusBadRequest, err)
		res.JSON(w)
		return
	}

	res = handler.UseCase.UpdateActivity(ctx, id, claims.ID, userInput)

	res.JSON(w)
}

func (handler *ActivityHandler) ReadActivity(w http.ResponseWriter, r *http.Request) {
	var res response.Response
	var userInput activitys.DateReq

	ctx := r.Context()
	c, err := r.Cookie("token")
	if err != nil {
		res = response.Error(response.StatusUnauthorized, err)
		res.JSON(w)
		return
	}

	tokenString := c.Value
	claims := &jwt.JWTclaim{}

	newJWT.ParseWithClaims(tokenString, claims, func(t *newJWT.Token) (interface{}, error) {
		return jwt.JWT_KEY, nil
	})

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		res = response.Error(response.StatusUnprocessableEntity, err)
		res.JSON(w)
		return
	}

	res = handler.UseCase.Riwayat(ctx, claims.ID, userInput)

	res.JSON(w)
}

func (handler *ActivityHandler) DeleteActivity(w http.ResponseWriter, r *http.Request) {
	var res response.Response
	ctx := r.Context()

	c, err := r.Cookie("checkin-token")
	if err != nil {
		res = response.Error(response.StatusUnauthorized, err)
		res.JSON(w)
		return
	}

	tokenString := c.Value
	claims := &jwt.JWTclaim{}

	newJWT.ParseWithClaims(tokenString, claims, func(t *newJWT.Token) (interface{}, error) {
		return jwt.JWT_KEY, nil
	})

	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	res = handler.UseCase.DeleteActivity(ctx, id, claims.ID)

	res.JSON(w)
}
