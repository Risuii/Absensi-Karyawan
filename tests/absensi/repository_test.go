package absensi_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/Risuii/helpers/constant"
	"github.com/Risuii/internal/absensi"
	"github.com/Risuii/models/absensis"
	"github.com/Risuii/tests/mock"
)

var currentTime = time.Date(2021, 12, 12, 0, 0, 0, 0, &time.Location{})
var absensiStruct = absensis.Absensi{
	ID:       1,
	UserID:   1,
	Name:     "test",
	Checkin:  currentTime,
	Checkout: currentTime,
}

func TestCheckinRepo(t *testing.T) {
	t.Run("Create Checkin Success", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := absensi.NewAbsensiRepositoryImpl(db, constant.TableAbsensi)

		defer db.Close()

		query := fmt.Sprintf(`INSERT INTO %s`, constant.TableAbsensi)
		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs(absensiStruct.UserID, absensiStruct.Name, absensiStruct.Checkin).WillReturnResult(sqlmock.NewResult(1, 1))

		ID, err := repo.Checkin(ctx, absensiStruct)

		assert.Equal(t, int64(1), ID)
		assert.NoError(t, err)
	})

	t.Run("Create Checkin Error", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := absensi.NewAbsensiRepositoryImpl(db, constant.TableAbsensi)

		defer db.Close()

		query := fmt.Sprintf(`INSERT INTO %s`, constant.TableAbsensi)
		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs(absensiStruct.UserID, absensiStruct.Name, absensiStruct.Checkin).WillReturnResult(sqlmock.NewResult(0, 0))

		ID, err := repo.Checkin(ctx, absensiStruct)

		assert.Equal(t, int64(0), ID)
		assert.NoError(t, err)
	})
}

func TestCheckoutRepo(t *testing.T) {
	t.Run("Update Checkout Success", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := absensi.NewAbsensiRepositoryImpl(db, constant.TableAbsensi)

		defer db.Close()

		query := fmt.Sprintf(`UPDATE %s SET`, constant.TableAbsensi)

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs(absensiStruct.Checkout).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.Checkout(ctx, absensiStruct.ID, absensiStruct)

		assert.NoError(t, err)
	})

	t.Run("Update Checkout Error", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := absensi.NewAbsensiRepositoryImpl(db, constant.TableAbsensi)

		defer db.Close()

		query := fmt.Sprintf(`UPDATE %s SET`, constant.TableAbsensi)

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs(absensiStruct.Checkout).WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.Checkout(ctx, absensiStruct.ID, absensiStruct)

		assert.Error(t, err)
	})
}

func TestRiwayatRepo(t *testing.T) {
	t.Run("Test Riwayat Success", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := absensi.NewAbsensiRepositoryImpl(db, constant.TableAbsensi)

		defer db.Close()

		query := fmt.Sprintf(`SELECT id, userID, name, checkin, checkout FROM %s WHERE name = '%s'`, constant.TableAbsensi, absensiStruct.Name)
		rows := sqlmock.NewRows([]string{"id", "userID", "name", "checkin", "checkout"}).AddRow(absensiStruct.ID, absensiStruct.UserID, absensiStruct.Name, absensiStruct.Checkin, absensiStruct.Checkout)

		ctx := context.TODO()

		mock.ExpectQuery(query).WillReturnRows(rows)

		absensiStruct, err := repo.Riwayat(ctx, absensiStruct.Name)

		assert.NotNil(t, absensiStruct)
		assert.NoError(t, err)
	})

	t.Run("Test Riwayat Error", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := absensi.NewAbsensiRepositoryImpl(db, constant.TableAbsensi)

		defer db.Close()

		query := fmt.Sprintf(`SELECT id, userID, name, checkin, checkout FROM %s WHERE name = '%s'`, constant.TableAbsensi, absensiStruct.Name)
		rows := sqlmock.NewRows([]string{"id", "userID", "name", "checkin", "checkout"})

		ctx := context.TODO()

		mock.ExpectQuery(query).WillReturnRows(rows)

		absensiStruct, err := repo.Riwayat(ctx, absensiStruct.Name)

		assert.Empty(t, absensiStruct)
		assert.NoError(t, err)
	})
}
