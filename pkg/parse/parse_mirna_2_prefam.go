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

func ParseMiRNA2PrefamTxt(ctx context.Context, r io.Reader) ([]*mirbase.MiRNA2Prefam, error) {
	r, lineCount, err := withLineCount(r)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	result := make([]*mirbase.MiRNA2Prefam, 0, lineCount)
	for scanner.Scan() {
		line := scanner.Text()
		record, err := parseMiRNA2Prefam(line)
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

func parseMiRNA2Prefam(line string) (record *mirbase.MiRNA2Prefam, err error) {
	fields := strings.Split(line, "\t")
	if len(fields) != 2 {
		return nil, fmt.Errorf("invalid number of fields (%d): %s", len(fields), line)
	}
	record = new(mirbase.MiRNA2Prefam)
	if record.AutoMiRNA, err = strconv.ParseInt(fields[0], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing AutoMiRNA: %w", err)
	}
	if record.AutoPrefam, err = strconv.ParseInt(fields[1], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing AutoPrefam: %w", err)
	}
	return record, nil
}
