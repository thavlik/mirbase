package parse

import (
	"bytes"
	"html"
	"io"
	"strconv"
	"strings"
)

func withLineCount(r io.Reader) (io.Reader, int, error) {
	body, err := io.ReadAll(r)
	if err != nil {
		return nil, 0, err
	}

	unescaped := html.UnescapeString(string(body))
	lineCount := strings.Count(unescaped, "\n")
	return bytes.NewBuffer([]byte(unescaped)), lineCount, nil
}

func isSet(s string) bool {
	return s != "" && s != "\\N"
}

func maybeString(s string) *string {
	if isSet(s) {
		return &s
	}
	return nil
}

func maybeInt(s string) (*int64, error) {
	if !isSet(s) {
		return nil, nil
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
