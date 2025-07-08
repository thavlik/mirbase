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

func ParseMatureDatabaseUrlTxt(ctx context.Context, r io.Reader) ([]*mirbase.MatureDatabaseUrl, error) {
	r, lineCount, err := withLineCount(r)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	result := make([]*mirbase.MatureDatabaseUrl, 0, lineCount)
	for scanner.Scan() {
		line := scanner.Text()
		record, err := parseMatureDatabaseUrl(line)
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

func parseMatureDatabaseUrl(line string) (record *mirbase.MatureDatabaseUrl, err error) {
	fields := strings.Split(line, "\t")
	if len(fields) != 4 {
		return nil, fmt.Errorf("invalid number of fields (%d): %s", len(fields), line)
	}
	record = &mirbase.MatureDatabaseUrl{
		DisplayName: fields[1],
		URL:         fields[2],
	}
	if record.AutoDB, err = strconv.ParseInt(fields[0], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing AutoDB: %w", err)
	}
	if record.Type, err = strconv.ParseInt(fields[3], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing Type: %w", err)
	}
	return record, nil
}
