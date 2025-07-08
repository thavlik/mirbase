package sql_store

import (
	"context"

	"github.com/pkg/errors"
	"github.com/thavlik/mirbase/pkg/mirbase"
)

func (s *sqlStore) InsertMiRNAChromosomeBuild(ctx context.Context, records []*mirbase.MiRNAChromosomeBuild) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}
	for _, record := range records {
		if _, err := tx.ExecContext(
			ctx,
			`INSERT INTO mirna_chromosome_build (
  auto_mirna,
  xsome,
  contig_start,
  contig_end,
  strand
) VALUES ($1, $2, $3, $4, $5)`,
			record.AutoMiRNA,
			record.Xsome,
			record.ContigStart,
			record.ContigEnd,
			record.Strand,
		); err != nil {
			return errors.Wrapf(err, "failed to insert record: %v", record)
		}
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}
	return nil
}
