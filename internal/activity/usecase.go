package activity

import (
	"context"
	"time"

	"github.com/Risuii/helpers/exception"
	"github.com/Risuii/helpers/response"
	"github.com/Risuii/models/activitys"
)

type (
	ActivityUseCase interface {
		AddActivity(ctx context.Context, userID int64, params activitys.Activity) response.Response
		UpdateActivity(ctx context.Context, id int64, userID int64, params activitys.Activity) response.Response
		DeleteActivity(ctx context.Context, id int64, userID int64) response.Response
		Riwayat(ctx context.Context, userID int64, params activitys.DateReq) response.Response
	}

	activityUseCaseImpl struct {
		repository ActivityRepository
	}
)

func NewActivityUseCaseImpl(repo ActivityRepository) ActivityUseCase {
	return &activityUseCaseImpl{
		repository: repo,
	}
}

func (au *activityUseCaseImpl) AddActivity(ctx context.Context, userID int64, params activitys.Activity) response.Response {

	activity := activitys.Activity{
		ID:          params.ID,
		UserID:      params.UserID,
		Description: params.Description,
		CreatedAt:   time.Now(),
	}

	ID, err := au.repository.AddActivity(ctx, userID, activity)
	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	activity.ID = ID
	activity.UserID = userID

	return response.Success(response.StatusOK, activity)
}

func (au *activityUseCaseImpl) UpdateActivity(ctx context.Context, id int64, userID int64, params activitys.Activity) response.Response {

	activity, err := au.repository.FindByID(ctx, id)
	if err == exception.ErrNotFound {
		return response.Error(response.StatusNotFound, exception.ErrNotFound)
	}

	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	activity.Description = params.Description
	activity.UpdateAt = time.Now()

	if activity.UserID != userID {
		return response.Error(response.StatusUnauthorized, exception.ErrUnauthorized)
	}

	if err := au.repository.UpdateActivity(ctx, id, activity); err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	return response.Success(response.StatusOK, activity)
}

func (au *activityUseCaseImpl) DeleteActivity(ctx context.Context, id int64, userID int64) response.Response {
	activity, err := au.repository.FindByID(ctx, id)
	if err == exception.ErrNotFound {
		return response.Error(response.StatusNotFound, exception.ErrNotFound)
	}

	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	if activity.UserID != userID {
		return response.Error(response.StatusUnauthorized, exception.ErrUnauthorized)
	}

	if err := au.repository.Delete(ctx, id); err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	msg := "Berhasil Menghapus Aktivitas"

	return response.Success(response.StatusOK, msg)
}

func (au *activityUseCaseImpl) Riwayat(ctx context.Context, userID int64, params activitys.DateReq) response.Response {

	activity, err := au.repository.Riwayat(ctx, userID, params)
	if err == exception.ErrNotFound {
		return response.Error(response.StatusNotFound, exception.ErrNotFound)
	}
	if err != nil {
		return response.Success(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	return response.Success(response.StatusOK, activity)
}
