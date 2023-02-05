package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/Risuii/helpers/response"
	"github.com/Risuii/models/users"
)

type UserHandler struct {
	Validate *validator.Validate
	UseCase  UserUseCase
}

func NewUserHandler(router *mux.Router, validate *validator.Validate, usecase UserUseCase) {
	handler := &UserHandler{
		Validate: validate,
		UseCase:  usecase,
	}

	router.HandleFunc("/register", handler.Register).Methods(http.MethodPost)
	router.HandleFunc("/login", handler.Login).Methods(http.MethodPost)
	router.HandleFunc("/logout", handler.Logout).Methods(http.MethodGet)
}

func (handler *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var res response.Response
	var userInput users.Employee
	ctx := r.Context()
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		res = response.Error(response.StatusUnprocessableEntity, err)
		res.JSON(w)
		return
	}

	err := handler.Validate.StructCtx(ctx, userInput)
	if err != nil {
		res = response.Error(response.StatusBadRequest, err)
		res.JSON(w)
		return
	}

	res = handler.UseCase.Register(ctx, userInput)

	res.JSON(w)
}

func (handler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var res response.Response
	var userInput users.EmployeeLogin
	ctx := r.Context()
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		res = response.Error(response.StatusUnprocessableEntity, err)
		res.JSON(w)
		return
	}

	err := handler.Validate.StructCtx(ctx, userInput)
	if err != nil {
		res = response.Error(response.StatusBadRequest, err)
		res.JSON(w)
		return
	}

	res, token := handler.UseCase.Login(ctx, userInput)

	if token.Token == "" {
		http.SetCookie(w, &http.Cookie{
			Name:     "",
			Path:     "",
			Value:    "",
			HttpOnly: true,
			MaxAge:   -1,
		})
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token.Token,
		HttpOnly: true,
	})

	res.JSON(w)
}

func (handler *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "checkin-token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})
}
