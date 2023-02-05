package absensi_test

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
	"github.com/Risuii/helpers/response"
	"github.com/Risuii/internal/absensi"
	"github.com/Risuii/models/absensis"
	"github.com/Risuii/models/token"
	"github.com/Risuii/tests/absensi/mocks"
)

func TestHandler_Checkin(t *testing.T) {
	t.Run("Checkin Success", func(t *testing.T) {
		type Data struct {
			Msg string
		}
		mockData := Data{
			Msg: "Berhasil Checkin",
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

		tokens, err := tokenAlgo.SignedString(jwt.JWT_KEY)
		if err != nil {
			t.Error(err)
			return
		}

		resp := response.Success(response.StatusOK, mockData)

		checkinUseCase := new(mocks.AbsensiUseCase)
		checkinUseCase.On("Checkin", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("string")).Return(resp, token.Token{})

		checkinHandler := absensi.AbsensiHandler{
			UseCase: checkinUseCase,
		}

		r := httptest.NewRequest(http.MethodPost, "/just/for/testing", nil)
		r.AddCookie(&http.Cookie{
			Name:  "token",
			Value: tokens,
		})
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(checkinHandler.Checkin)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusOK, rb.Status)
		assert.NotNil(t, rb.Data)

		checkinUseCase.AssertExpectations(t)
	})

	t.Run("Checkin Error Unauthorized", func(t *testing.T) {
		checkinUseCase := new(mocks.AbsensiUseCase)

		checkinHandler := absensi.AbsensiHandler{
			UseCase: checkinUseCase,
		}

		r := httptest.NewRequest(http.MethodPost, "/just/for/testing", nil)
		r.AddCookie(&http.Cookie{
			Name:  "",
			Value: "",
		})
		r.AddCookie(&http.Cookie{
			Name:  "",
			Value: "",
		})
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(checkinHandler.Checkin)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusUnauthorized, rb.Status)
		assert.Nil(t, rb.Data)

		checkinUseCase.AssertExpectations(t)
	})
}

func TestHandler_Checkout(t *testing.T) {
	t.Run("Checkout Success", func(t *testing.T) {
		type Data struct {
			Msg string
		}
		mockData := Data{
			Msg: "Berhasil Checkout",
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

		checkinUseCase := new(mocks.AbsensiUseCase)
		checkinUseCase.On("Checkout", mock.Anything, mock.AnythingOfType("int64")).Return(resp)

		checkinHandler := absensi.AbsensiHandler{
			UseCase: checkinUseCase,
		}

		r := httptest.NewRequest(http.MethodPost, "/just/for/testing", nil)
		r.AddCookie(&http.Cookie{
			Name:  "token",
			Value: token,
		})
		r.AddCookie(&http.Cookie{
			Name:  "checkin-token",
			Value: token,
		})
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(checkinHandler.Checkout)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusOK, rb.Status)
		assert.NotNil(t, rb.Data)

		checkinUseCase.AssertExpectations(t)
	})

	t.Run("Checkout Error First Token Unauthorized", func(t *testing.T) {
		checkinUseCase := new(mocks.AbsensiUseCase)

		checkinHandler := absensi.AbsensiHandler{
			UseCase: checkinUseCase,
		}

		r := httptest.NewRequest(http.MethodPost, "/just/for/testing", nil)
		r.AddCookie(&http.Cookie{
			Name:  "",
			Value: "",
		})
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(checkinHandler.Checkout)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusUnauthorized, rb.Status)
		assert.Nil(t, rb.Data)

		checkinUseCase.AssertExpectations(t)
	})
}

func TestHandler_Riwayat(t *testing.T) {
	t.Run("Get Riwayat Success", func(t *testing.T) {
		mockData := absensis.Riwayat{
			Name: "test",
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

		resp := response.Success(response.StatusOK, absensis.Absensi{})

		validate := validator.New()
		absensiUseCase := new(mocks.AbsensiUseCase)
		absensiUseCase.On("Riwayat", mock.Anything, mock.AnythingOfType("absensis.Riwayat")).Return(resp)

		absensiHandler := absensi.AbsensiHandler{
			Validate: validate,
			UseCase:  absensiUseCase,
		}

		newReq, err := json.Marshal(mockData)
		if err != nil {
			t.Error(err)
			return
		}

		r := httptest.NewRequest(http.MethodGet, "/just/for/testing", bytes.NewReader(newReq))
		r.AddCookie(&http.Cookie{
			Name:  "token",
			Value: token,
		})
		r.AddCookie(&http.Cookie{
			Name:  "checkin-token",
			Value: token,
		})
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(absensiHandler.Riwayat)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusOK, rb.Status)
		assert.NotNil(t, rb.Data)

		absensiUseCase.AssertExpectations(t)
	})

	t.Run("Get Riwayat Error Unauthorized", func(t *testing.T) {
		mockData := absensis.Riwayat{
			Name: "test",
		}

		validate := validator.New()
		absensiUseCase := new(mocks.AbsensiUseCase)

		absensiHandler := absensi.AbsensiHandler{
			Validate: validate,
			UseCase:  absensiUseCase,
		}

		newReq, err := json.Marshal(mockData)
		if err != nil {
			t.Error(err)
			return
		}

		r := httptest.NewRequest(http.MethodGet, "/just/for/testing", bytes.NewReader(newReq))
		r.AddCookie(&http.Cookie{
			Name:  "",
			Value: "",
		})
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(absensiHandler.Riwayat)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusUnauthorized, rb.Status)
		assert.Nil(t, rb.Data)

		absensiUseCase.AssertExpectations(t)
	})

	t.Run("Get Riwayat Error Bad Request", func(t *testing.T) {
		mockData := absensis.Riwayat{
			Name: "",
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
		absensiUseCase := new(mocks.AbsensiUseCase)

		absensiHandler := absensi.AbsensiHandler{
			Validate: validate,
			UseCase:  absensiUseCase,
		}

		newReq, err := json.Marshal(mockData)
		if err != nil {
			t.Error(err)
			return
		}

		r := httptest.NewRequest(http.MethodGet, "/just/for/testing", bytes.NewReader(newReq))
		r.AddCookie(&http.Cookie{
			Name:  "token",
			Value: token,
		})
		r.AddCookie(&http.Cookie{
			Name:  "checkin-token",
			Value: token,
		})
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(absensiHandler.Riwayat)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusBadRequest, rb.Status)
		assert.Nil(t, rb.Data)
	})

	t.Run("Get Riwayat Error Entity", func(t *testing.T) {
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
		absensiUseCase := new(mocks.AbsensiUseCase)

		absensiHandler := absensi.AbsensiHandler{
			Validate: validate,
			UseCase:  absensiUseCase,
		}

		r := httptest.NewRequest(http.MethodGet, "/just/for/testing", nil)
		r.AddCookie(&http.Cookie{
			Name:  "token",
			Value: token,
		})
		r.AddCookie(&http.Cookie{
			Name:  "checkin-token",
			Value: token,
		})
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(absensiHandler.Riwayat)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusUnprocessableEntity, rb.Status)
		assert.Nil(t, rb.Data)

		absensiUseCase.AssertExpectations(t)
	})
}
