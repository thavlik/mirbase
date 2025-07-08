package sql_store

import (
	"context"

	"github.com/pkg/errors"
	"github.com/thavlik/mirbase/pkg/mirbase"
)

func (s *sqlStore) InsertMiRNALiteratureReferences(ctx context.Context, records []*mirbase.MiRNALiteratureReferences) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}
	for _, record := range records {
		if _, err := tx.ExecContext(
			ctx,
			`INSERT INTO mirna_literature_references (
  auto_mirna,
  auto_lit,
  comment,
  order_added
) VALUES ($1, $2, $3, $4)`,
			record.AutoMiRNA,
			record.AutoLit,
			record.Comment,
			record.OrderAdded,
		); err != nil {
			return errors.Wrapf(err, "failed to insert record: %v", record)
		}
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}
	return nil
}
