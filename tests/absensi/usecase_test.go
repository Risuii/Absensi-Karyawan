package absensi_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Risuii/helpers/exception"
	"github.com/Risuii/internal/absensi"
	"github.com/Risuii/models/absensis"
	"github.com/Risuii/tests/absensi/mocks"
)

func TestCheckin(t *testing.T) {
	t.Run("Success Checkin", func(t *testing.T) {
		absensiRepository := new(mocks.AbsensiRepository)

		absensiRepository.On("SendMsg", mock.Anything, mock.AnythingOfType("absensis.Absensi")).Return(nil)
		absensiRepository.On("ReceiveMsg").Return(absensis.Absensi{}, nil)
		absensiRepository.On("Checkin", mock.Anything, mock.AnythingOfType("absensis.Absensi")).Return(int64(1), nil)

		absensiUseCase := absensi.NewAbsensiUseCase(
			absensiRepository,
		)

		ctx := context.TODO()

		params := absensis.Absensi{
			ID:       1,
			UserID:   1,
			Name:     "test",
			Checkin:  time.Time{},
			Checkout: time.Time{},
		}

		resp, _ := absensiUseCase.Checkin(ctx, params.UserID, params.Name)

		assert.NoError(t, resp.Err())
		absensiRepository.AssertExpectations(t)
	})

	t.Run("Internal Server Error Send Msg", func(t *testing.T) {
		absensiRepository := new(mocks.AbsensiRepository)

		absensiRepository.On("SendMsg", mock.Anything, mock.AnythingOfType("absensis.Absensi")).Return(exception.ErrInternalServer)

		absensiUseCase := absensi.NewAbsensiUseCase(
			absensiRepository,
		)

		ctx := context.TODO()

		params := absensis.Absensi{
			ID:       1,
			UserID:   1,
			Name:     "test",
			Checkin:  time.Time{},
			Checkout: time.Time{},
		}

		resp, _ := absensiUseCase.Checkin(ctx, params.UserID, params.Name)

		assert.Error(t, resp.Err())
		absensiRepository.AssertExpectations(t)
	})

	t.Run("Internal Server Error Receive Msg", func(t *testing.T) {
		absensiRepository := new(mocks.AbsensiRepository)

		absensiRepository.On("SendMsg", mock.Anything, mock.AnythingOfType("absensis.Absensi")).Return(nil)
		absensiRepository.On("ReceiveMsg").Return(absensis.Absensi{}, exception.ErrInternalServer)

		absensiUseCase := absensi.NewAbsensiUseCase(
			absensiRepository,
		)

		ctx := context.TODO()

		params := absensis.Absensi{
			ID:       1,
			UserID:   1,
			Name:     "test",
			Checkin:  time.Time{},
			Checkout: time.Time{},
		}

		resp, _ := absensiUseCase.Checkin(ctx, params.UserID, params.Name)

		assert.Error(t, resp.Err())
		absensiRepository.AssertExpectations(t)
	})

	t.Run("Internal Server Error Checkin", func(t *testing.T) {
		absensiRepository := new(mocks.AbsensiRepository)

		absensiRepository.On("SendMsg", mock.Anything, mock.AnythingOfType("absensis.Absensi")).Return(nil)
		absensiRepository.On("ReceiveMsg").Return(absensis.Absensi{}, nil)
		absensiRepository.On("Checkin", mock.Anything, mock.AnythingOfType("absensis.Absensi")).Return(int64(0), exception.ErrInternalServer)

		absensiUseCase := absensi.NewAbsensiUseCase(
			absensiRepository,
		)

		ctx := context.TODO()

		params := absensis.Absensi{
			ID:       1,
			UserID:   1,
			Name:     "test",
			Checkin:  time.Time{},
			Checkout: time.Time{},
		}

		resp, _ := absensiUseCase.Checkin(ctx, params.UserID, params.Name)

		assert.Error(t, resp.Err())
		absensiRepository.AssertExpectations(t)
	})
}

func TestCheckout(t *testing.T) {
	t.Run("Success Checkout", func(t *testing.T) {
		absensiRepository := new(mocks.AbsensiRepository)

		absensiRepository.On("Checkout", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("absensis.Absensi")).Return(nil)

		absensiUseCase := absensi.NewAbsensiUseCase(
			absensiRepository,
		)

		ctx := context.TODO()

		params := absensis.Absensi{
			ID:       1,
			UserID:   1,
			Name:     "test",
			Checkin:  time.Time{},
			Checkout: time.Time{},
		}

		resp := absensiUseCase.Checkout(ctx, params.ID)

		assert.NoError(t, resp.Err())
		absensiRepository.AssertExpectations(t)
	})

	t.Run("Error Not Found Checkout", func(t *testing.T) {
		absensiRepository := new(mocks.AbsensiRepository)

		absensiRepository.On("Checkout", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("absensis.Absensi")).Return(exception.ErrNotFound)

		absensiUseCase := absensi.NewAbsensiUseCase(
			absensiRepository,
		)

		ctx := context.TODO()

		params := absensis.Absensi{
			ID:       1,
			UserID:   1,
			Name:     "test",
			Checkin:  time.Time{},
			Checkout: time.Time{},
		}

		resp := absensiUseCase.Checkout(ctx, params.ID)

		assert.Error(t, resp.Err())
		absensiRepository.AssertExpectations(t)
	})

	t.Run("Internal Server Error Checkout", func(t *testing.T) {
		absensiRepository := new(mocks.AbsensiRepository)

		absensiRepository.On("Checkout", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("absensis.Absensi")).Return(exception.ErrInternalServer)

		absensiUseCase := absensi.NewAbsensiUseCase(
			absensiRepository,
		)

		ctx := context.TODO()

		params := absensis.Absensi{
			ID:       1,
			UserID:   1,
			Name:     "test",
			Checkin:  time.Time{},
			Checkout: time.Time{},
		}

		resp := absensiUseCase.Checkout(ctx, params.ID)

		assert.Error(t, resp.Err())
		absensiRepository.AssertExpectations(t)
	})
}

func TestRiwayat(t *testing.T) {
	t.Run("Get Riwayat Success", func(t *testing.T) {
		absensiRepository := new(mocks.AbsensiRepository)

		absensiRepository.On("Riwayat", mock.Anything, mock.AnythingOfType("string")).Return([]absensis.Absensi{}, nil)

		absensiUseCase := absensi.NewAbsensiUseCase(
			absensiRepository,
		)

		ctx := context.TODO()

		params := absensis.Riwayat{
			Name: "test",
		}

		resp := absensiUseCase.Riwayat(ctx, params)

		assert.NoError(t, resp.Err())
		absensiRepository.AssertExpectations(t)
	})

	t.Run("Not Found Error Riwayat", func(t *testing.T) {
		absensiRepository := new(mocks.AbsensiRepository)

		absensiRepository.On("Riwayat", mock.Anything, mock.AnythingOfType("string")).Return([]absensis.Absensi{}, exception.ErrNotFound)

		absensiUseCase := absensi.NewAbsensiUseCase(
			absensiRepository,
		)

		ctx := context.TODO()

		params := absensis.Riwayat{
			Name: "test",
		}

		resp := absensiUseCase.Riwayat(ctx, params)

		assert.Error(t, resp.Err())
		absensiRepository.AssertExpectations(t)
	})

	t.Run("Internal Server Error Riwayat", func(t *testing.T) {
		absensiRepository := new(mocks.AbsensiRepository)

		absensiRepository.On("Riwayat", mock.Anything, mock.AnythingOfType("string")).Return([]absensis.Absensi{}, exception.ErrInternalServer)

		absensiUseCase := absensi.NewAbsensiUseCase(
			absensiRepository,
		)

		ctx := context.TODO()

		params := absensis.Riwayat{
			Name: "test",
		}

		resp := absensiUseCase.Riwayat(ctx, params)

		assert.Error(t, resp.Err())
		absensiRepository.AssertExpectations(t)
	})
}
