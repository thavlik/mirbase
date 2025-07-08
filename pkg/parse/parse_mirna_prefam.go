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

func ParseMiRNAPrefamTxt(ctx context.Context, r io.Reader) ([]*mirbase.MiRNAPrefam, error) {
	r, lineCount, err := withLineCount(r)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	result := make([]*mirbase.MiRNAPrefam, 0, lineCount)
	for scanner.Scan() {
		line := scanner.Text()
		record, err := parseMiRNAPrefam(line)
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

func parseMiRNAPrefam(line string) (record *mirbase.MiRNAPrefam, err error) {
	fields := strings.Split(line, "\t")
	if len(fields) != 3 && len(fields) != 4 {
		return nil, fmt.Errorf("invalid number of fields (%d): %s", len(fields), line)
	}
	record = &mirbase.MiRNAPrefam{PrefamAcc: fields[1], PrefamID: fields[2]}
	if len(fields) == 4 {
		record.Description = &fields[3]
	}
	if record.AutoPrefam, err = strconv.ParseInt(fields[0], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing AutoPrefam: %w", err)
	}
	return record, nil
}
