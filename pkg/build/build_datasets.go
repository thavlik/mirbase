package build

import (
	"context"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/thavlik/mirbase/pkg/parse"
	"github.com/thavlik/mirbase/pkg/store"
)

func BuildConfidenceScore(ctx context.Context, output store.Store) error {
	r, err := open(filepath.Join(dataRoot, "confidence_score.txt"))
	if err != nil {
		return errors.Wrap(err, "open")
	}
	defer r.Close()
	result, err := parse.ParseConfidenceScoreTxt(ctx, r)
	if err != nil {
		return errors.Wrap(err, "parse")
	}
	if err := output.InsertConfidenceScore(ctx, result); err != nil {
		return errors.Wrap(err, "store")
	}
	return nil
}

func BuildConfidence(ctx context.Context, output store.Store) error {
	r, err := open(filepath.Join(dataRoot, "confidence.txt"))
	if err != nil {
		return errors.Wrap(err, "open")
	}
	defer r.Close()
	result, err := parse.ParseConfidenceTxt(ctx, r)
	if err != nil {
		return errors.Wrap(err, "parse")
	}
	if err := output.InsertConfidence(ctx, result); err != nil {
		return errors.Wrap(err, "store")
	}
	return nil
}

func BuildDeadMiRNA(ctx context.Context, output store.Store) error {
	r, err := open(filepath.Join(dataRoot, "dead_mirna.txt"))
	if err != nil {
		return errors.Wrap(err, "open")
	}
	defer r.Close()
	result, err := parse.ParseDeadMirnaTxt(ctx, r)
	if err != nil {
		return errors.Wrap(err, "parse")
	}
	if err := output.InsertDeadMiRNA(ctx, result); err != nil {
		return errors.Wrap(err, "store")
	}
	return nil
}

func BuildLiteratureReferences(ctx context.Context, output store.Store) error {
	r, err := open(filepath.Join(dataRoot, "literature_references.txt"))
	if err != nil {
		return errors.Wrap(err, "open")
	}
	defer r.Close()
	result, err := parse.ParseLiteratureReferencesTxt(ctx, r)
	if err != nil {
		return errors.Wrap(err, "parse")
	}
	if err := output.InsertLiteratureReferences(ctx, result); err != nil {
		return errors.Wrap(err, "store")
	}
	return nil
}

func BuildMatureDatabaseLinks(ctx context.Context, output store.Store) error {
	r, err := open(filepath.Join(dataRoot, "mature_database_links.txt"))
	if err != nil {
		return errors.Wrap(err, "open")
	}
	defer r.Close()
	result, err := parse.ParseMatureDatabaseLinksTxt(ctx, r)
	if err != nil {
		return errors.Wrap(err, "parse")
	}
	if err := output.InsertMatureDatabaseLinks(ctx, result); err != nil {
		return errors.Wrap(err, "store")
	}
	return nil
}

func BuildMatureDatabaseUrl(ctx context.Context, output store.Store) error {
	r, err := open(filepath.Join(dataRoot, "mature_database_url.txt"))
	if err != nil {
		return errors.Wrap(err, "open")
	}
	defer r.Close()
	result, err := parse.ParseMatureDatabaseUrlTxt(ctx, r)
	if err != nil {
		return errors.Wrap(err, "parse")
	}
	if err := output.InsertMatureDatabaseUrl(ctx, result); err != nil {
		return errors.Wrap(err, "store")
	}
	return nil
}

func BuildMiRNA2Prefam(ctx context.Context, output store.Store) error {
	r, err := open(filepath.Join(dataRoot, "mirna_2_prefam.txt"))
	if err != nil {
		return errors.Wrap(err, "open")
	}
	defer r.Close()
	result, err := parse.ParseMiRNA2PrefamTxt(ctx, r)
	if err != nil {
		return errors.Wrap(err, "parse")
	}
	if err := output.InsertMiRNA2Prefam(ctx, result); err != nil {
		return errors.Wrap(err, "store")
	}
	return nil
}

func BuildMiRNAChromosomeBuild(ctx context.Context, output store.Store) error {
	r, err := open(filepath.Join(dataRoot, "mirna_chromosome_build.txt"))
	if err != nil {
		return errors.Wrap(err, "open")
	}
	defer r.Close()
	result, err := parse.ParseMiRNAChromosomeBuildTxt(ctx, r)
	if err != nil {
		return errors.Wrap(err, "parse")
	}
	if err := output.InsertMiRNAChromosomeBuild(ctx, result); err != nil {
		return errors.Wrap(err, "store")
	}
	return nil
}

func BuildMiRNAContext(ctx context.Context, output store.Store) error {
	r, err := open(filepath.Join(dataRoot, "mirna_context.txt"))
	if err != nil {
		return errors.Wrap(err, "open")
	}
	defer r.Close()
	result, err := parse.ParseMiRNAContextTxt(ctx, r)
	if err != nil {
		return errors.Wrap(err, "parse")
	}
	if err := output.InsertMiRNAContext(ctx, result); err != nil {
		return errors.Wrap(err, "store")
	}
	return nil
}

func BuildMiRNADatabaseLinks(ctx context.Context, output store.Store) error {
	r, err := open(filepath.Join(dataRoot, "mirna_database_links.txt"))
	if err != nil {
		return errors.Wrap(err, "open")
	}
	defer r.Close()
	result, err := parse.ParseMiRNADatabaseLinksTxt(ctx, r)
	if err != nil {
		return errors.Wrap(err, "parse")
	}
	if err := output.InsertMiRNADatabaseLinks(ctx, result); err != nil {
		return errors.Wrap(err, "store")
	}
	return nil
}

func BuildMiRNADatabaseUrl(ctx context.Context, output store.Store) error {
	r, err := open(filepath.Join(dataRoot, "mirna_database_url.txt"))
	if err != nil {
		return errors.Wrap(err, "open")
	}
	defer r.Close()
	result, err := parse.ParseMiRNADatabaseUrlTxt(ctx, r)
	if err != nil {
		return errors.Wrap(err, "parse")
	}
	if err := output.InsertMiRNADatabaseUrl(ctx, result); err != nil {
		return errors.Wrap(err, "store")
	}
	return nil
}

func BuildMiRNALiteratureReferences(ctx context.Context, output store.Store) error {
	r, err := open(filepath.Join(dataRoot, "mirna_literature_references.txt"))
	if err != nil {
		return errors.Wrap(err, "open")
	}
	defer r.Close()
	result, err := parse.ParseMiRNALiteratureReferencesTxt(ctx, r)
	if err != nil {
		return errors.Wrap(err, "parse")
	}
	if err := output.InsertMiRNALiteratureReferences(ctx, result); err != nil {
		return errors.Wrap(err, "store")
	}
	return nil
}

func BuildMiRNAMature(ctx context.Context, output store.Store) error {
	r, err := open(filepath.Join(dataRoot, "mirna_mature.txt"))
	if err != nil {
		return errors.Wrap(err, "open")
	}
	defer r.Close()
	result, err := parse.ParseMiRNAMatureTxt(ctx, r)
	if err != nil {
		return errors.Wrap(err, "parse")
	}
	if err := output.InsertMiRNAMature(ctx, result); err != nil {
		return errors.Wrap(err, "store")
	}
	return nil
}

func BuildMiRNAPreMature(ctx context.Context, output store.Store) error {
	r, err := open(filepath.Join(dataRoot, "mirna_pre_mature.txt"))
	if err != nil {
		return errors.Wrap(err, "open")
	}
	defer r.Close()
	result, err := parse.ParseMiRNAPreMatureTxt(ctx, r)
	if err != nil {
		return errors.Wrap(err, "parse")
	}
	if err := output.InsertMiRNAPreMature(ctx, result); err != nil {
		return errors.Wrap(err, "store")
	}
	return nil
}

func BuildMiRNAPrefam(ctx context.Context, output store.Store) error {
	r, err := open(filepath.Join(dataRoot, "mirna_prefam.txt"))
	if err != nil {
		return errors.Wrap(err, "open")
	}
	defer r.Close()
	result, err := parse.ParseMiRNAPrefamTxt(ctx, r)
	if err != nil {
		return errors.Wrap(err, "parse")
	}
	if err := output.InsertMiRNAPrefam(ctx, result); err != nil {
		return errors.Wrap(err, "store")
	}
	return nil
}

func BuildMiRNASpecies(ctx context.Context, output store.Store) error {
	r, err := open(filepath.Join(dataRoot, "mirna_species.txt"))
	if err != nil {
		return errors.Wrap(err, "open")
	}
	defer r.Close()
	result, err := parse.ParseMiRNASpeciesTxt(ctx, r)
	if err != nil {
		return errors.Wrap(err, "parse")
	}
	if err := output.InsertMiRNASpecies(ctx, result); err != nil {
		return errors.Wrap(err, "store")
	}
	return nil
}

func BuildMiRNA(ctx context.Context, output store.Store) error {
	r, err := open(filepath.Join(dataRoot, "mirna.txt"))
	if err != nil {
		return errors.Wrap(err, "open")
	}
	defer r.Close()
	result, err := parse.ParseMiRNATxt(ctx, r)
	if err != nil {
		return errors.Wrap(err, "parse")
	}
	if err := output.InsertMiRNA(ctx, result); err != nil {
		return errors.Wrap(err, "store")
	}
	return nil
}
