package sql_store

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

func (s *sqlStore) Close() error {
	if s.create {
		fmt.Println("Closing database, running VACUUM to optimize space usage...")
		if _, err := s.db.Exec("VACUUM;"); err != nil {
			return errors.Wrap(err, "failed to vacuum database")
		}
		fmt.Println("VACUUM completed.")
	}
	return s.db.Close()
}
