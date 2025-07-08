package init_store

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/thavlik/mirbase/pkg/store"
	"github.com/thavlik/mirbase/pkg/store/sql_store"
)

func InitStore(path string, create bool) (store.Store, error) {
	query := []string{
		"_sync=FULL",
		"_journal_mode=PERSIST",
		"_auto_vacuum=FULL",
		"_foreign_keys=yes",
		"_defer_foreign_keys=no",
		"_case_sensitive_like=no",
		"cache=shared",
	}
	if create {
		query = append(query, "mode=rwc")
	} else {
		query = append(
			query,
			//"immutable=yes",
			"mode=rw", // ro
			"_query_only=yes",
		)
		info, err := os.Stat(path)
		if err != nil {
			return nil, fmt.Errorf("failed to stat database file: %v", err)
		}
		sizeInMb := info.Size() / 1000 / 1000
		if sizeInMb < 1 {
			return nil, fmt.Errorf("database file '%s' is too small: %d KiB", path, info.Size()/1000)
		}
	}
	db, err := sql.Open(
		"sqlite3",
		fmt.Sprintf(
			"file:%s?%s",
			path,
			strings.Join(query, "&"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	output, err := sql_store.NewSqlStore(db, create)
	if err != nil {
		return nil, fmt.Errorf("failed to create SQL store: %v", err)
	}
	return output, nil
}
