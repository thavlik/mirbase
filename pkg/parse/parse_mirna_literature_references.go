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

func ParseMiRNALiteratureReferencesTxt(ctx context.Context, r io.Reader) ([]*mirbase.MiRNALiteratureReferences, error) {
	r, lineCount, err := withLineCount(r)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	result := make([]*mirbase.MiRNALiteratureReferences, 0, lineCount)
	for scanner.Scan() {
		line := scanner.Text()
		record, err := parseMiRNALiteratureReferences(line)
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

func parseMiRNALiteratureReferences(line string) (record *mirbase.MiRNALiteratureReferences, err error) {
	fields := strings.Split(line, "\t")
	if len(fields) != 4 {
		return nil, fmt.Errorf("invalid number of fields (%d): %s", len(fields), line)
	}
	record = &mirbase.MiRNALiteratureReferences{Comment: fields[2]}
	if record.AutoMiRNA, err = strconv.ParseInt(fields[0], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing AutoMiRNA: %w", err)
	}
	if record.AutoLit, err = strconv.ParseInt(fields[1], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing AutoLit: %w", err)
	}
	if record.OrderAdded, err = strconv.ParseInt(fields[3], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing OrderAdded: %w", err)
	}
	return record, nil
}
