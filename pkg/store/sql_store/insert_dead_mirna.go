package sql_store

import (
	"context"

	"github.com/pkg/errors"
	"github.com/thavlik/mirbase/pkg/mirbase"
)

func (s *sqlStore) InsertDeadMiRNA(ctx context.Context, records []*mirbase.DeadMiRNA) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}
	for _, record := range records {
		if _, err := tx.ExecContext(
			ctx,
			`INSERT INTO dead_mirna (
  mirna_acc,
  mirna_id,
  previous_id,
  forward_to,
  comment
) VALUES ($1, $2, $3, $4, $5)`,
			record.MiRNAAcc,
			record.MiRNAID,
			record.PreviousID,
			record.ForwardTo,
			record.Comment,
		); err != nil {
			return errors.Wrapf(err, "failed to insert record: %v", record)
		}
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}
	return nil
}
