package parse

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/thavlik/mirbase/pkg/mirbase"
)

func ParseDeadMirnaTxt(ctx context.Context, r io.Reader) ([]*mirbase.DeadMiRNA, error) {
	r, lineCount, err := withLineCount(r)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	result := make([]*mirbase.DeadMiRNA, 0, lineCount)
	for scanner.Scan() {
		line := scanner.Text()
		record, err := parseDeadMiRNA(line)
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

func parseDeadMiRNA(line string) (*mirbase.DeadMiRNA, error) {
	fields := strings.Split(line, "\t")
	if len(fields) != 5 {
		return nil, fmt.Errorf("invalid number of fields (%d): %s", len(fields), line)
	}
	return &mirbase.DeadMiRNA{
		MiRNAAcc:   fields[0],
		MiRNAID:    fields[1],
		PreviousID: fields[2],
		ForwardTo:  fields[3],
		Comment:    fields[4],
	}, nil
}
