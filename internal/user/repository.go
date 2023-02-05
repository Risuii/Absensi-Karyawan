package user

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Risuii/helpers/exception"
	"github.com/Risuii/models/users"
)

type (
	UserRepository interface {
		Create(ctx context.Context, params users.Employee) (int64, error)
		FindByEmail(ctx context.Context, params string) (users.Employee, error)
	}

	userRepositoryImpl struct {
		db        *sql.DB
		tableName string
	}
)

func NewUserRepository(db *sql.DB, tableName string) UserRepository {
	return &userRepositoryImpl{
		db:        db,
		tableName: tableName,
	}
}

func (ur *userRepositoryImpl) Create(ctx context.Context, params users.Employee) (int64, error) {
	query := fmt.Sprintf(`INSERT INTO %s (name, password, email, created_at) VALUES (?,?,?,?)`, ur.tableName)
	stmt, err := ur.db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return 0, exception.ErrInternalServer
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		params.Name,
		params.Password,
		params.Email,
		params.CreatedAt,
	)
	if err != nil {
		log.Println(err)
		return 0, exception.ErrInternalServer
	}
	ID, _ := result.LastInsertId()
	return ID, nil
}

func (ur *userRepositoryImpl) FindByEmail(ctx context.Context, params string) (users.Employee, error) {
	var users users.Employee
	query := fmt.Sprintf(`SELECT id, name, password, email, created_at, update_at FROM %s WHERE email = ?`, ur.tableName)
	stmt, err := ur.db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return users, exception.ErrInternalServer
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, params)

	err = row.Scan(
		&users.ID,
		&users.Name,
		&users.Password,
		&users.Email,
		&users.CreatedAt,
		&users.UpdateAt,
	)
	if err != nil {
		log.Println(err)
		return users, exception.ErrNotFound
	}
	return users, nil
}
