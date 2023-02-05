package activity

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Risuii/helpers/exception"
	"github.com/Risuii/models/activitys"
)

type (
	ActivityRepository interface {
		AddActivity(ctx context.Context, userID int64, params activitys.Activity) (int64, error)
		UpdateActivity(ctx context.Context, id int64, params activitys.Activity) error
		FindByID(ctx context.Context, id int64) (activitys.Activity, error)
		Delete(ctx context.Context, id int64) error
		Riwayat(ctx context.Context, userID int64, params activitys.DateReq) ([]activitys.Activity, error)
	}

	activityRepositoryImpl struct {
		DB        *sql.DB
		TableName string
	}
)

func NewActivityRepositoryImpl(db *sql.DB, tableName string) ActivityRepository {
	return &activityRepositoryImpl{
		DB:        db,
		TableName: tableName,
	}
}

func (ar *activityRepositoryImpl) AddActivity(ctx context.Context, userID int64, params activitys.Activity) (int64, error) {
	query := fmt.Sprintf(`INSERT INTO %s (userID, deskripsi, created_at) VALUES (?, ?, ?)`, ar.TableName)
	stmt, err := ar.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return 0, exception.ErrInternalServer
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		userID,
		params.Description,
		params.CreatedAt,
	)

	if err != nil {
		log.Println(err)
		return 0, exception.ErrInternalServer
	}

	ID, err := result.LastInsertId()

	if err != nil {
		log.Println(err)
		return 0, exception.ErrInternalServer
	}

	return ID, nil
}

func (ar *activityRepositoryImpl) UpdateActivity(ctx context.Context, id int64, params activitys.Activity) error {
	query := fmt.Sprintf(`UPDATE %s SET deskripsi = ?, update_at = ? WHERE id = %d`, ar.TableName, id)
	stmt, err := ar.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return exception.ErrInternalServer
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		params.Description,
		params.UpdateAt,
	)

	if err != nil {
		log.Println(err)
		return exception.ErrInternalServer
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected < 1 {
		return exception.ErrNotFound
	}

	return nil
}

func (ar *activityRepositoryImpl) FindByID(ctx context.Context, id int64) (activitys.Activity, error) {
	activity := activitys.Activity{}

	query := fmt.Sprintf(`SELECT id, userID, deskripsi, created_at, update_at FROM %s WHERE id = ?`, ar.TableName)
	stmt, err := ar.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return activitys.Activity{}, exception.ErrInternalServer
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)

	err = row.Scan(
		&activity.ID,
		&activity.UserID,
		&activity.Description,
		&activity.CreatedAt,
		&activity.UpdateAt,
	)

	if err != nil {
		log.Println(err)
		return activitys.Activity{}, exception.ErrInternalServer
	}

	return activity, nil
}

func (ar *activityRepositoryImpl) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = %d`, ar.TableName, id)
	stmt, err := ar.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return exception.ErrInternalServer
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
	)
	if err != nil {
		log.Println(err)
		return exception.ErrInternalServer
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected < 1 {
		return exception.ErrNotFound
	}

	return nil
}

func (ar *activityRepositoryImpl) Riwayat(ctx context.Context, userID int64, params activitys.DateReq) ([]activitys.Activity, error) {
	activity := []activitys.Activity{}

	if params.To == "" {
		now := time.Now()
		rows, err := ar.DB.Query(fmt.Sprintf(`SELECT id, userID, deskripsi, created_at, update_at FROM '%s' WHERE DATE(created_at) BETWEEN '%s' AND '%s' ORDER BY created_at asc`, ar.TableName, params.From, now.Format("2006-01-02")))
		if err != nil {
			log.Println(err)
			return activity, exception.ErrInternalServer
		}

		defer rows.Close()

		for rows.Next() {
			var c activitys.Activity
			if err := rows.Scan(
				&c.ID,
				&c.UserID,
				&c.Description,
				&c.CreatedAt,
				&c.UpdateAt,
			); err != nil {
				log.Println(err)
				return activity, exception.ErrInternalServer
			}
			activity = append(activity, c)
		}

		if err = rows.Err(); err != nil {
			log.Println(err)
			return activity, exception.ErrInternalServer
		}

		return activity, nil
	} else if params.To != "" {
		rows, err := ar.DB.Query(fmt.Sprintf(`SELECT id, userID, deskripsi, created_at, update_at FROM %s WHERE DATE(created_at) BETWEEN '%s' AND '%s' ORDER BY created_at asc`, ar.TableName, params.From, params.To))
		if err != nil {
			log.Println(err)
			return activity, exception.ErrInternalServer
		}

		defer rows.Close()

		for rows.Next() {
			var c activitys.Activity
			if err := rows.Scan(
				&c.ID,
				&c.UserID,
				&c.Description,
				&c.CreatedAt,
				&c.UpdateAt,
			); err != nil {
				log.Println(err)
				return activity, exception.ErrInternalServer
			}
			activity = append(activity, c)
		}

		if err = rows.Err(); err != nil {
			log.Println(err)
			return activity, exception.ErrInternalServer
		}

		return activity, nil
	}

	return activity, nil
}
