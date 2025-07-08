package sql_store

import (
	"context"

	"github.com/pkg/errors"
	"github.com/thavlik/mirbase/pkg/mirbase"
)

func (s *sqlStore) InsertConfidence(ctx context.Context, records []*mirbase.Confidence) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}
	for _, record := range records {
		if _, err := tx.ExecContext(
			ctx,
			`INSERT INTO confidence (
  mirna_id,
  auto_mirna,
  exp_count,
  _5p_count,
  _5p_raw_count,
  _3p_count,
  _3p_raw_count,
  _5p_consistent,
  _5p_mature_consistent,
  _3p_consistent,
  _3p_mature_consistent,
  _5p_overhang,
  _3p_overhang,
  energy_precursor,
  energy_by_length,
  paired_hairpin,
  mirdeep_score
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)`,
			record.MiRNAID,
			record.AutoMiRNA,
			record.ExpCount,
			record.Z_5pCount,
			record.Z_5pRawCount,
			record.Z_3pCount,
			record.Z_3pRawCount,
			record.Z_5pConsistent,
			record.Z_5pMatureConsistent,
			record.Z_3pConsistent,
			record.Z_3pMatureConsistent,
			record.Z_5pOverhang,
			record.Z_3pOverhang,
			nullFloat(record.EnergyPrecursor),
			record.EnergyByLength,
			record.PairedHairpin,
			record.MirdeepScore,
		); err != nil {
			return errors.Wrapf(err, "failed to insert record: %v", record)
		}
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}
	return nil
}
