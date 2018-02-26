package database

import (
	"context"
	"database/sql"
	"log"
	"os"

	"fmt"
	"time"

	"github.com/Sharykhin/gl-mail-grpc"
	"github.com/Sharykhin/gl-mail-grpc-server/entity"
	_ "github.com/go-sql-driver/mysql" // mysql driver dependency
)

// Storage is a reference to a private struct with all necessary methods
var Storage storage

type storage struct {
	db *sql.DB
}

func (s storage) Create(ctx context.Context, fmr *api.FailMailRequest) (*entity.FailMail, error) {
	res, err := s.db.ExecContext(ctx, "INSERT INTO failed_mails(`action`, `payload`, `reason`, `created_at`) VALUES(?, ?, ?, NOW())", fmr.Action, string(fmr.Payload), fmr.Reason)
	if err != nil {
		return nil, fmt.Errorf("could not create a new failed message: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("could not get last insert id: %v", err)
	}

	return &entity.FailMail{
		ID:        id,
		Action:    fmr.Action,
		Payload:   entity.Payload(fmr.Payload),
		Reason:    fmr.Reason,
		CreatedAt: entity.JSONTime(time.Now()),
		DeletedAt: nil,
	}, nil

	//return &api.FailMailResponse{
	//	ID:        id,
	//	Action:    fmr.Action,
	//	Payload:   fmr.Payload,
	//	Reason:    fmr.Reason,
	//	CreatedAt: time.Now().Format(time.RFC822),
	//	DeletedAt: "",
	//}, nil
}

func (s storage) GetList(ctx context.Context, limit, offset int64) ([]api.FailMailResponse, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT `id`, `action`, `reason`, `payload`, `created_at` FROM failed_mails LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, fmt.Errorf("could not make a select statement: %v", err)
	}
	defer rows.Close() // nolint: errcheck

	var fm []api.FailMailResponse
	for rows.Next() {
		var m api.FailMailResponse
		err := rows.Scan(&m.ID, &m.Action, &m.Reason, &m.Payload, &m.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("could not scan a row to struct %v: %v", m, err)
		}
		fm = append(fm, m)
	}

	return fm, rows.Err()
}

func init() {
	dbSource := os.Getenv("DB_SOURCE")
	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		log.Fatalf("cone not connect to mysql: %v", err)
	}

	Storage = storage{db: db}
}
