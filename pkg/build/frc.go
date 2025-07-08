package build

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func open(path string) (io.ReadCloser, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("failed to stat file '%s': %v", path, err)
	}
	if info.Size() == 0 {
		return nil, fmt.Errorf("file '%s' is empty", path)
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file '%s': %v", path, err)
	}
	br := bufio.NewReader(f)
	r, _, err := br.ReadRune()
	if err != nil {
		return nil, err
	}
	if r != '\uFEFF' {
		br.UnreadRune() // Not a BOM -- put the rune back
	}
	fmt.Println("Input file opened:", path, "size:", info.Size())
	return &frc{br, f}, nil
}

type frc struct {
	r io.Reader
	f *os.File
}

func (b *frc) Read(p []byte) (n int, err error) {
	return b.r.Read(p)
}

func (b *frc) Close() error {
	return b.f.Close()
}
