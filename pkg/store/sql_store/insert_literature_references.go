package sql_store

import (
	"context"

	"github.com/pkg/errors"
	"github.com/thavlik/mirbase/pkg/mirbase"
)

func (s *sqlStore) InsertLiteratureReferences(ctx context.Context, records []*mirbase.LiteratureReferences) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}

	for _, record := range records {
		if _, err := tx.ExecContext(
			ctx,
			`INSERT INTO literature_references (
  auto_lit,
  medline,
  title,
  author,
  journal
) VALUES ($1, $2, $3, $4, $5)`,
			record.AutoLit,
			nullInt(record.Medline),
			record.Title,
			record.Author,
			record.Journal,
		); err != nil {
			return errors.Wrapf(err, "failed to insert record: %v", record)
		}
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}
	return nil
}
