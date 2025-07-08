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

func ParseConfidenceScoreTxt(ctx context.Context, r io.Reader) ([]*mirbase.ConfidenceScore, error) {
	r, lineCount, err := withLineCount(r)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	result := make([]*mirbase.ConfidenceScore, 0, lineCount)
	for scanner.Scan() {
		line := scanner.Text()
		record, err := parseConfidenceScore(line)
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

func parseConfidenceScore(line string) (record *mirbase.ConfidenceScore, err error) {
	fields := strings.Split(line, "\t")
	if len(fields) != 2 {
		return nil, fmt.Errorf("invalid number of fields (%d): %s", len(fields), line)
	}
	record = new(mirbase.ConfidenceScore)
	if record.AutoMiRNA, err = strconv.ParseInt(fields[0], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing AutoMiRNA: %w", err)
	}
	if record.Confidence, err = strconv.ParseInt(fields[1], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing Confidence: %w", err)
	}
	return record, nil
}
