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

func ParseMiRNAMatureTxt(ctx context.Context, r io.Reader) ([]*mirbase.MiRNAMature, error) {
	r, lineCount, err := withLineCount(r)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	result := make([]*mirbase.MiRNAMature, 0, lineCount)
	for scanner.Scan() {
		line := scanner.Text()
		record, err := parseMiRNAMature(line)
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

func parseMiRNAMature(line string) (record *mirbase.MiRNAMature, err error) {
	fields := strings.Split(line, "\t")
	if len(fields) != 8 {
		return nil, fmt.Errorf("invalid number of fields (%d): %s", len(fields), line)
	}
	record = &mirbase.MiRNAMature{
		MatureName:       fields[1],
		PreviousMatureID: fields[2],
		MatureAcc:        fields[3],
		Evidence:         fields[4],
		Experiment:       fields[5],
		Similarity:       fields[6],
		DeadFlag:         fields[7] == "1",
	}
	if record.AutoMature, err = strconv.ParseInt(fields[0], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing AutoMature: %w", err)
	}
	return record, nil
}
