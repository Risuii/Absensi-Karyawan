package user_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/Risuii/helpers/constant"
	"github.com/Risuii/internal/user"
	"github.com/Risuii/models/users"
	"github.com/Risuii/tests/mock"
)

var currentTime = time.Date(2021, 12, 12, 0, 0, 0, 0, &time.Location{})
var employeeStruct = users.Employee{
	ID:        1,
	Name:      "test",
	Password:  "test",
	Email:     "test@test.com",
	CreatedAt: currentTime,
	UpdateAt:  currentTime,
}

func TestCreate(t *testing.T) {
	t.Run("Create Success", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := user.NewUserRepository(db, constant.TableEmployee)

		defer db.Close()

		query := fmt.Sprintf(`INSERT INTO %s`, constant.TableEmployee)
		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs(employeeStruct.Name, employeeStruct.Password, employeeStruct.Email, employeeStruct.CreatedAt).WillReturnResult(sqlmock.NewResult(1, 1))

		ID, err := repo.Create(ctx, employeeStruct)

		assert.Equal(t, int64(1), ID)
		assert.NoError(t, err)
	})

	t.Run("Create Error", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := user.NewUserRepository(db, constant.TableEmployee)

		defer db.Close()

		query := fmt.Sprintf(`INSERT INTO %s`, constant.TableEmployee)
		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs(employeeStruct.Name, employeeStruct.Password, employeeStruct.Email, employeeStruct.CreatedAt).WillReturnResult(sqlmock.NewResult(0, 0))

		ID, err := repo.Create(ctx, employeeStruct)

		assert.Equal(t, int64(0), ID)
		assert.NoError(t, err)
	})
}

func TestFindByEmail(t *testing.T) {
	t.Run("FindByEmail Success", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := user.NewUserRepository(db, constant.TableEmployee)

		defer db.Close()

		query := fmt.Sprintf(`SELECT id, name, password, email, created_at, update_at FROM %s WHERE email = ?`, constant.TableEmployee)
		rows := sqlmock.NewRows([]string{"id", "name", "password", "email", "created_at", "update_at"}).AddRow(employeeStruct.ID, employeeStruct.Name, employeeStruct.Password, employeeStruct.Email, employeeStruct.CreatedAt, employeeStruct.UpdateAt)

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectQuery().WithArgs(employeeStruct.Email).WillReturnRows(rows)

		employeeStruct, err := repo.FindByEmail(ctx, employeeStruct.Email)

		assert.NotNil(t, employeeStruct)
		assert.NoError(t, err)
	})

	t.Run("FindByEmail Error", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := user.NewUserRepository(db, constant.TableEmployee)

		defer db.Close()

		query := fmt.Sprintf(`SELECT id, name, password, email, created_at, update_at FROM %s WHERE email = ?`, constant.TableEmployee)
		rows := sqlmock.NewRows([]string{"id", "name", "password", "email", "created_at", "update_at"})

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectQuery().WithArgs(employeeStruct.Email).WillReturnRows(rows)

		employeeStruct, err := repo.FindByEmail(ctx, employeeStruct.Email)

		assert.Empty(t, employeeStruct)
		assert.Error(t, err)
	})
}
