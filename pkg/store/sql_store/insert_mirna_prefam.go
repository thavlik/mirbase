package sql_store

import (
	"context"

	"github.com/pkg/errors"
	"github.com/thavlik/mirbase/pkg/mirbase"
)

func (s *sqlStore) InsertMiRNAPrefam(ctx context.Context, records []*mirbase.MiRNAPrefam) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}
	for _, record := range records {
		if _, err := tx.ExecContext(
			ctx,
			`INSERT INTO mirna_prefam (
  auto_prefam,
  prefam_acc,
  prefam_id,
  description
) VALUES ($1, $2, $3, $4)`,
			record.AutoPrefam,
			record.PrefamAcc,
			record.PrefamID,
			nullString(record.Description),
		); err != nil {
			return errors.Wrapf(err, "failed to insert record: %v", record)
		}
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}
	return nil
}
