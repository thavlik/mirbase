package sql_store

import (
	"context"

	"github.com/pkg/errors"
	"github.com/thavlik/mirbase/pkg/mirbase"
)

func (s *sqlStore) InsertMiRNADatabaseLinks(ctx context.Context, records []*mirbase.MiRNADatabaseLinks) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}
	for _, record := range records {
		if _, err := tx.ExecContext(
			ctx,
			`INSERT INTO mirna_database_links (
  auto_mirna,
  auto_db,
  link,
  display_name
) VALUES ($1, $2, $3, $4)`,
			record.AutoMiRNA,
			record.AutoDB,
			record.Link,
			record.DisplayName,
		); err != nil {
			return errors.Wrapf(err, "failed to insert record: %v", record)
		}
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}
	return nil
}
