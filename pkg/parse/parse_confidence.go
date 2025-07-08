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

func ParseConfidenceTxt(ctx context.Context, r io.Reader) ([]*mirbase.Confidence, error) {
	r, lineCount, err := withLineCount(r)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	result := make([]*mirbase.Confidence, 0, lineCount)
	for scanner.Scan() {
		line := scanner.Text()
		record, err := parseConfidenceLine(line)
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

func parseConfidenceLine(line string) (record *mirbase.Confidence, err error) {
	fields := strings.Split(line, "\t")
	if len(fields) != 17 {
		return nil, fmt.Errorf("invalid number of fields (%d): %s", len(fields), line)
	}
	record = &mirbase.Confidence{MiRNAID: fields[0]}
	if record.AutoMiRNA, err = strconv.ParseInt(fields[1], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing AutoMiRNA: %w", err)
	}
	if record.ExpCount, err = strconv.ParseInt(fields[2], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing ExpCount: %w", err)
	}
	if record.Z_5pCount, err = strconv.ParseFloat(fields[3], 64); err != nil {
		return nil, fmt.Errorf("error parsing 5pCount: %w", err)
	}
	if record.Z_5pRawCount, err = strconv.ParseFloat(fields[4], 64); err != nil {
		return nil, fmt.Errorf("error parsing 5pRawCount: %w", err)
	}
	if record.Z_3pCount, err = strconv.ParseFloat(fields[5], 64); err != nil {
		return nil, fmt.Errorf("error parsing 3pCount: %w", err)
	}
	if record.Z_3pRawCount, err = strconv.ParseFloat(fields[6], 64); err != nil {
		return nil, fmt.Errorf("error parsing 3pRawCount: %w", err)
	}
	if record.Z_5pConsistent, err = strconv.ParseFloat(fields[7], 64); err != nil {
		return nil, fmt.Errorf("error parsing 5pConsistent: %w", err)
	}
	if record.Z_5pMatureConsistent, err = strconv.ParseFloat(fields[8], 64); err != nil {
		return nil, fmt.Errorf("error parsing 5pMatureConsistent: %w", err)
	}
	if record.Z_3pConsistent, err = strconv.ParseFloat(fields[9], 64); err != nil {
		return nil, fmt.Errorf("error parsing 3pConsistent: %w", err)
	}
	if record.Z_3pMatureConsistent, err = strconv.ParseFloat(fields[10], 64); err != nil {
		return nil, fmt.Errorf("error parsing 3pMatureConsistent: %w", err)
	}
	if record.Z_5pOverhang, err = strconv.ParseInt(fields[11], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing 5pOverhang: %w", err)
	}
	if record.Z_3pOverhang, err = strconv.ParseInt(fields[12], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing 3pOverhang: %w", err)
	}
	if fields[13] != "\\N" {
		v, err := strconv.ParseFloat(fields[13], 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing EnergyPrecursor: %w", err)
		}
		record.EnergyPrecursor = &v
	}
	if record.EnergyByLength, err = strconv.ParseFloat(fields[14], 64); err != nil {
		return nil, fmt.Errorf("error parsing EnergyByLength: %w", err)
	}
	if record.PairedHairpin, err = strconv.ParseFloat(fields[15], 64); err != nil {
		return nil, fmt.Errorf("error parsing PairedHairpin: %w", err)
	}
	if record.MirdeepScore, err = strconv.ParseFloat(fields[16], 64); err != nil {
		return nil, fmt.Errorf("error parsing MirdeepScore: %w", err)
	}
	return record, nil
}
