package absensi

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"

	"github.com/Risuii/config"
	"github.com/Risuii/helpers/exception"
	"github.com/Risuii/models/absensis"
)

type (
	AbsensiRepository interface {
		Checkin(ctx context.Context, params absensis.Absensi) (int64, error)
		Checkout(ctx context.Context, checkinID int64, params absensis.Absensi) error
		Riwayat(ctx context.Context, name string) ([]absensis.Absensi, error)
		SendMsg(ctx context.Context, params absensis.Absensi) error
		ReceiveMsg() (absensis.Absensi, error)
	}

	absensiRepositoryImpl struct {
		db        *sql.DB
		tableName string
	}
)

func NewAbsensiRepositoryImpl(db *sql.DB, tableName string) AbsensiRepository {
	return &absensiRepositoryImpl{
		db:        db,
		tableName: tableName,
	}
}

func (ur *absensiRepositoryImpl) Checkin(ctx context.Context, params absensis.Absensi) (int64, error) {

	query := fmt.Sprintf(`INSERT INTO %s (userID, name, checkin) VALUES (?, ?, ?)`, ur.tableName)
	stmt, err := ur.db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return 0, exception.ErrInternalServer
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		params.UserID,
		params.Name,
		params.Checkin,
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

func (ur *absensiRepositoryImpl) Checkout(ctx context.Context, checkinID int64, params absensis.Absensi) error {
	query := fmt.Sprintf(`UPDATE %s SET checkout = ? WHERE id = %d`, ur.tableName, checkinID)
	stmt, err := ur.db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return exception.ErrInternalServer
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		params.Checkout,
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

func (ur *absensiRepositoryImpl) Riwayat(ctx context.Context, name string) ([]absensis.Absensi, error) {
	absensi := []absensis.Absensi{}

	rows, err := ur.db.Query(fmt.Sprintf(`SELECT id, userID, name, checkin, checkout FROM %s WHERE name = '%s'`, ur.tableName, name))
	if err != nil {
		log.Println(err)
		return absensi, exception.ErrInternalServer
	}

	defer rows.Close()

	for rows.Next() {
		var c absensis.Absensi
		var checkout sql.NullTime
		if err := rows.Scan(
			&c.ID,
			&c.UserID,
			&c.Name,
			&c.Checkin,
			&checkout,
		); err != nil {
			log.Println(err)
			return absensi, exception.ErrInternalServer
		}
		if checkout.Valid {
			c.Checkout = checkout.Time
		}
		absensi = append(absensi, c)
	}

	if err = rows.Err(); err != nil {
		log.Println(err)
		return absensi, exception.ErrInternalServer
	}

	return absensi, nil
}

func (ur *absensiRepositoryImpl) SendMsg(ctx context.Context, params absensis.Absensi) error {
	cfg := config.New()

	ch, err := cfg.Rabbitmq.RabbitCon.Channel()
	if err != nil {
		log.Println(err)
		return exception.ErrInternalServer
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Absensi",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
		return exception.ErrInternalServer
	}

	data, _ := json.Marshal(params)

	if err := ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	); err != nil {
		log.Println(err)
		return exception.ErrInternalServer
	}

	log.Println("Success Publish Message")

	return nil
}

func (ur *absensiRepositoryImpl) ReceiveMsg() (absensis.Absensi, error) {
	var absensi absensis.Absensi
	cfg := config.New()

	ch, err := cfg.Rabbitmq.RabbitCon.Channel()
	if err != nil {
		log.Println(err)
		return absensi, exception.ErrInternalServer
	}

	defer ch.Close()

	msg, err := ch.Consume(
		"Absensi",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	for d := range msg {
		err := json.Unmarshal(d.Body, &absensi)
		if err != nil {
			log.Println(err)
			return absensi, exception.ErrInternalServer
		}
		return absensi, nil
	}
	return absensi, nil
}
