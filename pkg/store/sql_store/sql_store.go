package sql_store

import (
	"database/sql"
	_ "embed"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/thavlik/mirbase/pkg/store"
)

//go:embed tables.sql
var tables string

func NewSqlStore(db *sql.DB, create bool) (store.Store, error) {
	if create {
		if err := createTables(db); err != nil {
			return nil, fmt.Errorf("failed to create tables: %v", err)
		}
	}
	return &sqlStore{db, create}, nil
}

type sqlStore struct {
	db     *sql.DB
	create bool
}

func createTables(db *sql.DB) error {
	statements := strings.Split(tables, ";")
	for i, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		if _, err := db.Exec(stmt); err != nil {
			return fmt.Errorf("failed to execute table creation script, statement #%d ('%s'): %v", i, stmt, err)
		}
	}
	return nil
}

func nullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: *s, Valid: true}
}

func nullInt(s *int64) sql.NullInt64 {
	if s == nil {
		return sql.NullInt64{}
	}
	return sql.NullInt64{Int64: *s, Valid: true}
}

func nullFloat(v *float64) sql.NullFloat64 {
	if v == nil {
		return sql.NullFloat64{}
	}
	return sql.NullFloat64{Float64: *v, Valid: true}
}
