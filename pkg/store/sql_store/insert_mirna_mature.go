package sql_store

import (
	"context"

	"github.com/pkg/errors"
	"github.com/thavlik/mirbase/pkg/mirbase"
)

func (s *sqlStore) InsertMiRNAMature(ctx context.Context, records []*mirbase.MiRNAMature) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}
	for _, record := range records {
		if _, err := tx.ExecContext(
			ctx,
			`INSERT INTO mirna_mature (
  auto_mature,
  mature_name,
  previous_mature_id,
  mature_acc,
  evidence,
  experiment,
  similarity,
  dead_flag
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
			record.AutoMature,
			record.MatureName,
			record.PreviousMatureID,
			record.MatureAcc,
			record.Evidence,
			record.Experiment,
			record.Similarity,
			record.DeadFlag,
		); err != nil {
			return errors.Wrapf(err, "failed to insert record: %v", record)
		}
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}
	return nil
}
