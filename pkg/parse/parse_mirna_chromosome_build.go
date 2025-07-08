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

func ParseMiRNAChromosomeBuildTxt(ctx context.Context, r io.Reader) ([]*mirbase.MiRNAChromosomeBuild, error) {
	r, lineCount, err := withLineCount(r)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	result := make([]*mirbase.MiRNAChromosomeBuild, 0, lineCount)
	for scanner.Scan() {
		line := scanner.Text()
		record, err := parseMiRNAChromosomeBuild(line)
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

func parseMiRNAChromosomeBuild(line string) (record *mirbase.MiRNAChromosomeBuild, err error) {
	fields := strings.Split(line, "\t")
	if len(fields) != 5 {
		return nil, fmt.Errorf("invalid number of fields (%d): %s", len(fields), line)
	}
	record = &mirbase.MiRNAChromosomeBuild{
		Xsome:  fields[1],
		Strand: fields[4],
	}
	if record.AutoMiRNA, err = strconv.ParseInt(fields[0], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing AutoMiRNA: %w", err)
	}
	if record.ContigStart, err = strconv.ParseInt(fields[2], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing ContigStart: %w", err)
	}
	if record.ContigEnd, err = strconv.ParseInt(fields[3], 10, 64); err != nil {
		return nil, fmt.Errorf("error parsing ContigEnd: %w", err)
	}
	return record, nil
}
