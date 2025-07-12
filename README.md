# mirbase
[<img src="https://img.shields.io/docker/image-size/thavlik/mirbase/latest">](https://hub.docker.com/r/thavlik/mirbase)
[<img src="https://img.shields.io/badge/maintenance%20status-actively%20developed-brightgreen">](https://github.com/thavlik/mirbase)
[<img src="https://img.shields.io/badge/Language-go-01add8.svg">](https://go.dev/)
[<img src="https://img.shields.io/badge/License-MIT-lightblue.svg">](./LICENSE)

This repository contains a Dockerfile for building a [sqlite](https://www.sqlite.org/) database from the [miRBase data](https://www.mirbase.org/download/). 

## Building
Run [./scripts/build.sh](./scripts/build.sh) or run the following command:
```bash
docker build -t thavlik/mirbase:latest .
```

## Schema
See [tables.sql](pkg/store/sql_store/tables.sql) for how the tables are created. Note that the schema used for sqlite differs from the official miRBase release.


## Usage
There are two ways to utilize the prebuilt image: as a base image, and the `mirbase` command line.

When used as a base image, where you can find the database file at `/mirbase.sqlite`:

```Dockerfile
# Create a reference to the database image as a build stage.
FROM thavlik/mirbase:latest AS db

# Create your image here.
FROM debian:latest
COPY --from=db /mirbase.sqlite /opt/mirbase.sqlite
#
# Now your docker image contains the prebuilt
# sqlite database file at /opt/mirbase.sqlite
#
```

Note: the only time you must build this image is if you want to change how data is handled. For example, you want to insert into a PostgreSQL server instead of writing to a sqlite file. If sqlite is sufficient for you, you are encouraged to use the [prebuilt image](https://hub.docker.com/r/thavlik/mirbase).

### Loading
See [init_store.go](pkg/store/init_store/init_store.go) for code that demonstrates how to open the database. It is advised to optimize your connection query for your application. Note that [go-sqlite3](https://github.com/mattn/go-sqlite3) requires compiling with `CGO_ENABLED=1`.

### Searching
[tables.sql](pkg/store/sql_store/tables.sql) creates an [fts5](https://www.sqlite.org/fts5.html) virtual table to fuzzy search for rows in the `mirna` table. The trigram tokenizer provides a fuzzy searching behavior. 

Here is some example code written in [Go](https://go.dev/) that demonstrates paginated fuzzy searching and applies bold (\<b\> and \</b\>) tags to the matched text. Note that [go-sqlite3](https://github.com/mattn/go-sqlite3) with the search function requires compiling with **both** `CGO_ENABLED=1` **and** `-tags "fts5"`.
```go
type MiRNASearchResult struct {
	MiRNAAcc    string `json:"mirna_acc"`
	MiRNAID     string `json:"mirna_id"`
	Description string `json:"description,omitempty"`
	Sequence    string `json:"sequence,omitempty"`
	Comment     string `json:"comment,omitempty"`
}

// SearchMiRNAs fuzzy searches the sqlite database for miRNAs,
// partially matching accession number, ID, description, sequence,
// or comment. Results are sorted by rank and paginated.
func SearchMiRNAs(
	ctx context.Context,
	db *sql.DB,
	query string,
	pageSize int,
	page int,
) ([]*MiRNASearchResult, error) {
	rows, err := db.QueryContext(
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
	results := []*MiRNASearchResult{}
	for rows.Next() {
		result := new(MiRNASearchResult)
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
```

The above function is implemented as the `search` subcommand, which can be ran as follows:
```bash
docker run -it --rm thavlik/mirbase:latest mirbase search 12180
```
```json
[
  {
    "mirna_acc": "MI0040618",
    "mirna_id": "mmu-mir-12180",
    "description": "Mus musculus miR-\u003cb\u003e12180\u003c/b\u003e stem-loop",
    "sequence": "AGUGUUCCAGCAUGGAAGGGGAGGGGUUCCUGAGCUUGUGUCUUUAACCAAGGAGCUGUGGACACUUGA"
  },
  {
    "mirna_acc": "MI0012180",
    "mirna_id": "bmo-mir-2731-2",
    "description": "Bombyx mori miR-2731-2 stem-loop",
    "sequence": "AAACCUAACAGAUGCGAGACCAUGGUAUGUGGAAAUAAAAGCCACCCGGUUUUUAUCUUUCCACACCAAAAGAUCCACAUUUCCCGAGGUG"
  }
]
```

## License
The code in this repository for building the database is [licensed under MIT](./LICENSE).