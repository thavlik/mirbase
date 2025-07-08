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

func ParseMiRNATxt(ctx context.Context, r io.Reader) ([]*mirbase.MiRNA, error) {
	r, lineCount, err := withLineCount(r)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	result := make([]*mirbase.MiRNA, 0, lineCount)
	for scanner.Scan() {
		line := scanner.Text()
		record, err := parseMiRNA(line)
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

func parseMiRNA(line string) (record *mirbase.MiRNA, err error) {
	fields := strings.Split(line, "\t")
	if len(fields) != 9 {
		return nil, fmt.Errorf("invalid number of fields (%d): %s", len(fields), line)
	}
	record = &mirbase.MiRNA{
		MiRNAAcc:        fields[1],
		MiRNAID:         fields[2],
		PreviousMiRNAID: fields[3],
		Description:     fields[4],
		Sequence:        []byte(fields[5]),
		Comment:         fields[6],
		DeadFlag:        fields[8] == "1",
	}
	if record.AutoMiRNA, err = strconv.ParseInt(fields[0], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing AutoMiRNA: %w", err)
	}
	if record.AutoSpecies, err = strconv.ParseInt(fields[7], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing AutoSpecies: %w", err)
	}
	return record, nil
}
