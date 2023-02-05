package activity_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Risuii/helpers/exception"
	"github.com/Risuii/internal/activity"
	"github.com/Risuii/models/activitys"
	"github.com/Risuii/tests/activity/mocks"
)

func TestAddActivity(t *testing.T) {
	t.Run("Add Activity Success", func(t *testing.T) {
		activityRepository := new(mocks.ActivityRepository)

		activityRepository.On("AddActivity", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("activitys.Activity")).Return(int64(1), nil)

		activityUseCase := activity.NewActivityUseCaseImpl(
			activityRepository,
		)

		ctx := context.TODO()

		params := activitys.Activity{
			ID:          1,
			UserID:      1,
			Description: "test",
			CreatedAt:   time.Time{},
			UpdateAt:    time.Time{},
		}

		resp := activityUseCase.AddActivity(ctx, params.UserID, params)

		assert.NoError(t, resp.Err())
		activityRepository.AssertExpectations(t)
	})

	t.Run("Add Activity Error", func(t *testing.T) {
		activityRepository := new(mocks.ActivityRepository)

		activityRepository.On("AddActivity", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("activitys.Activity")).Return(int64(0), exception.ErrInternalServer)

		activityUseCase := activity.NewActivityUseCaseImpl(
			activityRepository,
		)

		ctx := context.TODO()

		params := activitys.Activity{}

		resp := activityUseCase.AddActivity(ctx, params.UserID, params)

		assert.Error(t, resp.Err())
		activityRepository.AssertExpectations(t)
	})
}

func TestUpdateActivity(t *testing.T) {
	t.Run("Update Activity Success", func(t *testing.T) {

		mockData := activitys.Activity{
			ID:          1,
			UserID:      1,
			Description: "test",
			CreatedAt:   time.Time{},
			UpdateAt:    time.Time{},
		}

		activityRepository := new(mocks.ActivityRepository)

		activityRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockData, nil)
		activityRepository.On("UpdateActivity", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("activitys.Activity")).Return(nil)

		activityUseCase := activity.NewActivityUseCaseImpl(
			activityRepository,
		)

		ctx := context.TODO()
		params := activitys.Activity{
			ID:          1,
			UserID:      1,
			Description: "test",
			CreatedAt:   time.Time{},
			UpdateAt:    time.Time{},
		}

		if mockData.UserID != params.UserID {
			t.Error()
			return
		}

		resp := activityUseCase.UpdateActivity(ctx, params.ID, params.UserID, params)

		assert.NoError(t, resp.Err())
		activityRepository.AssertExpectations(t)

	})

	t.Run("Update Activity Error NotFound", func(t *testing.T) {

		activityRepository := new(mocks.ActivityRepository)

		activityRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int64")).Return(activitys.Activity{}, exception.ErrNotFound)

		activityUseCase := activity.NewActivityUseCaseImpl(
			activityRepository,
		)

		ctx := context.TODO()
		params := activitys.Activity{
			ID:          1,
			UserID:      1,
			Description: "test",
			CreatedAt:   time.Time{},
			UpdateAt:    time.Time{},
		}

		resp := activityUseCase.UpdateActivity(ctx, params.ID, params.UserID, params)

		assert.Error(t, resp.Err())
		activityRepository.AssertExpectations(t)
	})

	t.Run("Update Activity Error Internal Server", func(t *testing.T) {

		activityRepository := new(mocks.ActivityRepository)

		activityRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int64")).Return(activitys.Activity{}, exception.ErrInternalServer)

		activityUseCase := activity.NewActivityUseCaseImpl(
			activityRepository,
		)

		ctx := context.TODO()
		params := activitys.Activity{
			ID:          1,
			UserID:      1,
			Description: "test",
			CreatedAt:   time.Time{},
			UpdateAt:    time.Time{},
		}

		resp := activityUseCase.UpdateActivity(ctx, params.ID, params.UserID, params)

		assert.Error(t, resp.Err())
		activityRepository.AssertExpectations(t)
	})

	t.Run("Update Activity Error Unauthorized", func(t *testing.T) {
		mockData := activitys.Activity{
			ID:          2,
			UserID:      1,
			Description: "test",
			CreatedAt:   time.Time{},
			UpdateAt:    time.Time{},
		}

		activityRepository := new(mocks.ActivityRepository)

		activityRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockData, nil)
		activityRepository.On("UpdateActivity", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("activitys.Activity")).Return(nil)

		activityUseCase := activity.NewActivityUseCaseImpl(
			activityRepository,
		)

		ctx := context.TODO()
		params := activitys.Activity{
			ID:          1,
			UserID:      1,
			Description: "test",
			CreatedAt:   time.Time{},
			UpdateAt:    time.Time{},
		}

		if mockData.UserID != params.UserID {
			assert.Error(t, exception.ErrUnauthorized)
			return
		}

		resp := activityUseCase.UpdateActivity(ctx, params.ID, params.UserID, params)

		assert.NoError(t, resp.Err())
		activityRepository.AssertExpectations(t)
	})

	t.Run("Update Activity Error Internal Server Update", func(t *testing.T) {

		mockData := activitys.Activity{
			ID:          1,
			UserID:      1,
			Description: "test",
			CreatedAt:   time.Time{},
			UpdateAt:    time.Time{},
		}

		activityRepository := new(mocks.ActivityRepository)

		activityRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockData, nil)
		activityRepository.On("UpdateActivity", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("activitys.Activity")).Return(exception.ErrInternalServer)

		activityUseCase := activity.NewActivityUseCaseImpl(
			activityRepository,
		)

		ctx := context.TODO()
		params := activitys.Activity{
			ID:          1,
			UserID:      1,
			Description: "test",
			CreatedAt:   time.Time{},
			UpdateAt:    time.Time{},
		}

		if mockData.UserID != params.UserID {
			t.Error()
			return
		}

		resp := activityUseCase.UpdateActivity(ctx, params.ID, params.UserID, params)

		assert.Error(t, resp.Err())
		activityRepository.AssertExpectations(t)
	})
}

func TestDeleteActivity(t *testing.T) {
	t.Run("Delete Activity Success", func(t *testing.T) {
		mockData := activitys.Activity{
			ID:     1,
			UserID: 1,
		}
		activityRepository := new(mocks.ActivityRepository)

		activityRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockData, nil)
		activityRepository.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil)

		activityUseCase := activity.NewActivityUseCaseImpl(
			activityRepository,
		)

		ctx := context.TODO()

		params := activitys.Activity{
			ID:     1,
			UserID: 1,
		}

		if mockData.UserID != params.UserID {
			t.Error()
			return
		}

		resp := activityUseCase.DeleteActivity(ctx, params.ID, params.UserID)

		assert.NoError(t, resp.Err())
		activityRepository.AssertExpectations(t)
	})

	t.Run("Delete Activity Error Not Found", func(t *testing.T) {
		activityRepository := new(mocks.ActivityRepository)

		activityRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int64")).Return(activitys.Activity{}, exception.ErrNotFound)

		activityUseCase := activity.NewActivityUseCaseImpl(
			activityRepository,
		)

		ctx := context.TODO()

		params := activitys.Activity{
			ID:     1,
			UserID: 1,
		}

		resp := activityUseCase.DeleteActivity(ctx, params.ID, params.UserID)

		assert.Error(t, resp.Err())
		activityRepository.AssertExpectations(t)
	})

	t.Run("Delete Activity Error Internal Server", func(t *testing.T) {
		activityRepository := new(mocks.ActivityRepository)

		activityRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int64")).Return(activitys.Activity{}, exception.ErrInternalServer)

		activityUseCase := activity.NewActivityUseCaseImpl(
			activityRepository,
		)

		ctx := context.TODO()

		params := activitys.Activity{
			ID:     1,
			UserID: 1,
		}

		resp := activityUseCase.DeleteActivity(ctx, params.ID, params.UserID)

		assert.Error(t, resp.Err())
		activityRepository.AssertExpectations(t)
	})

	t.Run("Delete Activity Error Unauthorized", func(t *testing.T) {
		mockData := activitys.Activity{
			ID:     1,
			UserID: 1,
		}
		activityRepository := new(mocks.ActivityRepository)

		activityRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockData, nil)
		activityRepository.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil)

		activityUseCase := activity.NewActivityUseCaseImpl(
			activityRepository,
		)

		ctx := context.TODO()

		params := activitys.Activity{
			ID:     1,
			UserID: 2,
		}

		if mockData.UserID != params.UserID {
			assert.Error(t, exception.ErrUnauthorized)
			return
		}

		resp := activityUseCase.DeleteActivity(ctx, params.ID, params.UserID)

		assert.Error(t, resp.Err())
		activityRepository.AssertExpectations(t)
	})

	t.Run("Delete Activity Error Internal Server Delete", func(t *testing.T) {
		mockData := activitys.Activity{
			ID:     1,
			UserID: 1,
		}
		activityRepository := new(mocks.ActivityRepository)

		activityRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockData, nil)
		activityRepository.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(exception.ErrInternalServer)

		activityUseCase := activity.NewActivityUseCaseImpl(
			activityRepository,
		)

		ctx := context.TODO()

		params := activitys.Activity{
			ID:     1,
			UserID: 1,
		}

		if mockData.UserID != params.UserID {
			t.Error()
			return
		}

		resp := activityUseCase.DeleteActivity(ctx, params.ID, params.UserID)

		assert.Error(t, resp.Err())
		activityRepository.AssertExpectations(t)
	})
}

func TestRiwayatActivity(t *testing.T) {
	t.Run("Riwayat Success", func(t *testing.T) {
		activityRepository := new(mocks.ActivityRepository)

		activityRepository.On("Riwayat", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("activitys.DateReq")).Return([]activitys.Activity{}, nil)

		activityUseCase := activity.NewActivityUseCaseImpl(
			activityRepository,
		)

		params := activitys.Activity{
			UserID: 1,
		}

		data := activitys.DateReq{
			From: "test",
			To:   "test",
		}

		ctx := context.TODO()

		resp := activityUseCase.Riwayat(ctx, params.UserID, data)

		assert.NoError(t, resp.Err())
		activityRepository.AssertExpectations(t)
	})

	t.Run("Riwayat Error Not Found", func(t *testing.T) {
		activityRepository := new(mocks.ActivityRepository)

		activityRepository.On("Riwayat", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("activitys.DateReq")).Return([]activitys.Activity{}, exception.ErrNotFound)

		activityUseCase := activity.NewActivityUseCaseImpl(
			activityRepository,
		)

		params := activitys.Activity{
			UserID: 1,
		}

		data := activitys.DateReq{
			From: "test",
			To:   "test",
		}

		ctx := context.TODO()

		resp := activityUseCase.Riwayat(ctx, params.UserID, data)

		assert.Error(t, resp.Err())
		activityRepository.AssertExpectations(t)
	})

	t.Run("Riwayat Error Internal Server", func(t *testing.T) {
		activityRepository := new(mocks.ActivityRepository)

		activityRepository.On("Riwayat", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("activitys.DateReq")).Return([]activitys.Activity{}, exception.ErrInternalServer)

		activityUseCase := activity.NewActivityUseCaseImpl(
			activityRepository,
		)
		ctx := context.TODO()

		resp := activityUseCase.Riwayat(ctx, 0, activitys.DateReq{})

		assert.NoError(t, resp.Err())
		activityRepository.AssertExpectations(t)
	})
}
