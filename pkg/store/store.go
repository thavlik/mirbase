package store

import (
	"context"
	"fmt"

	"github.com/thavlik/mirbase/pkg/mirbase"
)

var ErrNotFound = fmt.Errorf("record not found")

type MiRNASearchResult struct {
	MiRNAAcc    string `json:"mirna_acc"`
	MiRNAID     string `json:"mirna_id"`
	Description string `json:"description,omitempty"`
	Sequence    string `json:"sequence,omitempty"`
	Comment     string `json:"comment,omitempty"`
}

type Store interface {
	Close() error

	InsertConfidenceScore(context.Context, []*mirbase.ConfidenceScore) error
	InsertConfidence(context.Context, []*mirbase.Confidence) error
	InsertDeadMiRNA(context.Context, []*mirbase.DeadMiRNA) error
	InsertLiteratureReferences(context.Context, []*mirbase.LiteratureReferences) error
	InsertMatureDatabaseLinks(context.Context, []*mirbase.MatureDatabaseLinks) error
	InsertMatureDatabaseUrl(context.Context, []*mirbase.MatureDatabaseUrl) error
	InsertMiRNA2Prefam(context.Context, []*mirbase.MiRNA2Prefam) error
	InsertMiRNAChromosomeBuild(context.Context, []*mirbase.MiRNAChromosomeBuild) error
	InsertMiRNAContext(context.Context, []*mirbase.MiRNAContext) error
	InsertMiRNADatabaseLinks(context.Context, []*mirbase.MiRNADatabaseLinks) error
	InsertMiRNADatabaseUrl(context.Context, []*mirbase.MiRNADatabaseUrl) error
	InsertMiRNALiteratureReferences(context.Context, []*mirbase.MiRNALiteratureReferences) error
	InsertMiRNAMature(context.Context, []*mirbase.MiRNAMature) error
	InsertMiRNAPreMature(context.Context, []*mirbase.MiRNAPreMature) error
	InsertMiRNAPrefam(context.Context, []*mirbase.MiRNAPrefam) error
	InsertMiRNASpecies(context.Context, []*mirbase.MiRNASpecies) error
	InsertMiRNA(context.Context, []*mirbase.MiRNA) error

	SearchMiRNAs(ctx context.Context, query string, pageSize int, page int) ([]*MiRNASearchResult, error)
}
