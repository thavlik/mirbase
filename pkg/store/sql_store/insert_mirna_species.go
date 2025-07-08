package sql_store

import (
	"context"

	"github.com/pkg/errors"
	"github.com/thavlik/mirbase/pkg/mirbase"
)

func (s *sqlStore) InsertMiRNASpecies(ctx context.Context, records []*mirbase.MiRNASpecies) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}
	for _, record := range records {
		if _, err := tx.ExecContext(
			ctx,
			`INSERT INTO mirna_species (
  auto_id,
  organism,
  division,
  name,
  taxon_id,
  taxonomy,
  genome_assembly,
  genome_accession,
  ensembl_db
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
			record.AutoID,
			nullString(record.Organism),
			nullString(record.Division),
			nullString(record.Name),
			nullInt(record.TaxonID),
			nullString(record.Taxonomy),
			record.GenomeAssembly,
			record.GenomeAccession,
			nullString(record.EnsemblDB),
		); err != nil {
			return errors.Wrapf(err, "failed to insert record: %v", record)
		}
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}
	return nil
}
