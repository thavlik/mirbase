package parse

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/thavlik/mirbase/pkg/mirbase"
)

func ParseMiRNASpeciesTxt(ctx context.Context, r io.Reader) ([]*mirbase.MiRNASpecies, error) {
	r, lineCount, err := withLineCount(r)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	result := make([]*mirbase.MiRNASpecies, 0, lineCount)
	for scanner.Scan() {
		line := scanner.Text()
		record, err := parseMiRNASpecies(line)
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func parseMiRNASpecies(line string) (record *mirbase.MiRNASpecies, err error) {
	fields := strings.Split(line, "\t")
	if len(fields) != 9 {
		return nil, fmt.Errorf("invalid number of fields (%d): %s", len(fields), line)
	}
	record = &mirbase.MiRNASpecies{
		Organism:        maybeString(fields[1]),
		Division:        maybeString(fields[2]),
		Name:            maybeString(fields[3]),
		Taxonomy:        maybeString(fields[5]),
		GenomeAssembly:  fields[6],
		GenomeAccession: fields[7],
		EnsemblDB:       maybeString(fields[8]),
	}
	if record.AutoID, err = strconv.ParseInt(fields[0], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing AutoID: %w", err)
	}
	if record.TaxonID, err = maybeInt(fields[4]); err != nil {
		return nil, fmt.Errorf("error parsing TaxonID: %w", err)
	}
	return record, nil
}
