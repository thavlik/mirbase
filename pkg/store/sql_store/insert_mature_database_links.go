package sql_store

import (
	"context"

	"github.com/pkg/errors"
	"github.com/thavlik/mirbase/pkg/mirbase"
)

func (s *sqlStore) InsertMatureDatabaseLinks(ctx context.Context, records []*mirbase.MatureDatabaseLinks) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}
	for _, record := range records {
		if _, err := tx.ExecContext(
			ctx,
			`INSERT INTO mature_database_links (
  auto_mature,
  auto_db,
  link,
  display_name
) VALUES ($1, $2, $3, $4)`,
			record.AutoMature,
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
