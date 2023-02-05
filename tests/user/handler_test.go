package user_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Risuii/helpers/exception"
	"github.com/Risuii/helpers/response"
	"github.com/Risuii/internal/user"
	"github.com/Risuii/models/token"
	"github.com/Risuii/models/users"
	"github.com/Risuii/tests/user/mocks"
)

func TestHandler_Register(t *testing.T) {
	t.Run("Register Success", func(t *testing.T) {
		mockData := users.Employee{
			ID:       1,
			Name:     "test",
			Password: "test",
			Email:    "test@test.com",
		}

		resp := response.Success(response.StatusCreated, users.Employee{})

		newReq, err := json.Marshal(mockData)
		if err != nil {
			t.Error(err)
			return
		}
		validate := validator.New()
		employeeUseCase := new(mocks.UserUseCase)
		employeeUseCase.On("Register", mock.Anything, mock.AnythingOfType("users.Employee")).Return(resp)

		employeeHandler := user.UserHandler{
			Validate: validate,
			UseCase:  employeeUseCase,
		}

		r := httptest.NewRequest(http.MethodPost, "/just/for/testing", bytes.NewReader(newReq))
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(employeeHandler.Register)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusCreated, rb.Status)
		assert.NotNil(t, rb.Data)
	})

	t.Run("Register Error Entity", func(t *testing.T) {
		validate := validator.New()
		employeeUseCase := new(mocks.UserUseCase)

		employeeHandler := user.UserHandler{
			Validate: validate,
			UseCase:  employeeUseCase,
		}

		r := httptest.NewRequest(http.MethodPost, "/just/for/testing", nil)
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(employeeHandler.Register)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusUnprocessableEntity, rb.Status)
		assert.Nil(t, rb.Data)
	})

	t.Run("Register Error Bad Request", func(t *testing.T) {
		type invalidReq struct {
			Data string
		}

		mockData := invalidReq{
			Data: "error",
		}

		resp := response.Error(response.StatusBadRequest, exception.ErrBadRequest)

		newReq, err := json.Marshal(mockData)
		if err != nil {
			t.Error(err)
			return
		}
		validate := validator.New()
		employeeUseCase := new(mocks.UserUseCase)
		employeeUseCase.On("Register", mock.Anything, mock.AnythingOfType("users.Employee")).Return(resp)

		employeeHandler := user.UserHandler{
			Validate: validate,
			UseCase:  employeeUseCase,
		}

		r := httptest.NewRequest(http.MethodPost, "/just/for/testing", bytes.NewReader(newReq))
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(employeeHandler.Register)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusBadRequest, rb.Status)
		assert.Nil(t, rb.Data)
	})
}

func TestHandler_Login(t *testing.T) {
	t.Run("Login Success", func(t *testing.T) {
		mockData := users.EmployeeLogin{
			Email:    "test@test.com",
			Password: "test",
		}

		resp := response.Success(response.StatusOK, users.Employee{})

		newReq, err := json.Marshal(mockData)
		if err != nil {
			t.Error(err)
			return
		}
		validate := validator.New()
		activityUseCase := new(mocks.UserUseCase)
		activityUseCase.On("Login", mock.Anything, mock.AnythingOfType("users.EmployeeLogin")).Return(resp, token.Token{})

		activityHandler := user.UserHandler{
			Validate: validate,
			UseCase:  activityUseCase,
		}

		r := httptest.NewRequest(http.MethodPost, "/just/for/testing", bytes.NewReader(newReq))
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(activityHandler.Login)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusOK, rb.Status)
		assert.NotNil(t, rb.Data)
	})

	t.Run("Login Error Entity", func(t *testing.T) {
		validate := validator.New()
		activityUseCase := new(mocks.UserUseCase)
		activityUseCase.On("Login", mock.Anything, mock.AnythingOfType("users.EmployeeLogin")).Return(nil, token.Token{})

		activityHandler := user.UserHandler{
			Validate: validate,
			UseCase:  activityUseCase,
		}

		r := httptest.NewRequest(http.MethodPost, "/just/for/testing", nil)
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(activityHandler.Login)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusUnprocessableEntity, rb.Status)
		assert.Nil(t, rb.Data)
	})

	t.Run("Login Error Bad Request", func(t *testing.T) {
		type invalidReq struct {
			Data string
		}

		mockData := invalidReq{
			Data: "error",
		}

		newReq, err := json.Marshal(mockData)
		if err != nil {
			t.Error(err)
			return
		}
		validate := validator.New()
		activityUseCase := new(mocks.UserUseCase)

		activityHandler := user.UserHandler{
			Validate: validate,
			UseCase:  activityUseCase,
		}

		r := httptest.NewRequest(http.MethodPost, "/just/for/testing", bytes.NewReader(newReq))
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(activityHandler.Login)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusBadRequest, rb.Status)
		assert.Nil(t, rb.Data)
	})
}
