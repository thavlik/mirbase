package build

import (
	"context"

	"github.com/thavlik/mirbase/pkg/store"

	"github.com/pkg/errors"
)

const dataRoot = "/mirbase"

func Build(ctx context.Context, output store.Store) error {
	if err := BuildConfidenceScore(ctx, output); err != nil {
		return errors.Wrap(err, "BuildConfidenceScore")
	}
	if err := BuildConfidence(ctx, output); err != nil {
		return errors.Wrap(err, "BuildConfidence")
	}
	if err := BuildDeadMiRNA(ctx, output); err != nil {
		return errors.Wrap(err, "BuildDeadMiRNA")
	}
	if err := BuildLiteratureReferences(ctx, output); err != nil {
		return errors.Wrap(err, "BuildLiteratureReferences")
	}
	if err := BuildMatureDatabaseLinks(ctx, output); err != nil {
		return errors.Wrap(err, "BuildMatureDatabaseLinks")
	}
	if err := BuildMatureDatabaseUrl(ctx, output); err != nil {
		return errors.Wrap(err, "BuildMatureDatabaseUrl")
	}
	if err := BuildMiRNA2Prefam(ctx, output); err != nil {
		return errors.Wrap(err, "BuildMiRNA2Prefam")
	}
	if err := BuildMiRNAChromosomeBuild(ctx, output); err != nil {
		return errors.Wrap(err, "BuildMiRNAChromosomeBuild")
	}
	if err := BuildMiRNAContext(ctx, output); err != nil {
		return errors.Wrap(err, "BuildMiRNAContext")
	}
	if err := BuildMiRNADatabaseLinks(ctx, output); err != nil {
		return errors.Wrap(err, "BuildMiRNADatabaseLinks")
	}
	if err := BuildMiRNADatabaseUrl(ctx, output); err != nil {
		return errors.Wrap(err, "BuildMiRNADatabaseUrl")
	}
	if err := BuildMiRNALiteratureReferences(ctx, output); err != nil {
		return errors.Wrap(err, "BuildMiRNALiteratureReferences")
	}
	if err := BuildMiRNAMature(ctx, output); err != nil {
		return errors.Wrap(err, "BuildMiRNAMature")
	}
	if err := BuildMiRNAPreMature(ctx, output); err != nil {
		return errors.Wrap(err, "BuildMiRNAPreMature")
	}
	if err := BuildMiRNAPrefam(ctx, output); err != nil {
		return errors.Wrap(err, "BuildMiRNAPrefam")
	}
	if err := BuildMiRNASpecies(ctx, output); err != nil {
		return errors.Wrap(err, "BuildMiRNASpecies")
	}
	if err := BuildMiRNA(ctx, output); err != nil {
		return errors.Wrap(err, "BuildMiRNA")
	}
	return nil
}
