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

func ParseMiRNADatabaseLinksTxt(ctx context.Context, r io.Reader) ([]*mirbase.MiRNADatabaseLinks, error) {
	r, lineCount, err := withLineCount(r)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	result := make([]*mirbase.MiRNADatabaseLinks, 0, lineCount)
	for scanner.Scan() {
		line := scanner.Text()
		record, err := parseMiRNADatabaseLinks(line)
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

func parseMiRNADatabaseLinks(line string) (record *mirbase.MiRNADatabaseLinks, err error) {
	fields := strings.Split(line, "\t")
	if len(fields) != 4 {
		return nil, fmt.Errorf("invalid number of fields (%d): %s", len(fields), line)
	}
	record = &mirbase.MiRNADatabaseLinks{
		Link:        fields[2],
		DisplayName: fields[3],
	}
	if record.AutoMiRNA, err = strconv.ParseInt(fields[0], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing AutoMiRNA: %w", err)
	}
	if record.AutoDB, err = strconv.ParseInt(fields[1], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing AutoDB: %w", err)
	}
	return record, nil
}
