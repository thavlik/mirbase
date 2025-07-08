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

func ParseMatureDatabaseLinksTxt(ctx context.Context, r io.Reader) ([]*mirbase.MatureDatabaseLinks, error) {
	r, lineCount, err := withLineCount(r)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	result := make([]*mirbase.MatureDatabaseLinks, 0, lineCount)
	for scanner.Scan() {
		line := scanner.Text()
		record, err := parseMatureDatabaseLink(line)
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

func parseMatureDatabaseLink(line string) (record *mirbase.MatureDatabaseLinks, err error) {
	fields := strings.Split(line, "\t")
	if len(fields) != 4 {
		return nil, fmt.Errorf("invalid number of fields (%d): %s", len(fields), line)
	}
	record = &mirbase.MatureDatabaseLinks{
		Link:        fields[2],
		DisplayName: fields[3],
	}
	if record.AutoMature, err = strconv.ParseInt(fields[0], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing AutoMature: %w", err)
	}
	if record.AutoDB, err = strconv.ParseInt(fields[1], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing AutoDB: %w", err)
	}
	return record, nil
}
