package sql_store

import (
	"context"

	"github.com/pkg/errors"
	"github.com/thavlik/mirbase/pkg/mirbase"
)

func (s *sqlStore) InsertMiRNAContext(ctx context.Context, records []*mirbase.MiRNAContext) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}
	for _, record := range records {
		if _, err := tx.ExecContext(
			ctx,
			`INSERT INTO mirna_context (
  auto_mirna,
  transcript_id,
  overlap_sense,
  overlap_type,
  number,
  transcript_source,
  transcript_name
) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
			record.AutoMiRNA,
			record.TranscriptID,
			record.OverlapSense,
			record.OverlapType,
			record.Number,
			nullString(record.TranscriptSource),
			nullString(record.TranscriptName),
		); err != nil {
			return errors.Wrapf(err, "failed to insert record: %v", record)
		}
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}
	return nil
}
