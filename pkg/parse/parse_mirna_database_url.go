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

func ParseMiRNADatabaseUrlTxt(ctx context.Context, r io.Reader) ([]*mirbase.MiRNADatabaseUrl, error) {
	r, lineCount, err := withLineCount(r)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	result := make([]*mirbase.MiRNADatabaseUrl, 0, lineCount)
	for scanner.Scan() {
		line := scanner.Text()
		record, err := parseMiRNADatabaseUrl(line)
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

func parseMiRNADatabaseUrl(line string) (record *mirbase.MiRNADatabaseUrl, err error) {
	fields := strings.Split(line, "\t")
	if len(fields) != 3 {
		return nil, fmt.Errorf("invalid number of fields (%d): %s", len(fields), line)
	}
	record = &mirbase.MiRNADatabaseUrl{
		DisplayName: fields[1],
		URL:         fields[2],
	}
	if record.AutoDB, err = strconv.ParseInt(fields[0], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing AutoDB: %w", err)
	}
	return record, nil
}
