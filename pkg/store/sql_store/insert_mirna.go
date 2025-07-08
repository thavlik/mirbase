package sql_store

import (
	"context"

	"github.com/pkg/errors"
	"github.com/thavlik/mirbase/pkg/mirbase"
)

func (s *sqlStore) InsertMiRNA(ctx context.Context, records []*mirbase.MiRNA) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}
	for _, record := range records {
		if _, err := tx.ExecContext(
			ctx,
			`INSERT INTO mirna (
  auto_mirna,
  mirna_acc,
  mirna_id,
  previous_mirna_id,
  description,
  sequence,
  comment,
  auto_species,
  dead_flag
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
			record.AutoMiRNA,
			record.MiRNAAcc,
			record.MiRNAID,
			record.PreviousMiRNAID,
			record.Description,
			record.Sequence,
			record.Comment,
			record.AutoSpecies,
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
