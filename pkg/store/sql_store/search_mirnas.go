package sql_store

import (
	"context"

	"github.com/pkg/errors"
	"github.com/thavlik/mirbase/pkg/store"
)

// SearchMiRNAs fuzzy searches the sqlite database for miRNAs,
// partially matching accession number, ID, description, sequence,
// or comment. Results are sorted by rank and paginated.
func (s *sqlStore) SearchMiRNAs(
	ctx context.Context,
	query string,
	pageSize int,
	page int,
) ([]*store.MiRNASearchResult, error) {
	rows, err := s.db.QueryContext(
		ctx,
		`SELECT mirna_acc, mirna_id,
  highlight(mirna_search, 2, '<b>', '</b>'),
  highlight(mirna_search, 3, '<b>', '</b>'),
  highlight(mirna_search, 4, '<b>', '</b>')
FROM mirna_search($1)
ORDER BY rank
LIMIT $2 OFFSET $3`,
		query,
		pageSize,
		(page-1)*pageSize,
	)
	if err != nil {
		return nil, errors.Wrap(err, "sql query")
	}
	defer rows.Close()
	results := []*store.MiRNASearchResult{}
	for rows.Next() {
		result := new(store.MiRNASearchResult)
		if err := rows.Scan(
			&result.MiRNAAcc,
			&result.MiRNAID,
			&result.Description,
			&result.Sequence,
			&result.Comment,
		); err != nil {
			return nil, errors.Wrap(err, "sql scan")
		}
		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "sql rows iterator")
	}
	return results, nil
}
