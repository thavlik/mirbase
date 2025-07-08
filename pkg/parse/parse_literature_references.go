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

func ParseLiteratureReferencesTxt(ctx context.Context, r io.Reader) ([]*mirbase.LiteratureReferences, error) {
	r, lineCount, err := withLineCount(r)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	result := make([]*mirbase.LiteratureReferences, 0, lineCount)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line == "\\" {
			// There's a bug in the database release where
			// some lines are just empty or a backslash.
			continue
		}
		record, err := parseLiteratureReference(line)
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

func parseLiteratureReference(line string) (record *mirbase.LiteratureReferences, err error) {
	fields := strings.Split(line, "\t")
	if len(fields) != 5 {
		return nil, fmt.Errorf("invalid number of fields (%d): %s", len(fields), line)
	}
	record = &mirbase.LiteratureReferences{
		Title:   fields[2],
		Author:  fields[3],
		Journal: fields[4],
	}
	if record.AutoLit, err = strconv.ParseInt(fields[0], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing AutoLit: %w", err)
	}
	if fields[1] != "\\N" {
		v, err := strconv.ParseInt(fields[1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing Medline: %w", err)
		}
		record.Medline = &v
	}
	return record, nil
}
