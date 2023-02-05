package activity_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/Risuii/helpers/constant"
	"github.com/Risuii/internal/activity"
	"github.com/Risuii/models/activitys"
	"github.com/Risuii/tests/mock"
)

var currentTime = time.Date(2021, 12, 12, 0, 0, 0, 0, &time.Location{})
var activityStruct = activitys.Activity{
	ID:          1,
	UserID:      1,
	Description: "test",
	CreatedAt:   currentTime,
	UpdateAt:    currentTime,
}
var dateStruct = activitys.DateReq{
	From: "2000-01-01",
	To:   "2000-01-02",
}

func TestAddActivityRepo(t *testing.T) {
	t.Run("Add Activity Success", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := activity.NewActivityRepositoryImpl(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf(`INSERT INTO %s`, constant.TableActivity)
		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs(activityStruct.UserID, activityStruct.Description, activityStruct.CreatedAt).WillReturnResult(sqlmock.NewResult(1, 1))

		ID, err := repo.AddActivity(ctx, activityStruct.ID, activityStruct)

		assert.Equal(t, int64(1), ID)
		assert.NoError(t, err)
	})

	t.Run("Add Activity Error", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := activity.NewActivityRepositoryImpl(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf(`INSERT INTO %s`, constant.TableActivity)
		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs(activityStruct.UserID, activityStruct.Description, activityStruct.CreatedAt).WillReturnResult(sqlmock.NewResult(0, 0))

		ID, err := repo.AddActivity(ctx, activityStruct.ID, activityStruct)

		assert.Equal(t, int64(0), ID)
		assert.NoError(t, err)
	})
}

func TestUpdateActivityRepo(t *testing.T) {
	t.Run("Update Activity Success", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := activity.NewActivityRepositoryImpl(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf(`UPDATE %s SET`, constant.TableActivity)

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs(activityStruct.Description, activityStruct.UpdateAt).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.UpdateActivity(ctx, activityStruct.ID, activityStruct)

		assert.NoError(t, err)
	})

	t.Run("Update Activity Error", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := activity.NewActivityRepositoryImpl(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf(`UPDATE %s SET`, constant.TableActivity)

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs(activityStruct.Description, activityStruct.UpdateAt).WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.UpdateActivity(ctx, int64(0), activityStruct)

		assert.Error(t, err)
	})
}

func TestFindByIDActivityRepo(t *testing.T) {
	t.Run("FindByID Success", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := activity.NewActivityRepositoryImpl(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf(`SELECT id, userID, deskripsi, created_at, update_at FROM %s WHERE id = ?`, constant.TableActivity)
		rows := sqlmock.NewRows([]string{"id", "userID", "deskripsi", "created_at", "update_at"}).AddRow(activityStruct.ID, activityStruct.UserID, activityStruct.Description, activityStruct.CreatedAt, activityStruct.UpdateAt)

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectQuery().WithArgs(activityStruct.ID).WillReturnRows(rows)

		activityStruct, err := repo.FindByID(ctx, activityStruct.ID)

		assert.NotNil(t, activityStruct)
		assert.NoError(t, err)
	})

	t.Run("FindByID Error", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := activity.NewActivityRepositoryImpl(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf(`SELECT id, userID, deskripsi, created_at, update_at FROM %s WHERE id = ?`, constant.TableActivity)
		rows := sqlmock.NewRows([]string{"id", "userID", "deskripsi", "created_at", "update_at"})

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectQuery().WithArgs(activityStruct.ID).WillReturnRows(rows)

		activityStruct, err := repo.FindByID(ctx, activityStruct.ID)

		assert.Empty(t, activityStruct)
		assert.Error(t, err)
	})
}

func TestDeleteRepo(t *testing.T) {
	t.Run("Delete Success", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := activity.NewActivityRepositoryImpl(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, constant.TableActivity)

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Delete(ctx, activityStruct.ID)

		assert.NoError(t, err)
	})

	t.Run("Delete Error", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := activity.NewActivityRepositoryImpl(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, constant.TableActivity)

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs().WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.Delete(ctx, activityStruct.ID)

		assert.Error(t, err)
	})
}

func TestRiwayatRepo(t *testing.T) {
	t.Run("Riwayat All Params Not Nill", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := activity.NewActivityRepositoryImpl(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf(`SELECT id, userID, deskripsi, created_at, update_at FROM %s WHERE DATE(created_at) BETWEEN '%s' AND '%s' ORDER BY created_at asc`, constant.TableActivity, dateStruct.From, dateStruct.To)
		rows := sqlmock.NewRows([]string{"id", "userID", "deskripsi", "created_at", "update_at"}).AddRow(activityStruct.ID, activityStruct.UserID, activityStruct.Description, activityStruct.CreatedAt, activityStruct.UpdateAt)
		ctx := context.TODO()

		mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

		activityStruct, err := repo.Riwayat(ctx, activityStruct.ID, dateStruct)

		assert.NotEmpty(t, activityStruct)
		assert.NoError(t, err)
	})

	t.Run("Riwayat Params To Nill", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := activity.NewActivityRepositoryImpl(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf(`SELECT id, userID, deskripsi, created_at, update_at FROM %s WHERE DATE(created_at) BETWEEN '%s' AND '%s' ORDER BY created_at asc`, constant.TableActivity, dateStruct.From, dateStruct.To)
		rows := sqlmock.NewRows([]string{"id", "userID", "deskripsi", "created_at", "update_at"}).AddRow(activityStruct.ID, activityStruct.UserID, activityStruct.Description, activityStruct.CreatedAt, activityStruct.UpdateAt)
		ctx := context.TODO()

		mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

		activityStruct, err := repo.Riwayat(ctx, activityStruct.ID, dateStruct)

		assert.NotEmpty(t, activityStruct)
		assert.NoError(t, err)
	})

	t.Run("Riwayat Error", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := activity.NewActivityRepositoryImpl(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf(`SELECT * FROM %s WHERE DATE(created_at) BETWEEN '%s' AND '%s' ORDER BY created_at asc`, constant.TableActivity, dateStruct.From, dateStruct.To)
		rows := sqlmock.NewRows([]string{"id", "userID", "deskripsi", "created_at", "update_at"}).AddRow(activityStruct.ID, activityStruct.UserID, activityStruct.Description, activityStruct.CreatedAt, activityStruct.UpdateAt)

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectQuery().WithArgs().WillReturnRows(rows)

		_, err := repo.Riwayat(ctx, activityStruct.UserID, dateStruct)

		assert.Error(t, err)
	})
}
