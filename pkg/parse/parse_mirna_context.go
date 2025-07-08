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

func ParseMiRNAContextTxt(ctx context.Context, r io.Reader) ([]*mirbase.MiRNAContext, error) {
	r, lineCount, err := withLineCount(r)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	result := make([]*mirbase.MiRNAContext, 0, lineCount)
	for scanner.Scan() {
		line := scanner.Text()
		record, err := parseMiRNAContext(line)
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

func parseMiRNAContext(line string) (record *mirbase.MiRNAContext, err error) {
	fields := strings.Split(line, "\t")
	if len(fields) != 7 {
		return nil, fmt.Errorf("invalid number of fields (%d): %s", len(fields), line)
	}
	record = &mirbase.MiRNAContext{
		TranscriptID: fields[1],
		OverlapSense: fields[2],
		OverlapType:  fields[3],
	}
	if fields[5] != "" {
		record.TranscriptSource = &fields[5]
	}
	if fields[6] != "" {
		record.TranscriptName = &fields[6]
	}
	if record.AutoMiRNA, err = strconv.ParseInt(fields[0], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing AutoMiRNA: %w", err)
	}
	if record.Number, err = strconv.ParseInt(fields[4], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing Number: %w", err)
	}
	return record, nil
}
