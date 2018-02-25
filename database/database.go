package database

import (
	"context"
	"database/sql"
	"log"
	"os"

	"fmt"
	"time"

	"github.com/Sharykhin/gl-mail-grpc"
	_ "github.com/go-sql-driver/mysql" // mysql driver dependency
)

// Storage is a reference to a private struct with all necessary methods
var Storage storage

type storage struct {
	db *sql.DB
}

func (s storage) Create(ctx context.Context, fmr *api.FailMailRequest) (*api.FailMailResponse, error) {
	res, err := s.db.ExecContext(ctx, "INSERT INTO failed_mails(`action`, `payload`, `reason`, `created_at`) VALUES(?, ?, ?, NOW())", fmr.Action, string(fmr.Payload), fmr.Reason)
	if err != nil {
		return nil, fmt.Errorf("could not create a new failed message: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("could not get last insert id: %v", err)
	}

	return &api.FailMailResponse{
		ID:        id,
		Action:    fmr.Action,
		Payload:   fmr.Payload,
		Reason:    fmr.Reason,
		CreatedAt: time.Now().Format(time.RFC822),
		DeletedAt: "",
	}, nil
}

func init() {
	dbSource := os.Getenv("DB_SOURCE")
	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		log.Fatalf("cone not connect to mysql: %v", err)
	}

	Storage = storage{db: db}
}
