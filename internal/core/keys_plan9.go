//go:build plan9
// +build plan9

package core

import (
	"errors"
	"io"
	"os"

	"github.com/reeflective/readline/internal/term"
)

// GetCursorPos returns the current cursor position in the terminal.
// On Plan 9, ANSI cursor-position queries are not supported, so this
// always returns -1, -1.
func (k *Keys) GetCursorPos() (x, y int) {
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return -1, -1
	}
	return -1, -1
}

func (k *Keys) readInputFiltered() (keys []byte, err error) {
	buf := make([]byte, keyScanBufSize)

	read, err := Stdin.Read(buf)
	if err != nil && errors.Is(err, io.EOF) {
		return
	}

	return buf[:read], nil
}
