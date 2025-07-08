package sql_store

import (
	"context"

	"github.com/pkg/errors"
	"github.com/thavlik/mirbase/pkg/mirbase"
)

func (s *sqlStore) InsertMiRNAPreMature(ctx context.Context, records []*mirbase.MiRNAPreMature) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}
	for _, record := range records {
		if _, err := tx.ExecContext(
			ctx,
			`INSERT INTO mirna_pre_mature (
  auto_mirna,
  auto_mature,
  mature_from,
  mature_to
) VALUES ($1, $2, $3, $4)`,
			record.AutoMiRNA,
			record.AutoMature,
			record.MatureFrom,
			record.MatureTo,
		); err != nil {
			return errors.Wrapf(err, "failed to insert record: %v", record)
		}
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}
	return nil
}
