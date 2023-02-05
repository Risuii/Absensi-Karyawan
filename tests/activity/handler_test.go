package activity_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	newJWT "github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Risuii/config/jwt"
	"github.com/Risuii/helpers/exception"
	"github.com/Risuii/helpers/response"
	"github.com/Risuii/internal/activity"

	"github.com/Risuii/models/activitys"
	"github.com/Risuii/tests/activity/mocks"
)

func TestHandler_AddActivity(t *testing.T) {
	t.Run("Add Activity Success", func(t *testing.T) {
		mockData := activitys.Activity{
			ID:          1,
			UserID:      1,
			Description: "test",
			CreatedAt:   time.Time{},
			UpdateAt:    time.Time{},
		}

		mockToken := &jwt.JWTclaim{
			ID:    1,
			Email: "test@test.com",
			StandardClaims: newJWT.StandardClaims{
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().Add(time.Hour * 24 * 1).Unix(),
			},
		}

		tokenAlgo := newJWT.NewWithClaims(newJWT.SigningMethodHS256, mockToken)

		token, err := tokenAlgo.SignedString(jwt.JWT_KEY)
		if err != nil {
			t.Error(err)
			return
		}

		resp := response.Success(response.StatusOK, activitys.Activity{})

		validate := validator.New()
		activityUseCase := new(mocks.ActivityUseCase)
		activityUseCase.On("AddActivity", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("activitys.Activity")).Return(resp)

		activityHandler := activity.ActivityHandler{
			Validate: validate,
			UseCase:  activityUseCase,
		}

		newReq, err := json.Marshal(mockData)
		if err != nil {
			t.Error(err)
			return
		}

		r := httptest.NewRequest(http.MethodPost, "/just/for/testing", bytes.NewReader(newReq))
		r.AddCookie(&http.Cookie{
			Name:  "checkin-token",
			Value: token,
		})
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(activityHandler.AddActivity)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusOK, rb.Status)
		assert.NotNil(t, rb.Data)

		activityUseCase.AssertExpectations(t)
	})

	t.Run("Add Activity Error Unauthorized", func(t *testing.T) {
		mockData := activitys.Activity{
			ID:          1,
			UserID:      1,
			Description: "test",
			CreatedAt:   time.Time{},
			UpdateAt:    time.Time{},
		}

		validate := validator.New()
		activityUseCase := new(mocks.ActivityUseCase)

		newReq, err := json.Marshal(mockData)
		if err != nil {
			t.Error(err)
			return
		}

		activityHandler := activity.ActivityHandler{
			Validate: validate,
			UseCase:  activityUseCase,
		}

		r := httptest.NewRequest(http.MethodGet, "/just/for/testing", bytes.NewReader(newReq))
		r.AddCookie(&http.Cookie{
			Name:  "",
			Value: "",
		})
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(activityHandler.AddActivity)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusUnauthorized, rb.Status)
		assert.Nil(t, rb.Data)
	})

	t.Run("Add Activity Error Bad Request", func(t *testing.T) {
		type invalidReq struct {
			Data string
		}
		mockData := invalidReq{
			Data: "error",
		}

		mockToken := &jwt.JWTclaim{
			ID:    1,
			Email: "test@test.com",
			StandardClaims: newJWT.StandardClaims{
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().Add(time.Hour * 24 * 1).Unix(),
			},
		}

		tokenAlgo := newJWT.NewWithClaims(newJWT.SigningMethodHS256, mockToken)

		token, err := tokenAlgo.SignedString(jwt.JWT_KEY)
		if err != nil {
			t.Error(err)
			return
		}

		validate := validator.New()
		activityUseCase := new(mocks.ActivityUseCase)

		newReq, err := json.Marshal(mockData)
		if err != nil {
			t.Error(err)
			return
		}

		activityHandler := activity.ActivityHandler{
			Validate: validate,
			UseCase:  activityUseCase,
		}

		r := httptest.NewRequest(http.MethodGet, "/just/for/testing", bytes.NewReader(newReq))
		r.AddCookie(&http.Cookie{
			Name:  "checkin-token",
			Value: token,
		})
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(activityHandler.AddActivity)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusBadRequest, rb.Status)
		assert.Nil(t, rb.Data)
	})

	t.Run("Add Activity Error Entity", func(t *testing.T) {

		mockToken := &jwt.JWTclaim{
			ID:    1,
			Email: "test@test.com",
			StandardClaims: newJWT.StandardClaims{
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().Add(time.Hour * 24 * 1).Unix(),
			},
		}

		tokenAlgo := newJWT.NewWithClaims(newJWT.SigningMethodHS256, mockToken)

		token, err := tokenAlgo.SignedString(jwt.JWT_KEY)
		if err != nil {
			t.Error(err)
			return
		}

		validate := validator.New()
		activityUseCase := new(mocks.ActivityUseCase)

		activityHandler := activity.ActivityHandler{
			Validate: validate,
			UseCase:  activityUseCase,
		}

		r := httptest.NewRequest(http.MethodGet, "/just/for/testing", nil)
		r.AddCookie(&http.Cookie{
			Name:  "checkin-token",
			Value: token,
		})
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(activityHandler.AddActivity)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusUnprocessableEntity, rb.Status)
		assert.Nil(t, rb.Data)
	})
}

func TestHandler_UpdateActivity(t *testing.T) {
	t.Run("Update Activity Success", func(t *testing.T) {
		mockData := activitys.Activity{
			ID:          1,
			UserID:      1,
			Description: "test",
			CreatedAt:   time.Time{},
			UpdateAt:    time.Time{},
		}

		mockToken := &jwt.JWTclaim{
			ID:    1,
			Email: "test@test.com",
			StandardClaims: newJWT.StandardClaims{
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().Add(time.Hour * 24 * 1).Unix(),
			},
		}

		tokenAlgo := newJWT.NewWithClaims(newJWT.SigningMethodHS256, mockToken)

		token, err := tokenAlgo.SignedString(jwt.JWT_KEY)
		if err != nil {
			t.Error(err)
			return
		}

		resp := response.Success(response.StatusOK, activitys.Activity{})

		validate := validator.New()
		activityUseCase := new(mocks.ActivityUseCase)
		activityUseCase.On("UpdateActivity", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("int64"), mock.AnythingOfType("activitys.Activity")).Return(resp)

		newReq, err := json.Marshal(mockData)
		if err != nil {
			t.Error(err)
			return
		}

		activityHandler := activity.ActivityHandler{
			Validate: validate,
			UseCase:  activityUseCase,
		}

		r := httptest.NewRequest(http.MethodPatch, "/just/for/testing", bytes.NewReader(newReq))
		r.AddCookie(&http.Cookie{
			Name:  "checkin-token",
			Value: token,
		})
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(activityHandler.UpdateActivity)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusOK, rb.Status)
		assert.NotNil(t, rb.Data)
	})

	t.Run("Update Activity Error Unauthorized", func(t *testing.T) {
		mockData := activitys.Activity{
			ID:          1,
			UserID:      1,
			Description: "test",
			CreatedAt:   time.Time{},
			UpdateAt:    time.Time{},
		}

		resp := response.Error(response.StatusUnauthorized, exception.ErrUnauthorized)

		validate := validator.New()
		activityUseCase := new(mocks.ActivityUseCase)
		activityUseCase.On("UpdateActivity", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("int64"), mock.AnythingOfType("activitys.Activity")).Return(resp)

		newReq, err := json.Marshal(mockData)
		if err != nil {
			t.Error(err)
			return
		}

		activityHandler := activity.ActivityHandler{
			Validate: validate,
			UseCase:  activityUseCase,
		}

		r := httptest.NewRequest(http.MethodPatch, "/just/for/testing", bytes.NewReader(newReq))
		r.AddCookie(&http.Cookie{
			Name:  "",
			Value: "",
		})
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(activityHandler.UpdateActivity)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusUnauthorized, rb.Status)
		assert.Nil(t, rb.Data)
	})

	t.Run("Update Activity Error Entity", func(t *testing.T) {
		mockToken := &jwt.JWTclaim{
			ID:    1,
			Email: "test@test.com",
			StandardClaims: newJWT.StandardClaims{
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().Add(time.Hour * 24 * 1).Unix(),
			},
		}

		tokenAlgo := newJWT.NewWithClaims(newJWT.SigningMethodHS256, mockToken)

		token, err := tokenAlgo.SignedString(jwt.JWT_KEY)
		if err != nil {
			t.Error(err)
			return
		}

		resp := response.Error(response.StatusUnprocessableEntity, exception.ErrUnprocessableEntity)

		validate := validator.New()
		activityUseCase := new(mocks.ActivityUseCase)
		activityUseCase.On("UpdateActivity", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("int64"), mock.AnythingOfType("activitys.Activity")).Return(resp)

		activityHandler := activity.ActivityHandler{
			Validate: validate,
			UseCase:  activityUseCase,
		}

		r := httptest.NewRequest(http.MethodPatch, "/just/for/testing", nil)
		r.AddCookie(&http.Cookie{
			Name:  "checkin-token",
			Value: token,
		})
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(activityHandler.UpdateActivity)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusUnprocessableEntity, rb.Status)
		assert.Nil(t, rb.Data)
	})

	t.Run("Update Activity Error Bad Request", func(t *testing.T) {
		mockData := activitys.Activity{
			ID:          1,
			UserID:      1,
			Description: "",
			CreatedAt:   time.Time{},
			UpdateAt:    time.Time{},
		}

		mockToken := &jwt.JWTclaim{
			ID:    1,
			Email: "test@test.com",
			StandardClaims: newJWT.StandardClaims{
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().Add(time.Hour * 24 * 1).Unix(),
			},
		}

		tokenAlgo := newJWT.NewWithClaims(newJWT.SigningMethodHS256, mockToken)

		token, err := tokenAlgo.SignedString(jwt.JWT_KEY)
		if err != nil {
			t.Error(err)
			return
		}

		resp := response.Error(response.StatusBadRequest, exception.ErrBadRequest)

		validate := validator.New()
		activityUseCase := new(mocks.ActivityUseCase)
		activityUseCase.On("UpdateActivity", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("int64"), mock.AnythingOfType("activitys.Activity")).Return(resp)

		newReq, err := json.Marshal(mockData)
		if err != nil {
			t.Error(err)
			return
		}

		activityHandler := activity.ActivityHandler{
			Validate: validate,
			UseCase:  activityUseCase,
		}

		r := httptest.NewRequest(http.MethodPatch, "/just/for/testing", bytes.NewReader(newReq))
		r.AddCookie(&http.Cookie{
			Name:  "checkin-token",
			Value: token,
		})
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(activityHandler.UpdateActivity)
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

func TestHandler_ReadActivity(t *testing.T) {
	t.Run("Read Activity Success", func(t *testing.T) {
		mockData := activitys.DateReq{
			From: "test",
			To:   "test",
		}

		mockToken := &jwt.JWTclaim{
			ID:    1,
			Email: "test@test.com",
			StandardClaims: newJWT.StandardClaims{
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().Add(time.Hour * 24 * 1).Unix(),
			},
		}

		tokenAlgo := newJWT.NewWithClaims(newJWT.SigningMethodHS256, mockToken)

		token, err := tokenAlgo.SignedString(jwt.JWT_KEY)
		if err != nil {
			t.Error(err)
			return
		}

		resp := response.Success(response.StatusOK, []activitys.Activity{})

		newReq, _ := json.Marshal(mockData)
		activityUseCase := new(mocks.ActivityUseCase)
		activityUseCase.On("Riwayat", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("activitys.DateReq")).Return(resp)

		activityHandler := activity.ActivityHandler{
			UseCase: activityUseCase,
		}

		r := httptest.NewRequest(http.MethodGet, "/just/for/testing", bytes.NewReader(newReq))
		r.AddCookie(&http.Cookie{
			Name:  "token",
			Value: token,
		})
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(activityHandler.ReadActivity)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusOK, rb.Status)
		assert.NotNil(t, rb.Data)
	})

	t.Run("Read Activity Error Unauthorized", func(t *testing.T) {
		mockData := activitys.DateReq{
			From: "test",
			To:   "test",
		}

		newReq, _ := json.Marshal(mockData)
		activityUseCase := new(mocks.ActivityUseCase)

		activityHandler := activity.ActivityHandler{
			UseCase: activityUseCase,
		}

		r := httptest.NewRequest(http.MethodGet, "/just/for/testing", bytes.NewReader(newReq))
		r.AddCookie(&http.Cookie{
			Name:  "",
			Value: "",
		})
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(activityHandler.ReadActivity)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusUnauthorized, rb.Status)
		assert.Nil(t, rb.Data)
	})

	t.Run("Read Activity Error Entity", func(t *testing.T) {
		mockToken := &jwt.JWTclaim{
			ID:    1,
			Email: "test@test.com",
			StandardClaims: newJWT.StandardClaims{
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().Add(time.Hour * 24 * 1).Unix(),
			},
		}

		tokenAlgo := newJWT.NewWithClaims(newJWT.SigningMethodHS256, mockToken)

		token, err := tokenAlgo.SignedString(jwt.JWT_KEY)
		if err != nil {
			t.Error(err)
			return
		}

		activityUseCase := new(mocks.ActivityUseCase)

		activityHandler := activity.ActivityHandler{
			UseCase: activityUseCase,
		}

		r := httptest.NewRequest(http.MethodGet, "/just/for/testing", nil)
		r.AddCookie(&http.Cookie{
			Name:  "token",
			Value: token,
		})
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(activityHandler.ReadActivity)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusUnprocessableEntity, rb.Status)
		assert.Nil(t, rb.Data)
	})
}

func TestHandler_DeleteActivity(t *testing.T) {
	t.Run("Delete Activity Success", func(t *testing.T) {
		type Data struct {
			Msg string
		}

		mockData := Data{
			Msg: "Berhasil Delete Activity",
		}

		mockToken := &jwt.JWTclaim{
			ID:    1,
			Email: "test@test.com",
			StandardClaims: newJWT.StandardClaims{
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().Add(time.Hour * 24 * 1).Unix(),
			},
		}

		tokenAlgo := newJWT.NewWithClaims(newJWT.SigningMethodHS256, mockToken)

		token, err := tokenAlgo.SignedString(jwt.JWT_KEY)
		if err != nil {
			t.Error(err)
			return
		}

		resp := response.Success(response.StatusOK, mockData)
		activityUseCase := new(mocks.ActivityUseCase)
		activityUseCase.On("DeleteActivity", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("int64")).Return(resp)

		activityHandler := activity.ActivityHandler{
			UseCase: activityUseCase,
		}

		r := httptest.NewRequest(http.MethodDelete, "/just/for/testing", nil)
		r.AddCookie(&http.Cookie{
			Name:  "checkin-token",
			Value: token,
		})
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(activityHandler.DeleteActivity)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusOK, rb.Status)
		assert.NotNil(t, rb.Data)
	})

	t.Run("Delete Activity Error Unauthorized", func(t *testing.T) {
		activityUseCase := new(mocks.ActivityUseCase)

		activityHandler := activity.ActivityHandler{
			UseCase: activityUseCase,
		}

		r := httptest.NewRequest(http.MethodDelete, "/just/for/testing", nil)
		r.AddCookie(&http.Cookie{
			Name:  "",
			Value: "",
		})
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(activityHandler.DeleteActivity)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusUnauthorized, rb.Status)
		assert.Nil(t, rb.Data)
	})
}
